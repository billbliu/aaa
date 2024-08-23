package business

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/redis_model"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/golang-module/carbon/v2"
)

type BaseService struct{}

var BaseServiceApp = new(BaseService)

// @author: [bill]
// @function: GetEmailSmsCode
// @description: 获取邮箱短信验证码
// @param: email string, smsType string
// @return: *response.GetSMSRes, error
func (s *BaseService) GetEmailSmsCode(email string, smsType string) (*response.GetSMSRes, error) {
	redisKey := s.newEmailRedisKey(smsType, email)
	// 查找邮箱是否已发送过邮箱验证码并且还未过期
	rSmsCode, err := s.getRedisSmsCode(redisKey)

	// 未找到,重新生成验证码并发送
	if err != nil {
		now := carbon.Now(carbon.PRC)
		code := utils.Captcha(6)
		item := redis_model.RedisSmsCode{
			Type:     smsType,
			SmsCode:  code,
			SendTime: carbon.DateTime{Carbon: now},
		}

		res, err1 := s.sendEmailSmsCode(email, smsType, code)
		if err1 != nil {
			return nil, err1
		}

		if err1 := s.setRedisSmsCode(redisKey, item); err1 != nil {
			global.GVA_LOG.Error(err1.Error())
		}

		return res, nil
	}

	// 找到
	if gte := carbon.Now(carbon.PRC).Gte(rSmsCode.SendTime.AddSeconds(int(global.SmsCodeSendInterval))); gte {
		// 已经存在验证码, 但是验证码超过间隔发送时间, 发送新的
		now := carbon.Now(carbon.PRC)
		code := utils.Captcha(6)
		item := redis_model.RedisSmsCode{
			Type:     smsType,
			SmsCode:  code,
			SendTime: carbon.DateTime{Carbon: now},
		}
		res, err1 := s.sendEmailSmsCode(email, smsType, code)
		if err1 != nil {
			return nil, err1
		}

		if err1 := s.setRedisSmsCode(redisKey, item); err1 != nil {
			global.GVA_LOG.Error(err1.Error())
		}

		return res, nil
	}

	// 已经存在验证码, 但是验证码没超过间隔发送时间
	restTime, err := global.GVA_REDIS.TTL(context.Background(), redisKey).Result()
	if err != nil {
		return nil, err
	}

	code := rSmsCode.SmsCode
	if global.GVA_CONFIG.System.Env == global.ENV_PORT { // 真实环境, 发送真实邮件
		code = "邮箱验证码已发送至邮箱"
	}

	sendRestSec := global.SmsCodeSendInterval - rSmsCode.SendTime.DiffAbsInSeconds(carbon.Now(carbon.PRC))
	responseStruct := response.GetSMSRes{
		IsAlreadySend: true,
		SmsCode:       code,
		SendRestSec:   sendRestSec,
		ExpireSec:     int64(restTime.Seconds()),
	}
	return &responseStruct, nil
}

func (s *BaseService) sendEmailSmsCode(email string, smsType string, code string) (*response.GetSMSRes, error) {
	if global.GVA_CONFIG.System.Env == global.ENV_PORT { // 真实环境, 发送真实邮件
		if err := s.sendRealEmailSmsCode(email, smsType, code, global.SmsCodeTimeOut/60); err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, errors.New("发送验证码失败, 请重试!")
		}

		responseStruct := response.GetSMSRes{
			IsAlreadySend: false,
			SmsCode:       "验证码已发送至邮箱",
			ExpireSec:     global.SmsCodeTimeOut,
			SendRestSec:   global.SmsCodeSendInterval,
		}
		return &responseStruct, nil
	}

	// 开发环境, 直接返回结果
	responseStruct := response.GetSMSRes{
		IsAlreadySend: false,
		SmsCode:       code,
		ExpireSec:     global.SmsCodeTimeOut,
		SendRestSec:   global.SmsCodeSendInterval,
	}
	return &responseStruct, nil
}

func (s *BaseService) sendRealEmailSmsCode(email string, smsType string, code string, codeTimeOut int64) error {
	template := `
	<div style="border-style: solid; border-width: thin; border-color:#dadce0; border-radius: 8px; padding: 30px 20px;">
		<div style="font-size: 24px;"> <b>%s</b> </div>
		<div style="font-size: 24px; padding-top: 15px; padding-bottom: 15px;"> %s </div>
		<div style="font-size: 20px; text-align: left;">
			<b>verification code %d minutes valid</b>
		</div>
	</div>
	`
	var subject, content string
	switch smsType {
	case global.SmsCodeLoginRegister:
		subject = "Login Or Register verification code"
	case global.SmsCodeForgot:
		subject = "Forgot password verification code" // content = fmt.Sprintf("<div><h2>%s</h2><h4>请使用以下验证码完成重制密码操作:</h4><h1>%s</h1><h4>验证码%d分钟内有效</h4></div>", subject, code, codeTimeOut)
	default:
		return errors.New("error verification code type not supported")
	}
	content = fmt.Sprintf(template, subject, code, codeTimeOut)
	global.GVA_LOG.Info(content)
	return utils.SendEmail("Quaqua", email, "[Tarot] Please verify your device", content)
}

// @author: [bill]
// @function: CheckEmailSmsCode
// @description: 验证邮箱验证码
// @param: email string, smsCode string, smsType string
// @return: *response.CheckSMSRes, error
func (s *BaseService) CheckEmailSmsCode(email string, smsCode string, smsType string) error {
	//短信存在？
	redisKey := s.newEmailRedisKey(smsType, email)
	rSmsCode, err := s.getRedisSmsCode(redisKey)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return errors.New("验证码错误")
	}

	if smsType != rSmsCode.Type || smsCode != rSmsCode.SmsCode {
		return errors.New("验证码错误")
	}

	// restTime, err := global.GVA_REDIS.TTL(context.Background(), redisKey).Result()
	// if err != nil {
	// 	global.GVA_LOG.Error(err.Error())
	// 	return errors.New("验证码错误")
	// }

	// responseStruct := response.CheckSMSRes{
	// 	IsExpired: false,
	// 	SmsCode:   rSmsCode.SmsCode,
	// 	ExpireSec: int(restTime.Seconds()),
	// }
	// return &responseStruct, nil

	return nil
}

// @author: [bill]
// @function: GetPhoneSmsCode
// @description: 获取手机短信验证码
// @param: phone string, smsType string
// @return: *response.GetSMSRes, error
func (s *BaseService) GetPhoneSmsCode(phone string, smsType string) (*response.GetSMSRes, error) {
	redisKey := s.newPhoneRedisKey(smsType, phone)
	// 查找手机是否已发送过验证码并且还未过期
	rSmsCode, err := s.getRedisSmsCode(redisKey)

	// 未找到,重新生成验证码并发送
	if err != nil {
		now := carbon.Now(carbon.PRC)
		code := utils.Captcha(6)
		item := redis_model.RedisSmsCode{
			Type:     smsType,
			SmsCode:  code,
			SendTime: carbon.DateTime{Carbon: now},
		}

		res, err1 := s.sendPhoneSmsCode(phone, smsType, code)
		if err1 != nil {
			return nil, err1
		}

		if err1 := s.setRedisSmsCode(redisKey, item); err1 != nil {
			global.GVA_LOG.Error(err1.Error())
		}

		return res, nil
	}

	// 找到
	if gte := carbon.Now(carbon.PRC).Gte(rSmsCode.SendTime.AddSeconds(int(global.SmsCodeSendInterval))); gte {
		// 已经存在验证码, 但是验证码超过间隔发送时间, 发送新的
		now := carbon.Now(carbon.PRC)
		code := utils.Captcha(6)
		item := redis_model.RedisSmsCode{
			Type:     smsType,
			SmsCode:  code,
			SendTime: carbon.DateTime{Carbon: now},
		}
		res, err1 := s.sendPhoneSmsCode(phone, smsType, code)
		if err1 != nil {
			return nil, err1
		}

		if err1 := s.setRedisSmsCode(redisKey, item); err1 != nil {
			global.GVA_LOG.Error(err1.Error())
		}

		return res, nil
	}

	// 已经存在验证码, 但是验证码没超过间隔发送时间
	restTime, err := global.GVA_REDIS.TTL(context.Background(), redisKey).Result()
	if err != nil {
		return nil, err
	}

	code := rSmsCode.SmsCode
	if global.GVA_CONFIG.System.Env == global.ENV_PORT { // 真实环境, 发送真实邮件
		code = "短信验证码已发送至手机"
	}

	sendRestSec := global.SmsCodeSendInterval - rSmsCode.SendTime.DiffAbsInSeconds(carbon.Now(carbon.PRC))
	responseStruct := response.GetSMSRes{
		IsAlreadySend: true,
		SmsCode:       code,
		SendRestSec:   sendRestSec,
		ExpireSec:     int64(restTime.Seconds()),
	}
	return &responseStruct, nil
}

func (s *BaseService) sendPhoneSmsCode(phone string, smsType string, code string) (*response.GetSMSRes, error) {
	// if global.GVA_CONFIG.System.Env == global.ENV_PORT { // 真实环境, 发送真实验证码
	// 	if err := s.sendRealEmailSmsCode(phone, smsType, code, global.SmsCodeTimeOut/60); err != nil {
	// 		global.GVA_LOG.Error(err.Error())
	// 		return nil, errors.New("发送验证码失败, 请重试!")
	// 	}

	// 	responseStruct := response.GetSMSRes{
	// 		IsAlreadySend: false,
	// 		SmsCode:       "验证码已发送至手机",
	// 		ExpireSec:     global.SmsCodeTimeOut,
	// 		SendRestSec:   global.SmsCodeSendInterval,
	// 	}
	// 	return &responseStruct, nil
	// }

	// 开发环境, 直接返回结果
	responseStruct := response.GetSMSRes{
		IsAlreadySend: false,
		SmsCode:       code,
		ExpireSec:     global.SmsCodeTimeOut,
		SendRestSec:   global.SmsCodeSendInterval,
	}
	return &responseStruct, nil
}

// @author: [bill]
// @function: CheckPhoneSmsCode
// @description: 验证手机验证码
// @param: phone string, smsCode string, smsType string
// @return: *response.CheckSMSRes, error
func (s *BaseService) CheckPhoneSmsCode(phone string, smsCode string, smsType string) error {
	//短信存在？
	redisKey := s.newPhoneRedisKey(smsType, phone)
	rSmsCode, err := s.getRedisSmsCode(redisKey)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return errors.New("验证码错误")
	}

	if smsType != rSmsCode.Type || smsCode != rSmsCode.SmsCode {
		return errors.New("验证码错误")
	}
	// restTime, err := global.GVA_REDIS.TTL(context.Background(), redisKey).Result()
	// if err != nil {
	// 	global.GVA_LOG.Error(err.Error())
	// 	return errors.New("验证码错误")
	// }

	// responseStruct := response.CheckSMSRes{
	// 	IsExpired: false,
	// 	SmsCode:   rSmsCode.SmsCode,
	// 	ExpireSec: int(restTime.Seconds()),
	// }
	// return &responseStruct, nil
	return nil
}

// @author: [bill]
// @function: setRedisSmsCode
// @description: 向redis设置验证码
// @param: email string, prefix string
// @return: *redis_model.RedisSmsCode, error
func (s *BaseService) setRedisSmsCode(key string, rSmsCode redis_model.RedisSmsCode) error {
	timer := time.Duration(global.SmsCodeTimeOut) * time.Second
	ub, err := json.Marshal(rSmsCode)
	if err != nil {
		return err
	}

	err = global.GVA_REDIS.Set(context.Background(), key, ub, timer).Err()
	return err
}

// @author: [bill]
// @function: getRedisSmsCode
// @description: 从redis获取已发送未过期的验证码
// @param: email string, prefix string
// @return: *redis_model.RedisSmsCode, error
func (s *BaseService) getRedisSmsCode(key string) (*redis_model.RedisSmsCode, error) {
	ub, err := global.GVA_REDIS.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	// 直接将　user set,再Unmarshal为结构体时会失败
	res := redis_model.RedisSmsCode{}
	err = json.Unmarshal([]byte(ub), &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *BaseService) newEmailRedisKey(prefix string, email string) string {
	return fmt.Sprintf("%s-%s-%s", prefix, "email", email)
}

func (s *BaseService) newPhoneRedisKey(prefix string, phone string) string {
	return fmt.Sprintf("%s-%s-%s", prefix, "phone", phone)
}

func (s *BaseService) LoginOrRegisterByEmailSMS(email string) (*response.LoginResponse, error) {
	// 验证码通过, 邮箱存在则登录, 否则则创建账号
	customer, err := business.BusCustomerDao.GetCustomerByEmail(global.GVA_DB, email)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, errors.New("查找邮箱失败")
	}

	if customer == nil { // 不存在, 注册账号
		customer = &business.BusCustomer{
			Email: email,
		}
		if err := global.GVA_DB.Create(&customer).Error; err != nil {
			return nil, errors.New("注册用户失败")
		}
	}

	res, err := s.LoginResponse(customer.ID)
	if err != nil || res == nil {
		return nil, errors.New("登录失败")
	}

	return res, nil
}

func (s *BaseService) LoginResponse(UserId uint) (*response.LoginResponse, error) {
	// //验证通过
	// j := &middleware.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	// nowAt := carbon.Now(carbon.PRC).TimestampWithSecond()
	// notBefore := nowAt - 1000
	// expiresAt := nowAt + global.Config.JWT.ExpiresTime

	// claims := request.CustomClaims{
	// 	ID:         UserId,
	// 	BufferTime: global.Config.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
	// 	StandardClaims: jwt.StandardClaims{
	// 		NotBefore: notBefore,  // 签名生效时间
	// 		ExpiresAt: expiresAt,  // 过期时间 7天  配置文件
	// 		Issuer:    "csgTarot", // 签名的发行者
	// 	},
	// }
	// //生成token
	// token, err := j.CreateToken(claims)
	// if err != nil {
	// 	global.Logger.Error("Login", zap.Any("err", err))
	// 	return nil, errors.New("create Token err")
	// }

	// tUser, err := tables.TUserDao.GetUserByID(global.DB, UserId)

	// if err != nil {
	// 	global.Logger.Error("GetUserByID", zap.Any("err", err))
	// 	return nil, err
	// }

	// //是否开多点登录
	// if global.Config.System.UseMultipoint {
	// 	res := response.LoginResponse{
	// 		UserId: UserId,
	// 		Token:  token,
	// 		// Email:    tUserPasswd.Email,
	// 		NickName: tUser.NickName,
	// 		Sex:      tUser.Sex,
	// 		Birth:    tUser.Birth,
	// 		//ExpiresAt: expiresAt,
	// 	}
	// 	return &res, nil
	// }

	// //存Redis数据库
	// if jwtStr, err := service.GetRedisJWT(UserId); err == redis.Nil {
	// 	global.Logger.Debug("Login ", zap.Any("GetRedisJWT == nil", UserId))
	// 	//Redis没有保存Token,就存Token
	// 	if err := service.SetRedisJWT(token, UserId); err != nil {
	// 		global.Logger.Error("Login", zap.Any("err", err))
	// 		return nil, err
	// 	}
	// 	res := response.LoginResponse{
	// 		UserId: UserId,
	// 		Token:  token,
	// 		// Email:    tUserPasswd.Email,
	// 		NickName: tUser.NickName,
	// 		Sex:      tUser.Sex,
	// 		Birth:    tUser.Birth,
	// 		//ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}
	// 	return &res, nil
	// } else if err != nil {
	// 	//Redis服务问题保存失败
	// 	global.Logger.Error("Login", zap.Any("err", err))
	// 	return nil, err
	// } else {
	// 	//Redis有Token记录，则添加到黑名单
	// 	var blackJWT tables.TJwtBlackList
	// 	blackJWT.Jwt = jwtStr
	// 	if err := service.JsonInBlacklist(blackJWT); err != nil {
	// 		global.Logger.Error("Login", zap.Any("err", err))
	// 		//return nil, errors.New("token作废失败")
	// 	}
	// 	global.Logger.Debug("Login ", zap.Any("jwtStr ", jwtStr))
	// 	//Redis有新的Token值
	// 	if err := service.SetRedisJWT(token, UserId); err != nil {
	// 		global.Logger.Error("Login", zap.Any("err", err))
	// 		return nil, err
	// 	}
	// 	res := response.LoginResponse{
	// 		UserId: UserId,
	// 		Token:  token,
	// 		// Email:    tUserPasswd.Email,
	// 		NickName: tUser.NickName,
	// 		Sex:      tUser.Sex,
	// 		Birth:    tUser.Birth,
	// 		//ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	// 	}
	// 	return &res, nil
	// }
	return nil, nil
}
