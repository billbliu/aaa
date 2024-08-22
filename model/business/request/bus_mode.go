package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CreateModeReq struct {
	Name string `json:"name" comment:"模式名"`
}

type UpdateModeReq struct {
	ID   uint   `json:"id" comment:"网站ID"`
	Name string `json:"name" comment:"模式名"`
}

// Mode分页条件查询及排序结构体
type SearchModeParams struct {
	// business.BusMode
	request.PageInfo
}

type ModeBindingCategoryWebsitesReq struct {
	ModeID       uint                       `json:"mode_id" comment:"模式ID"`
	CategoryList []ModeCategoryWebsitesList `json:"category_list" comment:"模式分类网站列表"`
}

type ModeCategoryWebsitesList struct {
	CategoryID uint   `json:"category_id" comment:"模式分类ID"`
	WebsiteIDs []uint `json:"website_ids" comment:"网站id数组"`
}
