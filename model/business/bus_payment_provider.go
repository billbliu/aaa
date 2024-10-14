package business

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 支付接口提供商表
type BusPaymentProvider struct {
	global.GVA_MODEL
	ProviderName string `gorm:"column:provider_name;comment:支付提供商名称" json:"providerName"`
	ProviderURL  string `gorm:"column:provider_url;comment:支付提供商接口URL" json:"providerUrl"`
	Description  string `gorm:"column:description;comment:支付提供商描述" json:"description"`
	Status       int    `gorm:"column:status;comment:支付提供商状态(1: 启用, 0: 停用)" json:"status"`
}

var BusPaymentProviderDao = new(BusPaymentProvider)

func (t *BusPaymentProvider) TableName() string {
	return "bus_payment_provider"
}
