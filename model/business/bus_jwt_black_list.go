package business

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type BusJwtBlackList struct {
	global.GVA_MODEL

	Jwt string `gorm:"column:jwt;size:512;unique;comment:jwt token" json:"jwt"`
}

var BusJwtBlackListDao = new(BusJwtBlackList)

func (jwt *BusJwtBlackList) TableName() string {
	return "bus_jwt_black_list"
}
