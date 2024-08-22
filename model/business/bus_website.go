package business

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type BusWebsite struct {
	global.GVA_MODEL
	Title       string `gorm:"column:title;comment:网站标题" json:"title"`
	Url         string `gorm:"column:url;comment:网站URL" json:"url"`
	Description string `gorm:"column:description;type:text;comment:网站描述" json:"description"`
	FaviconUrl  string `gorm:"column:favicon_url;comment:网站图标URL" json:"faviconUrl"`
}

var BusWebsiteDao = new(BusWebsite)

func (p *BusWebsite) TableName() string {
	return "bus_website"
}

func (p *BusWebsite) TableComment(db *gorm.DB) {
	db.Exec(fmt.Sprintf("ALTER TABLE %s COMMENT '网站信息表';", p.TableName()))
}

func (t *BusWebsite) GetBusWebsiteById(db *gorm.DB, id uint) (*BusWebsite, error) {
	busWebsite := BusWebsite{}
	if err := db.Table(t.TableName()).Where("id = ?", id).First(&busWebsite).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &busWebsite, nil
}
