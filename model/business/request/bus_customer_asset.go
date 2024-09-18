package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type CustomerAssetBillGetReq struct {
	request.PageInfo
	Behavior uint8 `form:"behavior"  search:"type:behavior;column:behavior;table:user_asset_bill" vd:"$>=0&&$<=3" comment:"行为"`
}

type DepositByAlipayReq struct {
	Uuid  string  `json:"uuid"`
	Money float64 `json:"money" vd:"$>0&&$<=100000000"` // 元
}

type DepositByWxpayReq struct {
	Uuid  string  `json:"uuid"`
	Money float64 `json:"money" vd:"$>0&&$<=100000000"` // 元
}

type DepositByAlipayNotifyReq struct {
	OrderTradeNo string `json:"order_trade_no"` // 商家订单号
	TradeNo      string `json:"trade_no"`       // 支付宝交易号
	TradeState   string `json:"trade_state"`    // 交易状态
	TotalAmount  string `json:"total_amount"`   // 订单金额
}

type DepositByWxpayNotifyReq struct {
	OrderTradeNo string `json:"order_trade_no"` // 商家订单号
	TradeNo      string `json:"trade_no"`       // 支付宝交易号
	TradeState   string `json:"trade_state"`    // 交易状态
	TotalAmount  string `json:"total_amount"`   // 订单金额
}

type DepositOrderStatusReq struct {
	OutTradeNo string `uri:"out_trade_no" form:"out_trade_no" json:"out_trade_no"` // 商家订单号
}

type WithdrawByAlipayReq struct {
	Money   float64 `json:"money" vd:"$>0&&$<=100000000"`         // 元
	Account string  `json:"account" vd:"email($)||phone($,'CN')"` // 支付宝登录号
	Name    string  `json:"name"`                                 // 支付宝账号姓名
}

type WithdrawByWxpayReq struct {
	Money   float64 `json:"money" vd:"$>=1&&$<=100000000"` // 元
	Account string  `json:"account"`                       //微信账号
}

type WithdrawOrderStatusReq struct {
	OutTradeNo string `uri:"out_trade_no" form:"out_trade_no" json:"out_trade_no"` // 商家订单号
}
