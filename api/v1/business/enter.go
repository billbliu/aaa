package business

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BusWebsiteApi
	BusWebsiteCategoryApi
	BusModeApi
	BusBaseApi
	BusCustomerAssetApi
}

var (
	websiteService         = service.ServiceGroupApp.BusinessServiceGroup.WebsiteService
	websiteCategoryService = service.ServiceGroupApp.BusinessServiceGroup.WebsiteCategoryService
	modeService            = service.ServiceGroupApp.BusinessServiceGroup.ModeService
	baseService            = service.ServiceGroupApp.BusinessServiceGroup.BaseService
	customerAssetService   = service.ServiceGroupApp.BusinessServiceGroup.CustomerAssetService
)
