package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type GpuCenterProductMap struct {
	global.GVA_MODEL
	BusLabelPageModeId uint `gorm:"column:bus_label_page_mode_id;comment:标签页模式表id" json:"busLabelPageModeId"`
	BusWebsiteId       uint `gorm:"column:bus_website_id;comment:网站信息表id" json:"busWebsiteId"`
}

func (t *GpuCenterProductMap) TableName() string {
	return "bus_label_page_model_website_map"
}
