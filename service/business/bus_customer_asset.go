package business

import (
	"encoding/json"
	"errors"
	"fmt"

	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business/response"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-pay/xlog"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm/clause"
)

type CustomerAssetService struct{}

var CustomerAssetServiceApp = new(CustomerAssetService)

var mu sync.Mutex

const UserAdminId = 1

const (
	// 支付宝交易状态
	ALI_TRADE_STATE_WAIT_BUYER_PAY = "WAIT_BUYER_PAY" // 交易创建，等待买家付款
	ALI_TRADE_STATE_TRADE_SUCCESS  = "TRADE_SUCCESS"  // 交易支付成功
	ALI_TRADE_STATE_TRADE_CLOSED   = "TRADE_CLOSED"   // 未付款交易超时关闭，或支付完成后全额退款
	ALI_TRADE_STATE_TRADE_FINISHED = "TRADE_FINISHED" // 交易结束，不可退款

	// 支付宝提现转账单据状态
	ALI_TRANS_STATE_SUCCESS = "SUCCESS" // 该笔转账交易成功
	ALI_TRANS_STATE_FAIL    = "FAIL"    // 失败（具体失败原因请参见error_code以及fail_reason返回值）
	ALI_TRANS_STATE_DEALING = "DEALING" // 处理中（适用于"单笔转账到银行卡"）
	ALI_TRANS_STATE_REFUND  = "FAIL"    // 退票（适用于"单笔转账到银行卡"）；

	// 微信交易状态
	WX_TRADE_STATE_SUCCESS   = "SUCCESS"   // 支付成功
	WX_TRADE_STATE_REFUND    = "REFUND"    // 转入退款
	WX_TRADE_STATE_NOTPAY    = "NOTPA"     // 未支付
	WX_TRADE_STATE_CLOSED    = "CLOSED"    // 已关闭
	WX_TRADE_STATE_REVOKED   = "REVOKED"   // 已撤销（仅付款码支付会返回）
	WX_TRADE_STATE_USERPYING = "USERPYING" //用户支付中（仅付款码支付会返回）
	WX_TRADE_STATE_PAYERROR  = "PAYERROR"  // 支付失败（仅付款码支付会返回
)

// GetCustomerAsset 获取用户资产
func (ser *CustomerAssetService) GetCustomerAsset(userId uint) (*response.CustomerAssetGetRes, error) {
	customerAsset, err := business.BusCustomerAssetDao.GetCustomerAssetByID(global.GVA_DB, userId)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	res := response.CustomerAssetGetRes{
		Total:  customerAsset.Total,
		Freeze: customerAsset.Freeze,
		Avail:  customerAsset.Total.Sub(customerAsset.Freeze),
	}
	return &res, nil
}

// GetCustomerAssetBill 获取用户资产账单
func (ser *CustomerAssetService) GetCustomerAssetBill(userId uint, req request.CustomerAssetBillGetReq) (*[]response.CustomerAssetBillGetRes, int64, error) {
	bills := []business.BusCustomerAssetBill{}

	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)

	orm := global.GVA_DB.Debug().Model(&business.BusCustomerAssetBill{})

	if userId > 0 {
		orm = orm.Where("user_id = ?", userId)
	}

	if req.Behavior == 1 || req.Behavior == 2 || req.Behavior == 3 {
		orm = orm.Where("behavior = ?", req.Behavior)
	}

	var total int64

	if err := orm.Count(&total).Error; err != nil {
		return nil, total, err
	}

	err := orm.Limit(limit).Offset(offset).Order("created_at desc").Find(&bills).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, total, err
	}

	l := len(bills)

	if l > 0 {
		res := make([]response.CustomerAssetBillGetRes, l)
		for i := 0; i < l; i++ {
			behavior := ""
			if bills[i].Behavior == business.BEHAVIOR_DEPOSIT {
				behavior = "充值"
			} else if bills[i].Behavior == business.BEHAVIOR_MEMBER_PURCHASE {
				behavior = "会员购买"
			}
			payDirection := ""
			if bills[i].PayDirection == business.PAY_DIRECTION_ADD {
				payDirection = "+"
			} else if bills[i].PayDirection == business.PAY_DIRECTION_SUB {
				payDirection = "-"
			}
			res[i].Before = bills[i].Before.String()
			res[i].Amount = fmt.Sprintf("%s%s", payDirection, bills[i].Amount.String())
			res[i].After = bills[i].After.String()
			res[i].Behavior = behavior
			res[i].CreatedAt = bills[i].CreatedAt
		}
		return &res, total, nil
	}
	return nil, total, nil
}

// MallOrderPay 商城订单支付
func (ser *CustomerAssetService) MallOrderPay(formUserId uint, toUserId uint, amount float64, mallOrderId int64, orderType business.OrderType) (bool, error) {
	money := decimal.NewFromFloat(amount)

	// 订单完成，结算
	if err := ser.mallOrderClearing(formUserId, toUserId, uint(mallOrderId), money, orderType); err != nil {
		global.GVA_LOG.Error(err.Error())
		return false, err
	}
	return true, nil
}

// mallOrderClearing 商城订单结算
// params:
// formUserId 资产支出账户
// toUserId 资产流入账户
// orderID 商城订单ID
// orderAmount 订单金额
func (ser *CustomerAssetService) mallOrderClearing(fromUserId uint, toUserId uint, orderID uint, orderAmount decimal.Decimal, orderType business.OrderType) error {
	fromCustomerAsset, err := business.BusCustomerAssetDao.GetCustomerAssetByID(global.GVA_DB, fromUserId)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return err
	}

	if fromCustomerAsset.Total.Sub(fromCustomerAsset.Freeze).LessThan(orderAmount) {
		return errors.New("可用余额不足!")
	}

	toCustomerAsset, err := business.BusCustomerAssetDao.GetCustomerAssetByID(global.GVA_DB, toUserId)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return err
	}

	tx := global.GVA_DB.Begin()
	// 1 记录账单
	fromUserBill := business.BusCustomerAssetBill{
		Before:        fromCustomerAsset.Total,
		Amount:        orderAmount,
		After:         fromCustomerAsset.Total.Sub(orderAmount),
		Behavior:      business.BEHAVIOR_MEMBER_PURCHASE,
		PayDirection:  business.PAY_DIRECTION_SUB,
		BusCustomerId: fromCustomerAsset.BusCustomerId,
		OrderId:       orderID,
		OrderType:     orderType,
	}
	if err := tx.Create(&fromUserBill).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return err
	}

	toUserBill := business.BusCustomerAssetBill{
		Before:        toCustomerAsset.Total,
		Amount:        orderAmount,
		After:         toCustomerAsset.Total.Add(orderAmount),
		Behavior:      business.BEHAVIOR_MEMBER_PURCHASE,
		PayDirection:  business.PAY_DIRECTION_ADD,
		BusCustomerId: toCustomerAsset.BusCustomerId,
		OrderId:       orderID,
		OrderType:     orderType,
	}

	if err := tx.Create(&toUserBill).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return err
	}

	// 2 修改资产 - 商城算力交易使用, 目前不走冻结流程, 直接结算
	if err := tx.Model(&fromCustomerAsset).Updates(business.BusCustomerAsset{Total: fromCustomerAsset.Total.Sub(orderAmount)}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return err
	}

	if err := tx.Model(&toCustomerAsset).Updates(business.BusCustomerAsset{Total: toCustomerAsset.Total.Add(orderAmount)}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// 官方文档：https://opendocs.alipay.com/open/203/105288?pathHash=f2308e24
// 帮助中心：https://opensupport.alipay.com/support/FAQ
// DepositByAlipay 用户通过支付宝充值
func (ser *CustomerAssetService) DepositByAlipay(ctx *gin.Context, req *request.DepositByAlipayReq, customerId uint) (*response.DepositByAlipayRes, error) {
	// 查找订单是否存在
	orderInfo, err := business.BusCustomerDepositOrderDao.GetCustomerDepositOrderByUuid(global.GVA_DB, req.Uuid)
	if err == nil && orderInfo.Status == business.PAY_STATUS_NOPAY { // 已经存在订单并且未完成支付
		res := response.DepositByAlipayRes{
			OutTradeNo: orderInfo.OrderNo,
			PayUrl:     orderInfo.PayUrl,
		}
		return &res, nil
	}

	// 生成商家订单号
	tradeNo := strings.ReplaceAll(uuid.New().String(), "-", "")
	log.Println("tradeNo: ", tradeNo)

	// 不存在订单，生成新订单
	newOrder := business.BusCustomerDepositOrder{
		Uuid:          req.Uuid,
		OrderNo:       tradeNo,
		PaymentType:   business.PAY_ALIPAY,
		BusCustomerId: customerId,
		Amount:        decimal.NewFromFloat(req.Money),
		PayUrl:        "",
		Status:        business.PAY_STATUS_NOPAY,
	}
	if err := global.GVA_DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uuid"}},
		UpdateAll: true,
	}).Create(&newOrder).Error; err != nil {
		return nil, err
	}

	loc, _ := time.LoadLocation("UTC")

	expire := time.Now().In(loc).Add(490 * time.Minute).Format(time.DateTime)
	// 构造支付参数
	bm := make(gopay.BodyMap)
	bm.Set("subject", fmt.Sprintf("%s-%s", global.GVA_CONFIG.PayPlatform.Title, "充值")).
		Set("out_trade_no", tradeNo).
		Set("total_amount", req.Money).
		Set("time_expire", expire).
		Set("product_code", "QUICK_WAP_WAY").
		Set("quit_url", "https://www.google.com")
	// Set("product_code", "FAST_INSTANT_TRADE_PAY")
	// Set("notify_url", fmt.Sprintf("%s%s", global.GVA_CONFIG.PayPlatform.Alipay.NotifyUrl, "/api/v1/mall/asset/deposit/alipay/notify"))
	// Set("return_url", fmt.Sprintf("%s%s", global.GVA_CONFIG.PayPlatform.Alipay.ReturnUrl, "/api/v1/mall/asset/deposit/alipay/notify"))

	// 发起支付
	payUrl, err := global.PaymentClient.Alipay.TradeWapPay(ctx, bm)
	if err != nil {
		xlog.Errorf("client.TradeWapPay(%+v),error:%+v", bm, err)
		return nil, err
	}

	// 更新付款链接
	if err := global.GVA_DB.Model(&newOrder).Update("pay_url", payUrl).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	// return
	res := response.DepositByAlipayRes{
		OutTradeNo: tradeNo,
		PayUrl:     payUrl,
	}
	return &res, nil
}

// DepositByAlipayNotify 用户通过支付宝充值回调通知
// 只有TRADE_SUCCESS 支付成功才会触发这个通知
func (ser *CustomerAssetService) DepositByAlipayNotify(ctx *gin.Context) string {
	// 解析请求参数
	bm, err := alipay.ParseNotifyToBodyMap(ctx.Request)
	if err != nil {
		log.Fatalln("err:", err)
		return "fail"
	}
	log.Println("notifyReq:", bm)

	// 验签
	ok, err := alipay.VerifySign(global.GVA_CONFIG.PayPlatform.Alipay.AlipayPublicCertContent, bm)
	if err != nil {
		xlog.Error("err:", err)
		return "fail"
	}
	log.Println("支付宝验签是否通过:", ok)

	if ok {
		outTradeNo := bm.Get("out_trade_no")  // 原支付请求的商户订单号
		tradeNo := bm.Get("trade_no")         // 支付宝交易凭证号
		tradeStatus := bm.Get("trade_status") // 交易目前所处的状态

		totalAmount, err := strconv.ParseFloat(bm.Get("total_amount"), 32) // 订单金额
		if err != nil {
			global.GVA_LOG.Error(err.Error())
		}

		tx := global.GVA_DB.Begin()

		// 更新订单状态
		order, err := business.BusCustomerDepositOrderDao.UpdateCustomerDepositOrderStatusByOrderNo(tx, outTradeNo, business.PAY_STATUS_SUCCESS)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			tx.Rollback()
			return "fail"
		}

		// 记录支付日志
		paymentInfo := business.BusCustomerDepositPayment{
			OrderNo:       outTradeNo,
			PaymentType:   business.PAY_ALIPAY,
			TransactionId: tradeNo,
			TradeType:     "支付码支付",
			TradeState:    tradeStatus,
			Amount:        decimal.NewFromFloat(totalAmount),
			Content:       bm.JsonBody(),
		}
		err = tx.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "order_no"}},
			UpdateAll: true,
		}).Create(&paymentInfo).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			tx.Rollback()
			return "fail"
		}

		tx.Commit()

		// 订单完成，结算
		isClear, _ := business.BusCustomerAssetBillDao.IsClearingByOrderIdAndBehavior(global.GVA_DB, order.ID, business.BEHAVIOR_DEPOSIT)
		if !isClear { // 未结算
			if err = ser.depositOrderClearing(order.BusCustomerId, order.ID, order.Amount); err != nil {
				global.GVA_LOG.Error(err.Error())
				return "fail"
			}
		}

		return "success"
	} else {
		return "fail"
	}

}

// DepositByWxpay 用户通过微信充值
func (ser *CustomerAssetService) DepositByWxpay(ctx *gin.Context, req *request.DepositByWxpayReq, customerId uint) (*response.DepositByWxpayRes, error) {
	// 订单金额不能超过限制
	if req.Money <= 0 {
		return nil, errors.New("订单金额必须大于0")
	}

	if req.Money > 100000000 {
		return nil, errors.New("订单金额超过限制, 最大100000000!")
	}

	// 查找订单是否存在
	orderInfo, err := business.BusCustomerDepositOrderDao.GetCustomerDepositOrderByUuid(global.GVA_DB, req.Uuid)
	if err == nil && orderInfo.Status == business.PAY_STATUS_NOPAY { // 已经存在订单并且未完成支付
		res := response.DepositByWxpayRes{
			OutTradeNo: orderInfo.OrderNo,
			QrCode:     orderInfo.PayUrl,
		}
		return &res, nil
	}

	// 生成商家订单号
	tradeNo := strings.ReplaceAll(uuid.New().String(), "-", "")
	xlog.Infof("tradeNo: %s", tradeNo)

	// 不存在订单，生成新订单
	newOrder := business.BusCustomerDepositOrder{
		Uuid:          req.Uuid,
		OrderNo:       tradeNo,
		PaymentType:   business.PAY_WECHAT,
		BusCustomerId: customerId,
		Amount:        decimal.NewFromFloat(req.Money),
		PayUrl:        "",
		Status:        business.PAY_STATUS_NOPAY,
	}
	if err := global.GVA_DB.Create(&newOrder).Error; err != nil {
		return nil, err
	}

	// 构造支付参数
	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
	bm := make(gopay.BodyMap)
	bm.Set("appid", global.GVA_CONFIG.PayPlatform.Wechat.Appid).
		Set("description", fmt.Sprintf("%s-%s", global.GVA_CONFIG.PayPlatform.Title, "充值")).
		Set("out_trade_no", tradeNo).
		Set("time_expire", expire).
		Set("notify_url", fmt.Sprintf("%s%s", global.GVA_CONFIG.System.Domain, "/customer/asset/deposit/wxpay/notify")).
		SetBodyMap("amount", func(b gopay.BodyMap) {
			b.Set("total", req.Money).Set("currency", "CNY")
		}).
		SetBodyMap("scene_info", func(b gopay.BodyMap) {
			b.SetBodyMap("h5_info", func(b gopay.BodyMap) {
				b.Set("type", "Wap").
					Set("app_name", "王者荣耀").
					Set("app_url", "https://pay.qq.com").
					Set("bundle_id", "com.tencent.wzryiOS")
			})
		})
	// 发起支付
	wxRsp, err := global.PaymentClient.Wxpay.V3TransactionH5(ctx, bm)
	xlog.Error("wxRsp:", wxRsp)
	if err != nil {
		xlog.Error(err)
		return nil, err
	}

	if wxRsp.Code != wechat.Success {
		return nil, errors.New("调用微信支付接口失败")
	}

	// 更新付款链接
	if err := global.GVA_DB.Model(&newOrder).Update("pay_url", wxRsp.Response.H5Url).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
	}

	// 总账号冻结订单金额

	// return
	res := response.DepositByWxpayRes{
		OutTradeNo: tradeNo,
		QrCode:     wxRsp.Response.H5Url,
	}
	return &res, nil
}

// DepositByWxpay 用户通过微信充值
// func (ser *CustomerAssetService) DepositByWxpay(ctx *gin.Context, req *request.DepositByWxpayReq, customerId uint) (*response.DepositByWxpayRes, error) {
// 	// 订单金额不能超过限制
// 	if req.Money <= 0 {
// 		return nil, errors.New("订单金额必须大于0")
// 	}

// 	if req.Money > 100000000 {
// 		return nil, errors.New("订单金额超过限制, 最大100000000!")
// 	}

// 	// 查找订单是否存在
// 	orderInfo, err := business.BusCustomerDepositOrderDao.GetCustomerDepositOrderByUuid(global.GVA_DB, req.Uuid)
// 	if err == nil && orderInfo.Status == business.PAY_STATUS_NOPAY { // 已经存在订单并且未完成支付
// 		res := response.DepositByWxpayRes{
// 			OutTradeNo: orderInfo.OrderNo,
// 			QrCode:     orderInfo.PayUrl,
// 		}
// 		return &res, nil
// 	}

// 	// 生成商家订单号
// 	tradeNo := strings.ReplaceAll(uuid.New().String(), "-", "")
// 	xlog.Infof("tradeNo: %s", tradeNo)

// 	// 不存在订单，生成新订单
// 	newOrder := business.BusCustomerDepositOrder{
// 		Uuid:          req.Uuid,
// 		OrderNo:       tradeNo,
// 		PaymentType:   business.PAY_WECHAT,
// 		BusCustomerId: customerId,
// 		Amount:        decimal.NewFromFloat(req.Money),
// 		PayUrl:        "",
// 		Status:        business.PAY_STATUS_NOPAY,
// 	}
// 	if err := global.GVA_DB.Create(&newOrder).Error; err != nil {
// 		return nil, err
// 	}

// 	// 构造支付参数
// 	expire := time.Now().Add(10 * time.Minute).Format(time.RFC3339)
// 	bm := make(gopay.BodyMap)
// 	bm.Set("appid", global.GVA_CONFIG.PayPlatform.Wechat.Appid).
// 		Set("description", fmt.Sprintf("%s-%s", global.GVA_CONFIG.PayPlatform.Title, "充值")).
// 		Set("out_trade_no", tradeNo).
// 		Set("time_expire", expire).
// 		Set("notify_url", fmt.Sprintf("%s%s", global.GVA_CONFIG.System.Domain, "/customer/asset/deposit/wxpay/notify")).
// 		SetBodyMap("amount", func(bm gopay.BodyMap) {
// 			bm.Set("total", req.Money).Set("currency", "CNY")
// 		})
// 	// 发起支付
// 	wxRsp, err := global.PaymentClient.Wxpay.V3TransactionNative(ctx, bm)
// 	xlog.Error("wxRsp:", wxRsp)
// 	if err != nil {
// 		xlog.Error(err)
// 		return nil, err
// 	}

// 	if wxRsp.Code != wechat.Success {
// 		return nil, errors.New("调用微信支付接口失败")
// 	}

// 	// 更新付款链接
// 	if err := global.GVA_DB.Model(&newOrder).Update("pay_url", wxRsp.Response.CodeUrl).Error; err != nil {
// 		global.GVA_LOG.Error(err.Error())
// 	}

// 	// 总账号冻结订单金额

// 	// return
// 	res := response.DepositByWxpayRes{
// 		OutTradeNo: tradeNo,
// 		QrCode:     wxRsp.Response.CodeUrl,
// 	}
// 	return &res, nil
// }

// DepositByWxpayNotify 用户通过微信充值回调通知
// 只有TRADE_SUCCESS 支付成功才会触发这个通知
func (ser *CustomerAssetService) DepositByWxpayNotify(ctx *gin.Context) (*wechat.V3NotifyRsp, error) {
	// 解析参数
	v3NotifuReq, err := wechat.V3ParseNotify(ctx.Request)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return nil, err
	}

	// 验证签名
	if err = v3NotifuReq.VerifySignByPKMap(global.PaymentClient.Wxpay.WxPublicKeyMap()); err != nil {
		log.Println("err1:", err)
		return nil, err
	}

	// 解密加密信息
	res, err := v3NotifuReq.DecryptPayCipherText(global.GVA_CONFIG.PayPlatform.Wechat.ApiV3Key)
	if err != nil {
		log.Println("err2:", err)
		return nil, err
	}
	log.Println("res:", res)
	content, _ := json.Marshal(res)

	tx := global.GVA_DB.Begin()

	// 更新订单状态
	order, err := business.BusCustomerDepositOrderDao.UpdateCustomerDepositOrderStatusByOrderNo(tx, res.OutTradeNo, business.PAY_STATUS_SUCCESS)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return nil, err
	}

	// 记录支付日志
	paymentInfo := business.BusCustomerDepositPayment{
		OrderNo:       res.OutTradeNo,
		PaymentType:   business.PAY_WECHAT,
		TransactionId: res.TransactionId,
		TradeType:     "NATIVE支付",
		TradeState:    res.TradeState,
		Amount:        decimal.NewFromInt(int64(res.Amount.PayerTotal)),
		Content:       string(content),
	}
	err = global.GVA_DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "order_no"}},
		UpdateAll: true,
	}).Create(&paymentInfo).Error
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	// 订单完成，结算
	isClear, _ := business.BusCustomerAssetBillDao.IsClearingByOrderIdAndBehavior(global.GVA_DB, order.ID, business.BEHAVIOR_DEPOSIT)
	if !isClear { // 未结算
		if err = ser.depositOrderClearing(order.BusCustomerId, order.ID, order.Amount); err != nil {
			global.GVA_LOG.Error(err.Error())
			return nil, err
		}
	}

	rsp := new(wechat.V3NotifyRsp)
	rsp.Code = gopay.SUCCESS
	rsp.Message = "OK"
	return rsp, nil
}

// depositOrderClearing 用户充值结算 更新用户资产和资产账单记录
// params:
// userId 充值账户
// orderID 充值订单表ID
// orderAmount 订单金额
func (ser *CustomerAssetService) depositOrderClearing(busCustomerId uint, orderID uint, orderAmount decimal.Decimal) error {
	CustomerAsset, err := business.BusCustomerAssetDao.GetCustomerAssetByID(global.GVA_DB, busCustomerId)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return err
	}

	tx := global.GVA_DB.Begin()
	// 1 记录账单
	userBill := business.BusCustomerAssetBill{
		Before:        CustomerAsset.Total,
		Amount:        orderAmount,
		After:         CustomerAsset.Total.Add(orderAmount),
		Behavior:      business.BEHAVIOR_DEPOSIT,
		PayDirection:  business.PAY_DIRECTION_ADD,
		BusCustomerId: busCustomerId,
		OrderId:       orderID,
		OrderType:     business.ORDER_TYPE_DEPOSIT,
	}

	if err := tx.Create(&userBill).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return err
	}

	// 2 资产添加充值金额
	if err := tx.Model(&CustomerAsset).Updates(business.BusCustomerAsset{Total: CustomerAsset.Total.Add(orderAmount)}).Error; err != nil {
		global.GVA_LOG.Error(err.Error())
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

// GetDepositWxpayOrderStatus 获取充值订单状态
// 微信订单查询文档：https://pay.weixin.qq.com/docs/merchant/apis/native-payment/query-by-out-trade-no.html
func (ser *CustomerAssetService) GetDepositOrderStatus(ctx *gin.Context, req *request.DepositOrderStatusReq) business.DepositStatusType {
	// 查找订单是否存在
	orderInfo, err := business.BusCustomerDepositOrderDao.GetCustomerDepositOrderByOrderNo(global.GVA_DB, req.OutTradeNo)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return business.PAY_STATUS_NOPAY
	}

	// 订单未支付状态，查询订单状态接口
	if orderInfo.Status == business.PAY_STATUS_NOPAY {
		if orderInfo.PaymentType == business.PAY_ALIPAY {
			return ser.getDepositAlipayOrderStatus(ctx, orderInfo)
		} else if orderInfo.PaymentType == business.PAY_WECHAT {
			return ser.getDepositWxpayOrderStatus(ctx, orderInfo)
		}
	}

	// 其他状态，直接返回
	return orderInfo.Status
}

// getDepositAlipayOrderStatus 获取支付宝充值订单状态
// 支付宝订单查询接口：https://opendocs.alipay.com/open/82ea786a_alipay.trade.query?scene=23&pathHash=0745ecea
func (ser *CustomerAssetService) getDepositAlipayOrderStatus(ctx *gin.Context, orderInfo *business.BusCustomerDepositOrder) business.DepositStatusType {
	// 查询支付宝订单
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", orderInfo.OrderNo)
	aliRsp, err := global.PaymentClient.Alipay.TradeQuery(ctx, bm)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return business.PAY_STATUS_NOPAY
	}

	log.Println("res:", aliRsp.Response)

	if aliRsp.Response.TradeStatus == ALI_TRADE_STATE_WAIT_BUYER_PAY { // 等待支付状态
		return orderInfo.Status
	} else if aliRsp.Response.TradeStatus == ALI_TRADE_STATE_TRADE_CLOSED { // 交易关闭状态
		// 更新订单状态
		order, err := business.BusCustomerDepositOrderDao.UpdateCustomerDepositOrderStatusByOrderNo(global.GVA_DB, orderInfo.OrderNo, business.PAY_STATUS_CLOSED)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return business.PAY_STATUS_NOPAY
		}
		return order.Status
	} else if aliRsp.Response.TradeStatus == ALI_TRADE_STATE_TRADE_SUCCESS || aliRsp.Response.TradeStatus == ALI_TRADE_STATE_TRADE_FINISHED { // 交易支付成功或者交易结束(不可退款)
		status, err := business.BusCustomerDepositOrderDao.GetCustomerDepositOrderStatusByOrderNo(global.GVA_DB, orderInfo.OrderNo)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			return business.PAY_STATUS_NOPAY
		}

		// 已经更新了状态的直接返回
		if status != business.PAY_STATUS_NOPAY {
			return status
		}
		mu.Lock()
		defer mu.Unlock()
		tx := global.GVA_DB.Begin()
		// 更新订单状态
		order, err := business.BusCustomerDepositOrderDao.UpdateCustomerDepositOrderStatusByOrderNo(tx, orderInfo.OrderNo, business.PAY_STATUS_SUCCESS)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			tx.Rollback()
			return business.PAY_STATUS_NOPAY
		}

		totalAmount, err := strconv.ParseFloat(aliRsp.Response.TotalAmount, 64)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			tx.Rollback()
			return business.PAY_STATUS_NOPAY
		}

		content, _ := json.Marshal(aliRsp.Response)

		// 记录支付日志
		paymentInfo := business.BusCustomerDepositPayment{
			OrderNo:       aliRsp.Response.OutTradeNo,
			PaymentType:   business.PAY_ALIPAY,
			TransactionId: aliRsp.Response.TradeNo,
			TradeType:     "支付码支付",
			TradeState:    aliRsp.Response.TradeStatus,
			Amount:        decimal.NewFromFloat(totalAmount),
			Content:       string(content),
		}
		err = global.GVA_DB.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "order_no"}},
			UpdateAll: true,
		}).Create(&paymentInfo).Error
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			tx.Rollback()
			return business.PAY_STATUS_NOPAY
		}

		tx.Commit()

		// 订单完成，结算
		isClear, _ := business.BusCustomerAssetBillDao.IsClearingByOrderIdAndBehavior(global.GVA_DB, order.ID, business.BEHAVIOR_DEPOSIT)
		if !isClear { // 未结算
			if err = ser.depositOrderClearing(order.BusCustomerId, order.ID, order.Amount); err != nil {
				global.GVA_LOG.Error(err.Error())
				return business.PAY_STATUS_NOPAY
			}
		}
		return order.Status
	}
	return business.PAY_STATUS_NOPAY
}

// getDepositWxpayOrderStatus 获取微信充值订单状态
// 微信订单查询文档：https://pay.weixin.qq.com/docs/merchant/apis/native-payment/query-by-out-trade-no.html
func (ser *CustomerAssetService) getDepositWxpayOrderStatus(ctx *gin.Context, orderInfo *business.BusCustomerDepositOrder) business.DepositStatusType {
	//查询订单
	wxRsp, err := global.PaymentClient.Wxpay.V3TransactionQueryOrder(ctx, wechat.OutTradeNo, orderInfo.OrderNo)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return business.PAY_STATUS_NOPAY
	}

	log.Println("res:", wxRsp.Response)

	// 调用成功
	if wxRsp.Code == wechat.Success {
		if wxRsp.Response.TradeState == WX_TRADE_STATE_NOTPAY { // 等待支付状态
			return orderInfo.Status
		} else if wxRsp.Response.TradeState == WX_TRADE_STATE_CLOSED { // 交易关闭状态
			// 更新订单状态
			order, err := business.BusCustomerDepositOrderDao.UpdateCustomerDepositOrderStatusByOrderNo(global.GVA_DB, orderInfo.OrderNo, business.PAY_STATUS_CLOSED)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				return business.PAY_STATUS_NOPAY
			}
			return order.Status
		} else if wxRsp.Response.TradeState == WX_TRADE_STATE_SUCCESS { // 交易支付成功
			status, err := business.BusCustomerDepositOrderDao.GetCustomerDepositOrderStatusByOrderNo(global.GVA_DB, orderInfo.OrderNo)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				return business.PAY_STATUS_NOPAY
			}

			// 已经更新了状态的直接返回
			if status != business.PAY_STATUS_NOPAY {
				return status
			}
			mu.Lock()
			defer mu.Unlock()
			tx := global.GVA_DB.Begin()
			// 更新订单状态
			order, err := business.BusCustomerDepositOrderDao.UpdateCustomerDepositOrderStatusByOrderNo(tx, orderInfo.OrderNo, business.PAY_STATUS_SUCCESS)
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				tx.Rollback()
				return business.PAY_STATUS_NOPAY
			}

			content, _ := json.Marshal(wxRsp.Response)

			// 记录支付日志
			paymentInfo := business.BusCustomerDepositPayment{
				OrderNo:       wxRsp.Response.OutTradeNo,
				PaymentType:   business.PAY_WECHAT,
				TransactionId: wxRsp.Response.TransactionId,
				TradeType:     "NATIVE支付",
				TradeState:    wxRsp.Response.TradeState,
				Amount:        decimal.NewFromInt(int64(wxRsp.Response.Amount.Total)),
				Content:       string(content),
			}
			err = global.GVA_DB.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "order_no"}},
				UpdateAll: true,
			}).Create(&paymentInfo).Error
			if err != nil {
				global.GVA_LOG.Error(err.Error())
				tx.Rollback()
				return business.PAY_STATUS_NOPAY
			}

			tx.Commit()

			// 订单完成，结算
			isClear, _ := business.BusCustomerAssetBillDao.IsClearingByOrderIdAndBehavior(global.GVA_DB, order.ID, business.BEHAVIOR_DEPOSIT)
			if !isClear { // 未结算
				if err = ser.depositOrderClearing(order.BusCustomerId, order.ID, order.Amount); err != nil {
					global.GVA_LOG.Error(err.Error())
					return business.PAY_STATUS_NOPAY
				}
			}
			return order.Status
		}
	}

	return business.PAY_STATUS_NOPAY
}
