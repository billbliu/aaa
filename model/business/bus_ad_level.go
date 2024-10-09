package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type BusAdLevel struct {
	global.GVA_MODEL
	Level           uint8 `gorm:"column:level;comment:广告等级" json:"level"`
	TotalFlowReword int64 `gorm:"column:total_flow_reword;comment:总流量奖励(M)" json:"totalFlowReword"`
	OneFlowReword   int64 `gorm:"column:one_flow_reword;comment:单个流量奖励(M)" json:"oneFlowReword"`
}

var BusAdLevelDao = new(BusAdLevel)

func (t *BusAdLevel) TableName() string {
	return "bus_ad_level"
}
