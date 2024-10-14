package business

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 支付渠道表
type BusPaymentChannel struct {
	global.GVA_MODEL
	ChannelName string `gorm:"column:channel_name;comment:支付渠道名称（微信、支付宝）" json:"channelName"`
	Description string `gorm:"column:description;comment:支付渠道描述" json:"description"`
	Status      int    `gorm:"column:status;comment:支付渠道状态(1: 启用, 0: 停用)" json:"status"`
}

var BusPaymentChannelDao = new(BusPaymentChannel)

func (t *BusPaymentChannel) TableName() string {
	return "bus_payment_channel"
}
