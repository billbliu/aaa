package business

import (
	"errors"
	"log"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type DepositStatusType uint8

const (
	PAY_STATUS_NOPAY   DepositStatusType = 1 // 未支付
	PAY_STATUS_SUCCESS DepositStatusType = 2 // 支付成功
	PAY_STATUS_CLOSED  DepositStatusType = 3 // 交易关闭
)

type PayType uint8

const (
	PAY_ALIPAY PayType = 1 // 支付宝
	PAY_WECHAT PayType = 2 // 微信
)

// 用户充值订单
type BusCustomerDepositOrder struct {
	global.GVA_MODEL
	Uuid          string            `gorm:"column:uuid;unique;comment:前后端唯一标识" db:"uuid" json:"uuid" form:"uuid"`
	OrderNo       string            `gorm:"column:order_no;unique;comment:商户订单编号" db:"order_no" json:"order_no" form:"order_no"`
	PaymentType   PayType           `gorm:"column:payment_type;comment:支付类型" db:"payment_type" json:"payment_type" form:"payment_type"`
	BusCustomerId uint              `gorm:"column:bus_customer_id;comment:客户ID" db:"bus_customer_id"  json:"bus_customer_id" form:"bus_customer_id"`
	Amount        decimal.Decimal   `gorm:"column:amount;type:decimal(12,2);comment:订单金额" db:"amount" json:"amount"  form:"amount"`
	CodeUrl       string            `gorm:"column:code_url;comment:订单二维码链接" db:"code_url" json:"code_url" form:"code_url"`
	Status        DepositStatusType `gorm:"column:status;comment:订单状态: 0未支付、1已支付、2交易关闭" db:"status" json:"status"  form:"status"`
}

var BusCustomerDepositOrderDao = new(BusCustomerDepositOrder)

func (t *BusCustomerDepositOrder) TableName() string {
	return "bus_customer_deposit_order"
}

// GetCustomerDepositOrderByUuid 通过uuid查询存在的充值订单
func (t *BusCustomerDepositOrder) GetCustomerDepositOrderByUuid(db *gorm.DB, uuid string) (*BusCustomerDepositOrder, error) {
	order := BusCustomerDepositOrder{}
	if err := db.Table(t.TableName()).Where(&BusCustomerDepositOrder{Uuid: uuid}).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetCustomerDepositOrderByOrderNo 通过orderNo查询存在的充值订单
func (t *BusCustomerDepositOrder) GetCustomerDepositOrderByOrderNo(db *gorm.DB, orderNo string) (*BusCustomerDepositOrder, error) {
	order := BusCustomerDepositOrder{}
	if err := db.Table(t.TableName()).Where(&BusCustomerDepositOrder{OrderNo: orderNo}).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

// GetCustomerDepositOrderStatusByOrderNo 通过订单号查询状态
func (t *BusCustomerDepositOrder) GetCustomerDepositOrderStatusByOrderNo(db *gorm.DB, orderNo string) (DepositStatusType, error) {
	order := BusCustomerDepositOrder{}
	if err := db.Table(t.TableName()).Select("status").Where(&BusCustomerDepositOrder{OrderNo: orderNo}).First(&order).Error; err != nil {
		return PAY_STATUS_NOPAY, err
	}
	return order.Status, nil
}

// UpdateCustomerDepositOrderStatusByOrderNo 更新订单状态
func (t *BusCustomerDepositOrder) UpdateCustomerDepositOrderStatusByOrderNo(db *gorm.DB, orderNo string, newStatus DepositStatusType) (*BusCustomerDepositOrder, error) {
	order := BusCustomerDepositOrder{}
	if err := db.Table(t.TableName()).Where(&BusCustomerDepositOrder{OrderNo: orderNo}).First(&order).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在则创建
			return &order, errors.New("订单号不存在")
		} else {
			return &order, err
		}
	}

	if err := db.Model(&order).Update("status", newStatus).Error; err != nil {
		return &order, err
	}

	return &order, nil
}

// FindAllNopayOrder 查找所有未支付订单
func (t *BusCustomerDepositOrder) FindAllNopayOrder(db *gorm.DB) (*[]BusCustomerDepositOrder, error) {
	var orders []BusCustomerDepositOrder
	if err := db.Model(&BusCustomerDepositOrder{}).Where(&BusCustomerDepositOrder{Status: PAY_STATUS_NOPAY}).Find(&orders).Error; err != nil {
		return nil, err
	}
	return &orders, nil
}

// CloseAllNopayTimeoutOrder 关闭所有未支付并且超时订单
func (t *BusCustomerDepositOrder) CloseAllNopayTimeoutOrder(db *gorm.DB) error {
	var orders []BusCustomerDepositOrder
	if err := db.Model(&BusCustomerDepositOrder{}).Where(&BusCustomerDepositOrder{Status: PAY_STATUS_NOPAY}).Find(&orders).Error; err != nil {
		return err
	}
	for _, order := range orders {
		if order.CreatedAt.Add(10 * time.Minute).Before(time.Now()) {
			if err := db.Model(&order).Update("status", PAY_STATUS_CLOSED).Error; err != nil {
				log.Println("db err:", err)
			}
		}
	}
	return nil
}
