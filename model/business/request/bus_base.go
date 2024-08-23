package request

type GetEmailSMSReq struct {
	Email   string `uri:"email" form:"email" json:"email" binding:"required" label:"邮箱"`
	SmsType string `uri:"sms_type" form:"sms_type" json:"sms_type" binding:"required" label:"短信类型"`
}

type LoginOrRegisterByEmailSMSReq struct {
	Email   string `uri:"email" form:"email" json:"email" binding:"required" label:"邮箱"`
	SmsCode string `uri:"sms_code" form:"sms_code" json:"sms_code" binding:"required" label:"邮箱验证码"`
}

type GetPhoneSMSReq struct {
	Phone   string `uri:"phone" form:"phone" json:"phone" binding:"required" label:"手机号"`
	SmsType string `uri:"sms_type" form:"sms_type" json:"sms_type" binding:"required" label:"短信类型"`
}

type LoginOrRegisterByPhoneSMSReq struct {
	Phone   string `uri:"phone" form:"phone" json:"phone" binding:"required" label:"手机号"`
	SmsCode string `uri:"sms_code" form:"sms_code" json:"sms_code" binding:"required" label:"手机验证码"`
}
