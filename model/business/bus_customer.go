package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/golang-module/carbon/v2"
)

type BusCustomer struct {
	global.GVA_MODEL
	Phone     string      `gorm:"column:phone;default:'';comment:电话号码" json:"phone"`
	Passwd    string      `gorm:"column:passwd;comment:密码" json:"-"`
	Salt      string      `gorm:"column:salt;comment:必填" json:"-"`
	Email     string      `gorm:"column:email;default:'';comment:邮箱" json:"email"`
	NickName  string      `gorm:"column:nick_name;comment:昵称" json:"nickName"`
	AvatarUrl string      `gorm:"column:avatar_url;default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像url" json:"avatarUrl" ` // 用户头像
	Sex       uint8       `gorm:"column:sex;comment:性别 1 男 2 女" json:"sex"`
	Birth     carbon.Date `gorm:"column:birth;comment:出生日期" json:"birth"`
	CountryId uint        `gorm:"column:country_id;default:0;comment:国家id" json:"countryId"`
}

var BusCustomerDao = new(BusCustomer)

func (p *BusCustomer) TableName() string {
	return "bus_customer"
}
