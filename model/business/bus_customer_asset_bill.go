package business

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BehaviorType uint8

const (
	BEHAVIOR_DEPOSIT BehaviorType = 1 // 充值
	// BEHAVIOR_WITHDRAW     BehaviorType = 2 // 提现
	BEHAVIOR_MEMBER_PURCHASE BehaviorType = 2 // 会员购买
)

type PayDirectionType uint8

const (
	PAY_DIRECTION_SUB PayDirectionType = 1
	PAY_DIRECTION_ADD PayDirectionType = 2
)

type OrderType uint8

const (
	ORDER_TYPE_DEPOSIT    OrderType = 1 // 充值
	// ORDER_TYPE_WITHDRAW   OrderType = 2 // 提现
	ORDER_TYPE_MEMBER_PURCHASE   OrderType = 2 // 会员购买
)

type BusCustomerAssetBill struct {
	global.GVA_MODEL
	Before decimal.Decimal `gorm:"column:before;type:decimal(14,2);comment:交易前金额" db:"before" json:"before"  form:"before"`
	Amount decimal.Decimal `gorm:"column:amount;type:decimal(14,2);comment:交易金额" db:"amount" json:"amount"  form:"amount"`
	After  decimal.Decimal `gorm:"column:after;type:decimal(14,2);comment:交易后金额" db:"after" json:"after"  form:"after"`

	Behavior      BehaviorType     `gorm:"column:behavior;comment:行为:1充值、2提现、3商城订单" db:"behavior" json:"behavior"  form:"behavior"`
	PayDirection  PayDirectionType `gorm:"column:pay_direction;comment:行为:1减少、2增加" db:"pay_direction" json:"pay_direction"  form:"pay_direction"`
	BusCustomerId uint             `gorm:"column:bus_customer_id;comment:客户ID" db:"bus_customer_id"  json:"bus_customer_id" form:"bus_customer_id"`
	OrderId       uint             `gorm:"column:order_id;comment:订单表ID" db:"order_id"  json:"order_id" form:"order_id"`
	OrderType     OrderType        `gorm:"column:order_type;comment:订单类型" db:"order_type"  json:"order_type" form:"order_type"`
}

var BusCustomerAssetBillDao = new(BusCustomerAssetBill)

func (t *BusCustomerAssetBill) TableName() string {
	return "bus_customer_asset_bill"
}

func (t *BusCustomerAssetBill) IsClearingByOrderIdAndBehavior(db *gorm.DB, orderId uint, behavior BehaviorType) (bool, error) {
	bills := []BusCustomerAssetBill{}
	if err := db.Table(t.TableName()).Select("id").Where(&BusCustomerAssetBill{OrderId: orderId, Behavior: behavior}).Find(&bills).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}
	if len(bills) == 0 {
		return false, nil
	}
	return true, nil
}
