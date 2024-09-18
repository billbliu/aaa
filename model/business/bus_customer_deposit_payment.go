package business

import (
	"github.com/shopspring/decimal"
)

// 用户充值支付记录
type BusCustomerDepositPayment struct {
	ID            uint            `gorm:"column:id;primaryKey;autoIncrement" db:"id" json:"id" form:"id"`
	OrderNo       string          `gorm:"column:order_no;unique;comment:商户订单编号" db:"order_no" json:"order_no" form:"order_no"`
	PaymentType   PayType         `gorm:"column:payment_type;comment:支付类型" db:"payment_type" json:"payment_type" form:"payment_type"`
	TransactionId string          `gorm:"column:transaction_id;comment:支付系统交易编号" db:"transaction_id" json:"transaction_id" form:"transaction_id"`
	TradeType     string          `gorm:"column:trade_type;comment:交易类型" db:"trade_type" json:"trade_type" form:"trade_type"`
	TradeState    string          `gorm:"column:trade_state;comment:交易状态" db:"trade_state" json:"trade_state" form:"trade_state"`
	Amount        decimal.Decimal `gorm:"column:amount;type:decimal(12,2);comment:支付金额" db:"amount" json:"amount"  form:"amount"`
	Content       string          `gorm:"column:content;comment:通知参数" db:"content" json:"content" form:"content"`
}

func (t *BusCustomerDepositPayment) TableName() string {
	return "user_deposit_payment"
}
