package business

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 支付配置表
type BusPaymentConfig struct {
	global.GVA_MODEL
	ChannelID  uint    `gorm:"column:channel_id;comment:'支付渠道ID'" json:"channelId"`
	ProviderID uint    `gorm:"column:provider_id;comment:'支付提供商ID'" json:"providerId"`
	Amount     float64 `gorm:"column:amount;comment:'支付金额'" json:"amount"`
	Status     int     `gorm:"column:status;comment:'配置状态(1: 启用, 0: 停用)'" json:"status"`
}

var BusPaymentConfigDao = new(BusPaymentConfig)

func (t *BusPaymentConfig) TableName() string {
	return "bus_payment_config"
}
