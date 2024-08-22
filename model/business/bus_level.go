package business

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BusLevel struct {
	global.GVA_MODEL
	Level        string          `gorm:"column:level;default:'游客';comment:客户等级" json:"level"`
	UsdtPrice    decimal.Decimal `gorm:"column:usdt_price;comment:usdt价格" json:"usdtPrice"`
	RmbPrice     decimal.Decimal `gorm:"column:rmb_price;comment:rmb价格" json:"rmbPrice"`
	FlowDown     int64           `gorm:"column:flow_down;comment:下行速度" json:"flowDown"`
	FlowUp       int64           `gorm:"column:flow_up;comment:上行速度" json:"flowUp"`
	FlowInfinite bool            `gorm:"column:flow_infinite;default:true;comment:是否无限流量" json:"flowInfinite"`
}

var BusLevelDao = new(BusLevel)

func (p *BusLevel) TableName() string {
	return "bus_level"
}

func (p *BusLevel) TableComment(db *gorm.DB) {
	db.Exec(fmt.Sprintf("ALTER TABLE %s COMMENT '会员等级表';", p.TableName()))
}
