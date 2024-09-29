// 用户表
package business

import (
	"github.com/golang-module/carbon/v2"
	"github.com/shopspring/decimal"
)

type TOperationReport struct {
	ID        uint            `gorm:"column:id;primaryKey;autoIncrement" db:"id" json:"id" form:"id"`
	CreatedAt carbon.DateTime `gorm:"column:created_at" db:"created_at" json:"-" form:"created_at"`
	UpdatedAt carbon.DateTime `gorm:"column:updated_at" db:"updated_at" json:"-" form:"updated_at"`

	WhichDay           carbon.Date `gorm:"column:which_day;type:date;comment:统计时间" db:"which_day" json:"which_day" form:"which_day"`
	TotalMembers       int64       `gorm:"column:total_members;comment:总会员数" db:"total_members" json:"total_members"  form:"total_members"`
	DayRegisterMembers int64       `gorm:"column:day_register_members;comment:当天新注册会员数" db:"day_register_members" json:"day_register_members"  form:"day_register_members"`
	DayActiveMembers   int64       `gorm:"column:day_active_members;comment:当天活跃会员数" db:"day_active_members" json:"day_active_members"  form:"day_active_members"`

	TotalDepositUsdtAmount decimal.Decimal `gorm:"column:total_deposit_usdt_amount;type:decimal(12,2);comment:总usdt充值金额" db:"total_deposit_usdt_amount" json:"total_deposit_usdt_amount"  form:"total_deposit_usdt_amount"`
	DayDepositUsdtAmount   decimal.Decimal `gorm:"column:day_deposit_usdt_amount;type:decimal(12,2);comment:当天usdt充值金额" db:"day_deposit_usdt_amount" json:"day_deposit_usdt_amount"  form:"day_deposit_usdt_amount"`

	TotalDepositCoin             decimal.Decimal `gorm:"column:total_deposit_coin;type:decimal(12,2);comment:总金币充值金额" db:"total_deposit_coin" json:"total_deposit_coin"  form:"total_deposit_coin"`
	DayBinaryDivinationCoin      decimal.Decimal `gorm:"column:day_binary_divination_coin;type:decimal(12,2);comment:当天二选一占卜金币" db:"day_binary_divination_coin" json:"day_binary_divination_coin"  form:"day_binary_divination_coin"`
	DayRandomDivinationCoin      decimal.Decimal `gorm:"column:day_random_divination_coin;type:decimal(12,2);comment:当天无牌阵占卜金币" db:"day_random_divination_coin" json:"day_random_divination_coin"  form:"day_random_divination_coin"`
	DayTrinityDivinationCoin     decimal.Decimal `gorm:"column:day_trinity_divination_coin;type:decimal(12,2);comment:当天圣三角占卜金币" db:"day_trinity_divination_coin" json:"day_trinity_divination_coin"  form:"day_trinity_divination_coin"`
	DayCelticCrossDivinationCoin decimal.Decimal `gorm:"column:day_celtic_cross_divination_coin;type:decimal(12,2);comment:当天凯尔特十字占卜金币" db:"day_celtic_cross_divination_coin" json:"day_celtic_cross_divination_coin"  form:"day_celtic_cross_divination_coin"`
}

var TOperationReportDao = new(TOperationReport)

func (t *TOperationReport) TableName() string {
	return "t_operation_report"
}
