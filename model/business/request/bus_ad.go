package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/golang-module/carbon/v2"
)

type CreateAdReq struct {
	// 广告标题
	Title string `json:"title"`
	// 广告描述
	Description string `json:"description"`
	// 广告类型
	AdType business.AdType `json:"adType"`
	// 广告内容(图片/视频 URL)
	ContentUrl string `json:"contentUrl"`
	// 广告点击跳转链接
	TargetUrl string `json:"targetUrl"`

	// 广告开始展示时间
	StartTime carbon.DateTime `json:"startTime"`
	// 广告结束展示时间
	EndTime carbon.DateTime `json:"endTime"`
	// 是否启用
	Enable bool `json:"enable"`
	// 广告等级id
	BusAdLevelId uint `json:"busAdLevelId"`
}

type UpdateAdReq struct {
	// 广告ID
	ID uint `json:"id"`
	// 广告标题
	Title string `json:"title"`
	// 广告描述
	Description string `json:"description"`
	// 广告类型
	AdType business.AdType `json:"adType"`
	// 广告内容(图片/视频 URL)
	ContentUrl string `json:"contentUrl"`
	// 广告点击跳转链接
	TargetUrl string `json:"targetUrl"`

	// 广告开始展示时间
	StartTime carbon.DateTime `json:"startTime"`
	// 广告结束展示时间
	EndTime carbon.DateTime `json:"endTime"`
	// 是否启用
	Enable bool `json:"enable"`
	// 广告等级id
	BusAdLevelId uint `json:"busAdLevelId"`
}

// Ad分页条件查询及排序结构体
type SearchAdParams struct {
	// business.BusAd
	request.PageInfo
}

type CreateAdLevelReq struct {
	// 广告等级
	Level uint8 `json:"level"`
	// 总流量奖励(M)
	TotalFlowReword int64 `json:"totalFlowReword"`
	// 单个流量奖励(M)
	OneFlowReword int64 `json:"oneFlowReword"`
}

type UpdateAdLevelReq struct {
	// 广告等级ID
	ID uint `json:"id"`
	// 广告等级
	Level uint8 `json:"level"`
	// 总流量奖励(M)
	TotalFlowReword int64 `json:"totalFlowReword"`
	// 单个流量奖励(M)
	OneFlowReword int64 `json:"oneFlowReword"`
}
