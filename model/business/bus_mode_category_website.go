package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type BusModeCategoryWebsite struct {
	global.GVA_MODEL
	BusModeID            uint `gorm:"column:bus_mode_id;uniqueIndex:idx_bus_mode_category_website" json:"busModeID"`
	BusWebsiteCategoryID uint `gorm:"column:bus_website_category_id;uniqueIndex:idx_bus_mode_category_website" json:"busWebsiteCategoryID"`
	BusWebsiteID         uint `gorm:"column:bus_website_id;uniqueIndex:idx_bus_mode_category_website" json:"busWebsiteID"`
}

func (t *BusModeCategoryWebsite) TableName() string {
	return "bus_mode_category_website"
}
