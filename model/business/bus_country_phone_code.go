package business

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// https://www.okx.com/v3/users/support/common/countriesIncludeCommon?functionType=6&t=1724049583702
type BusCountryPhoneCode struct {
	global.GVA_MODEL
	CountryCode   string `gorm:"column:country_code;comment:国际域名缩写" json:"countryCode"`
	TelephoneCode string `gorm:"column:telephone_code;comment:电话代码" json:"telephoneCode"`
	ChineseName   string `gorm:"column:chinese_name;comment:中文名" json:"chineseName"`
	EnglishName   string `gorm:"column:english_name;comment:英文名" json:"englishName"`
	FormalName    string `gorm:"column:formal_name;comment:正式的名字" json:"formalName"`
	StartChar     string `gorm:"column:start_char;comment:首字母" json:"startChar"`
	NationalFlag  string `gorm:"column:national_flag;comment:国旗" json:"nationalFlag"`
	Enable        bool   `gorm:"column:enable;comment:是否启用" form:"enable"`
}

var BusCountryPhoneCodeDao = new(BusCountryPhoneCode)

func (t *BusCountryPhoneCode) TableName() string {
	return "bus_country_phone_code"
}
