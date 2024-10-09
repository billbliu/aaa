package business

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	WebsiteRouter
	WebsiteCategoryRouter
	ModeRouter
	BaseRouter
	CustomerAssetRouter
	AdRouter
}

var (
	busWebsiteApi         = api.ApiGroupApp.BusinessApiGroup.BusWebsiteApi
	busWebsiteCategoryApi = api.ApiGroupApp.BusinessApiGroup.BusWebsiteCategoryApi
	busModeApi            = api.ApiGroupApp.BusinessApiGroup.BusModeApi
	busBaseApi            = api.ApiGroupApp.BusinessApiGroup.BusBaseApi
	busCustomerAssetApi   = api.ApiGroupApp.BusinessApiGroup.BusCustomerAssetApi
	busAdApi              = api.ApiGroupApp.BusinessApiGroup.BusAdApi
)
