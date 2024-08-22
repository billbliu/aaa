package business

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	WebsiteRouter
	WebsiteCategoryRouter
	ModeRouter
}

var (
	busWebsiteApi         = api.ApiGroupApp.BusinessApiGroup.BusWebsiteApi
	busWebsiteCategoryApi = api.ApiGroupApp.BusinessApiGroup.BusWebsiteCategoryApi
	busModeApi            = api.ApiGroupApp.BusinessApiGroup.BusModeApi
)
