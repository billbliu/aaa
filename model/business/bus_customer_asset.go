package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

type BusCustomerAsset struct {
	global.GVA_MODEL
	TotalAmount   decimal.Decimal `gorm:"column:total_amount;type:decimal(12,2);comment:冻结金额" json:"totalAmount"`
	FreezeAmount  decimal.Decimal `gorm:"column:freeze_amount;type:decimal(12,2);comment:冻结金额" json:"freezeAmount"`
	BusCustomerId uint            `gorm:"column:bus_customer_id;comment:用户ID" json:"busCustomerId"`
	BusCustomer   BusCustomer     `gorm:"foreignKey:bus_customer_id;comment:用户" json:"user"`
}

var BusCustomerAssetDao = new(BusCustomerAsset)

func (t *BusCustomerAsset) TableName() string {
	return "bus_customer_asset"
}
