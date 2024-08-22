package business

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type BusWebsiteCategory struct {
	global.GVA_MODEL
	Name string `gorm:"column:name;comment:分类名称" json:"name"`
}

var BusWebsiteCategoryDao = new(BusWebsiteCategory)

func (p *BusWebsiteCategory) TableName() string {
	return "bus_website_category"
}

func (p *BusWebsiteCategory) TableComment(db *gorm.DB) {
	db.Exec(fmt.Sprintf("ALTER TABLE %s COMMENT '网站分类表';", p.TableName()))
}

func (t *BusWebsiteCategory) GetBusWebsiteCategoryById(db *gorm.DB, id uint) (*BusWebsiteCategory, error) {
	busWebsiteCategory := BusWebsiteCategory{}
	if err := db.Table(t.TableName()).Where("id = ?", id).First(&busWebsiteCategory).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &busWebsiteCategory, nil
}
