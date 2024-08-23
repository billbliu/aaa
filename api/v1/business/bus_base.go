package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusBaseApi struct{}

// GetEmailSMS
// @Tags      BusBaseApi
// @Summary   获取邮箱验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.GetEmailSMSReq    true  "邮箱和验证码类型"
// @Success   200   {object}  response.Response{data=response.GetSMSRes}  "获取邮箱验证码"
// @Router    /base/getEmailSMS [get]
func (bus *BusBaseApi) GetEmailSMS(c *gin.Context) {
	var req businessReq.GetEmailSMSReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := baseService.GetEmailSmsCode(req.Email, req.SmsType)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(res, c)
}

// LoginOrRegisterByEmailSMS
// @Tags      BusBaseApi
// @Summary   通过邮箱登录或注册
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.LoginOrRegisterByEmailSMSReq        true  "邮箱验证码参数"
// @Success   200   {object}  response.Response{data=response.CheckSMSRes}    "通过邮箱登录或注册"
// @Router    /base/loginOrRegisterByEmailSMS [post]
func (s *BusBaseApi) LoginOrRegisterByEmailSMS(c *gin.Context) {
	var req businessReq.LoginOrRegisterByEmailSMSReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := baseService.CheckEmailSmsCode(req.Email, req.SmsCode, global.SmsCodeLoginRegister); err != nil {
		global.GVA_LOG.Error("检查验证码失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := baseService.LoginOrRegisterByEmailSMS(req.Email)
	if err != nil {
		global.GVA_LOG.Error("登录或注册失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(res, c)
}

// GetPhoneSMS
// @Tags      BusBaseApi
// @Summary   获取手机验证码
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.GetPhoneSMSReq    true  "手机和验证码类型"
// @Success   200   {object}  response.Response{data=response.GetSMSRes}  "获取手机验证码"
// @Router    /base/GetPhoneSMS [get]
func (bus *BusBaseApi) GetPhoneSMS(c *gin.Context) {
	var req businessReq.GetPhoneSMSReq
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := baseService.GetPhoneSmsCode(req.Phone, req.SmsType)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithData(res, c)
}

// LoginOrRegisterByPhoneSMS
// @Tags      BusBaseApi
// @Summary   通过手机登录或注册
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.LoginOrRegisterByPhoneSMSReq        true  "手机验证码"
// @Success   200   {object}  response.Response{data=response.CheckSMSRes}    "检查手机验证码"
// @Router    /base/loginOrRegisterByPhoneSMS [post]
func (s *BusBaseApi) LoginOrRegisterByPhoneSMS(c *gin.Context) {
	var req businessReq.LoginOrRegisterByPhoneSMSReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := baseService.CheckPhoneSmsCode(req.Phone, req.SmsCode, global.SmsCodeLoginRegister); err != nil {
		global.GVA_LOG.Error("检查验证码失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := baseService.LoginOrRegisterByEmailSMS(req.Phone)
	if err != nil {
		global.GVA_LOG.Error("登录或注册失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithData(res, c)
}
