package business

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type BusCustomerBuyLevelRecord struct {
	global.GVA_MODEL
	BusCustomerId uint            `gorm:"column:bus_customer_id;comment:客户id" json:"busCustomerId"`
	BusLevelId    uint            `gorm:"column:bus_level_id;comment:会员等级id" json:"busLevelId"`
	BuyAt         carbon.DateTime `gorm:"column:buy_at;comment:购买时间" json:"buyAt"`
}

var BusCustomerBuyLevelRecordDao = new(BusCustomerBuyLevelRecord)

func (p *BusCustomerBuyLevelRecord) TableName() string {
	return "bus_customer_buy_level_record"
}

func (p *BusCustomerBuyLevelRecord) TableComment(db *gorm.DB) {
	db.Exec(fmt.Sprintf("ALTER TABLE %s COMMENT '客户购买会员等级记录表';", p.TableName()))
}
