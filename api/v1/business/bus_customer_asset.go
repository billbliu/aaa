package business

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	BusResponse "github.com/flipped-aurora/gin-vue-admin/server/model/business/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusCustomerAssetApi struct{}

// GetCustomerAsset
// @Summary 获取用户资产信息
// @Description 获取JSON
// @Tags 用户中心-资产
// @Success 200 {object} response.Response{data=BusResponse.CustomerAssetGetRes,msg=string} "用户资产"
// @Router /customer/asset/my [get]
// @Security  ApiKeyAuth
func (api BusCustomerAssetApi) GetCustomerAsset(c *gin.Context) {
	uid := utils.GetUserID(c)
	res, err := customerAssetService.GetCustomerAsset(uid)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}

	response.OkWithData(res, c)
}

// GetCustomerAssetBill
// @Summary 获取用户资产账单信息
// @Description 获取JSON
// @Tags 用户中心-资产
// @Param behavior query int false "行为 1:充值、2:会员购买"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=BusResponse.CustomerAssetGetRes,msg=string} "用户资产账单"
// @Router /customer/asset/bill [get]
// @Security  ApiKeyAuth
func (api BusCustomerAssetApi) GetCustomerAssetBill(c *gin.Context) {
	req := request.CustomerAssetBillGetReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)

	list, total, err := customerAssetService.GetCustomerAssetBill(uid, req)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)
}

// DepositByAlipay
// @Summary 资产充值-支付宝
// @Description 获取JSON
// @Tags 用户中心-资产
// @Accept  application/json
// @Product application/json
// @Param data body request.DepositByAlipayReq true "data"
// @Success 200 {object} response.Response{data=BusResponse.DepositByAlipayRes,msg=string} "支付宝充值"
// @Router /customer/asset/deposit/alipay [post]
// @Security  ApiKeyAuth
func (api BusCustomerAssetApi) DepositByAlipay(c *gin.Context) {
	req := request.DepositByAlipayReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)

	res, err := customerAssetService.DepositByAlipay(c, &req, uid)
	if err != nil {
		global.GVA_LOG.Error("支付宝充值失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("支付宝充值失败！错误详情：%s", err.Error()), c)
		return
	}
	response.OkWithData(res, c)
}

// DepositByAlipayNotify
// @Summary 资产充值-支付宝(回调通知)
// @Description 获取JSON
// @Tags 用户中心-资产
// @Accept  application/json
// @Product application/json
// @Success 200 {object}  response.Response{data=string,msg=string} "用户资产"
// @Router /customer/asset/deposit/alipay/notify [post]
func (api BusCustomerAssetApi) DepositByAlipayNotify(c *gin.Context) {
	res := customerAssetService.DepositByAlipayNotify(c)
	response.OkWithData(res, c)
}

// DepositByWxpay
// @Summary 资产充值-微信
// @Description 获取JSON
// @Tags 用户中心-资产
// @Accept  application/json
// @Product application/json
// @Param data body request.DepositByWxpayReq true "data"
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /customer/asset/deposit/wxpay [post]
// @Security  ApiKeyAuth
func (api BusCustomerAssetApi) DepositByWxpay(c *gin.Context) {
	req := request.DepositByWxpayReq{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uid := utils.GetUserID(c)

	res, err := customerAssetService.DepositByWxpay(c, &req, uid)
	if err != nil {
		global.GVA_LOG.Error("微信充值失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("微信充值失败！错误详情：%s", err.Error()), c)
		return
	}
	response.OkWithData(res, c)
}

// DepositByWxpayNotify
// @Summary 资产充值-微信(回调通知)
// @Description 获取JSON
// @Tags 用户中心-资产
// @Accept  application/json
// @Product application/json
// @Success 200 {object} response.Response "{"code": 200, "data": [...]}"
// @Router /customer/asset/deposit/wxpay/notify [post]
func (api BusCustomerAssetApi) DepositByWxpayNotify(c *gin.Context) {
	res, err := customerAssetService.DepositByWxpayNotify(c)
	if err != nil {
		global.GVA_LOG.Error("微信充值失败!", zap.Error(err))
		response.FailWithMessage(fmt.Sprintf("微信充值失败！错误详情：%s", err.Error()), c)
		return
	}
	response.OkWithData(res, c)
}

// GetDepositAlipayOrderStatus
// @Summary 资产充值-查询充值订单状态
// @Description 获取JSON
// @Tags 用户中心-资产
// @Accept  application/json
// @Product application/json
// @Param out_trade_no query string true "商家订单编号"
// @Success 200 {object} response.Response "{"code": 200, "data": ""}"
// @Router /customer/asset/deposit/order/status [get]
// @Security  ApiKeyAuth
func (api BusCustomerAssetApi) GetDepositOrderStatus(c *gin.Context) {
	req := request.DepositOrderStatusReq{}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	status := customerAssetService.GetDepositOrderStatus(c, &req)

	res := BusResponse.DepositOrderStatusRes{
		OutTradeNo: req.OutTradeNo,
		Status:     status,
	}
	response.OkWithData(res, c)
}
