package business

import (
	"github.com/gin-gonic/gin"
)

type CustomerAssetRouter struct{}

func (s *CustomerAssetRouter) InitCustomerAssetPrivateRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("customer/asset")

	apiRouter.GET("/my", busCustomerAssetApi.GetCustomerAsset) // 获取用户资产信息
	apiRouter.GET("/bill", busCustomerAssetApi.GetCustomerAssetBill)
	apiRouter.POST("/deposit/alipay", busCustomerAssetApi.DepositByAlipay)
	apiRouter.POST("/deposit/wxpay", busCustomerAssetApi.DepositByWxpay)
	apiRouter.GET("/deposit/order/status", busCustomerAssetApi.GetDepositOrderStatus)
}

func (s *CustomerAssetRouter) InitCustomerAssetPublicRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("customer/asset")

	apiRouter.POST("/deposit/alipay/notify", busCustomerAssetApi.DepositByAlipayNotify)
	apiRouter.POST("/deposit/wxpay/notify", busCustomerAssetApi.DepositByWxpayNotify)
}
