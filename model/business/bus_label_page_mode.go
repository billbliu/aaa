package business

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type BusLabelPageMode struct {
	global.GVA_MODEL
	Name string `gorm:"column:name;comment:模式名称" json:"name"`
}

var BusLabelPageModeDao = new(BusLabelPageMode)

func (p *BusLabelPageMode) TableName() string {
	return "bus_label_page_mode"
}

func (p *BusLabelPageMode) TableComment(db *gorm.DB) {
	db.Exec(fmt.Sprintf("ALTER TABLE %s COMMENT '标签页模式表';", p.TableName()))
}
