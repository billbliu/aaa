package redis_model

import "github.com/golang-module/carbon/v2"

type RedisSmsCode struct {
	Type     string          `json:"type"` //SmsCodeLogin, modify_passwd
	SmsCode  string          `json:"sms_code"`
	SendTime carbon.DateTime `json:"send_time"`
}
