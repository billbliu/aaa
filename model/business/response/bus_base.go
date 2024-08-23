package response

import "github.com/golang-module/carbon/v2"

type GetSMSRes struct {
	IsAlreadySend bool   `json:"is_already_send"`
	SmsCode       string `json:"sms_code"`
	SendRestSec   int64  `json:"send_rest_sec"`
	ExpireSec     int64  `json:"expire_sec"`
}

type CheckSMSRes struct {
	SmsCode   string `json:"sms_code"`
	IsExpired bool   `json:"is_expired"`
	ExpireSec int    `json:"expire_sec"`
}

type LoginResponse struct {
	UserId   uint        `json:"user_id"` //用户id
	Token    string      `json:"token"`   //登录成功
	NickName string      `json:"nick_name"`
	Email    string      `json:"email"`
	Phone    string      `json:"phone"`
	Sex      uint8       `json:"sex"`
	Birth    carbon.Date `json:"birth"`
}
