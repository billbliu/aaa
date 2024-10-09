package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/golang-module/carbon/v2"
)

type AdType uint8

const (
	AD_SPLASH AdType = 1 // 开屏
	AD_BANNER AdType = 2 // banner
	AD_POPUP  AdType = 3 // 弹窗
	// AD_REWARD_VIDEO AdType = 4 // 激励视频
)

type BusAd struct {
	global.GVA_MODEL
	Title       string `gorm:"column:title;comment:广告标题" json:"title"`
	Description string `gorm:"column:description;type:text;comment:广告描述" json:"description"`
	AdType      AdType `gorm:"column:ad_type;comment:广告类型" json:"adType"`

	ContentUrl string `gorm:"column:content_url;comment:广告内容(图片/视频 URL)" json:"contentUrl"`
	TargetUrl  string `gorm:"column:target_url;comment:广告点击跳转链接" json:"targetUrl"`

	StartTime carbon.DateTime `gorm:"column:start_time;comment:广告开始展示时间" json:"startTime"`
	EndTime   carbon.DateTime `gorm:"column:end_time;comment:广告结束展示时间" json:"endTime"`
	Enable    bool            `gorm:"column:enable;comment:是否启用" json:"enable"`

	BusAdLevelId uint `gorm:"column:bus_ad_level_id;comment:广告等级ID" json:"busAdLevelId"`
}

var BusAdDao = new(BusAd)

func (t *BusAd) TableName() string {
	return "bus_ad"
}
