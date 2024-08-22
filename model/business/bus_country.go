package business

type BusCountry struct {
	ID          uint    `gorm:"column:id;primaryKey;" db:"id" json:"id" form:"id"`
	Code        uint    `gorm:"column:code;comment:国家编号" db:"code" json:"code" form:"code"`
	CountryCode string  `gorm:"column:country_code;comment:国际域名缩写" db:"country_code" json:"country_code" form:"country_code"`
	ChineseName string  `gorm:"column:chinese_name;comment:中文名" db:"chinese_name" json:"chinese_name" form:"chinese_name"`
	EnglishName string  `gorm:"column:english_name;comment:英文名" db:"english_name" json:"english_name" form:"english_name"`
	Latitude    float32 `gorm:"column:latitude;comment:纬度" db:"latitude" json:"latitude" form:"latitude"`
	Longitude   float32 `gorm:"column:longitude;comment:经度" db:"longitude" json:"longitude" form:"longitude"`
}

var BusCountryDao = new(BusCountry)

func (t *BusCountry) TableName() string {
	return "t_country"
}
