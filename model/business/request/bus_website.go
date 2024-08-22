package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type CreateWebsiteReq struct {
	Title       string `json:"title" comment:"网站标题"`
	Url         string `json:"url" comment:"网站URL"`
	Description string `json:"description" comment:"网站描述"`
	FaviconUrl  string `json:"faviconUrl" comment:"网站图标URL"`
}

type UpdateWebsiteReq struct {
	ID          uint   `json:"id" comment:"网站ID"`
	Title       string `json:"title" comment:"网站标题"`
	Url         string `json:"url" comment:"网站URL"`
	Description string `json:"description" comment:"网站描述"`
	FaviconUrl  string `json:"faviconUrl" comment:"网站图标URL"`
}

// Website分页条件查询及排序结构体
type SearchWebsiteParams struct {
	// business.BusWebsite
	request.PageInfo
}
