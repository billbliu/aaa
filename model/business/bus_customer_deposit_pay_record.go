package business

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DepositPayType uint8

const (
	PAY_PAYPAL DepositPayType = 1 // Paypal
)

// 用户充值支付记录
type BusCustomerDepositPayRecord struct {
	global.GVA_MODEL
	OrderNo       string          `gorm:"column:order_no;unique;comment:商户订单编号" json:"orderNo"`
	PaymentType   DepositPayType  `gorm:"column:payment_type;comment:支付类型" json:"paymentType"`
	TransactionId string          `gorm:"column:transaction_id;comment:支付系统交易编号" json:"transactionId"`
	TradeType     string          `gorm:"column:trade_type;comment:交易类型" json:"tradeType"`
	TradeState    string          `gorm:"column:trade_state;comment:交易状态" json:"tradeState"`
	Amount        decimal.Decimal `gorm:"column:amount;type:decimal(12,2);comment:支付金额" json:"amount"`
	Content       string          `gorm:"column:content;type:text;comment:通知参数" json:"content"`
}

var BusCustomerDepositPayRecordDao = new(BusCustomerDepositPayRecord)

func (t *BusCustomerDepositPayRecord) TableName() string {
	return "bus_customer_deposit_pay_record"
}

// IsExistByTransactionId 通过TransactionId判断记录是否存在
func (t *BusCustomerDepositPayRecord) IsExistByTransactionId(db *gorm.DB, transactionId string) (bool, error) {
	tUserDepositPayment := BusCustomerDepositPayRecord{}
	if err := db.Table(t.TableName()).Where("transaction_id = ?", transactionId).First(&tUserDepositPayment).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
