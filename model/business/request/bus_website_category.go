package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CreateWebsiteCategoryReq struct {
	Name string `json:"name" comment:"模式分类名称"`
}

type UpdateWebsiteCategoryReq struct {
	ID   uint   `json:"id" comment:"网站ID"`
	Name string `json:"name" comment:"模式分类名称"`
}

// WebsiteCategory分页条件查询及排序结构体
type SearchWebsiteCategoryParams struct {
	// business.BusWebsiteCategory
	request.PageInfo
}
