package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
)

type BehaviorType uint8

const (
	BEHAVIOR_DEPOSIT BehaviorType = 1 // 充值
	BEHAVIOR_USE     BehaviorType = 2 // 使用
)

type StatusType uint8

const (
	STATUS_NEW     StatusType = 1 // 新建
	STATUS_SUCCESS StatusType = 2 // 成功
	STATUS_FAIL    StatusType = 3 // 失败
)

type BusCustomerAssetOrder struct {
	global.GVA_MODEL

	OrderId    string          `gorm:"column:order_id;unique;comment:订单ID" json:"orderId"`
	Amount     decimal.Decimal `gorm:"column:amount;type:decimal(12,2);comment:金额" json:"amount"`
	UsdtAmount decimal.Decimal `gorm:"column:usdt_amount;type:decimal(12,2);comment:usdt金额" json:"usdtAmount"`
	Status     StatusType      `gorm:"column:status;comment:状态: 1未完成、2已完成" json:"status"`
	Info       string          `gorm:"column:info;comment:订单描述" json:"info"`
	Behavior   BehaviorType    `gorm:"column:behavior;comment:行为:1充值、2使用" json:"behavior"`
	UserId     uint            `gorm:"column:user_id;comment:订单用户ID" json:"userId"`
}

var BusCustomerAssetOrderDao = new(BusCustomerAssetOrder)

func (t *BusCustomerAssetOrder) TableName() string {
	return "bus_customer_asset_order"
}
