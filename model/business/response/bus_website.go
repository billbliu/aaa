package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/business"

type GetModeBindingCategoryWebsitesRes struct {
	Mode         business.BusMode           `json:"mode" comment:"模式"`
	CategoryList []ModeCategoryWebsitesList `json:"categoryList" comment:"模式分类网站列表"`
}

type ModeCategoryWebsitesList struct {
	Category    business.BusWebsiteCategory `json:"category" comment:"模式分类"`
	WebsiteList []business.BusWebsite       `json:"websiteList" comment:"网站列表"`
}
