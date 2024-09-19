package response

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/shopspring/decimal"
)

type CustomerAssetGetRes struct {
	Total  decimal.Decimal `json:"total"`
	Freeze decimal.Decimal `json:"freeze"`
	Avail  decimal.Decimal `json:"Avail"`
}

type CustomerAssetBillGetRes struct {
	Before   string `json:"before"`
	Amount   string `json:"amount"`
	After    string `json:"after"`
	Behavior string `json:"behavior"`
	// PayDirection string          `json:"pay_direction"`
	// UserId       uint            `json:"user_id"`
	// OrderId      uint            `json:"order_id"`
	CreatedAt time.Time `json:"createdAt"`
}

type DepositByAlipayRes struct {
	OutTradeNo string `json:"out_trade_no"` // 商家订单号
	PayUrl     string `json:"pay_url"`      // 支付链接
}

type DepositByWxpayRes struct {
	OutTradeNo string `json:"out_trade_no"` // 商家订单号
	QrCode     string `json:"qr_code"`      // 支付链接
}

type DepositOrderStatusRes struct {
	OutTradeNo string                     `json:"out_trade_no"` // 商家订单号
	Status     business.DepositStatusType `json:"status"`       //
}

type WithdrawByAlipayRes struct {
	OutTradeNo string `json:"out_trade_no"` // 商家订单号
	IsSuccess  bool   `json:"is_success"`   // 是否已转账成功
}

type WithdrawByWxpayRes struct {
	OutTradeNo string `json:"out_trade_no"` // 商家订单号
}
