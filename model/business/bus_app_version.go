package business

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// BusAppUpdate 表示应用更新版本的信息
type BusAppUpdate struct {
	global.GVA_MODEL
	AppName         string    `gorm:"column:app_name;type:varchar(255);comment:应用名称" json:"appName"`
	Version         string    `gorm:"column:version;type:varchar(50);comment:应用版本号" json:"version"`
	ReleaseDate     time.Time `gorm:"column:release_date;not null;comment:版本发布的日期" json:"releaseDate"`
	UpdateURL       string    `gorm:"column:update_url;type:varchar(255);not null;comment:版本更新包的下载链接" json:"updateUrl"`
	Changelog       string    `gorm:"column:changelog;type:text;comment:更新内容说明" json:"changelog"`
	MandatoryUpdate bool      `gorm:"column:mandatory_update;default:false;comment:是否强制更新" json:"mandatoryUpdate"`
	Platform        string    `gorm:"column:platform;type:enum('iOS', 'Android');not null;comment:适用的平台" json:"platform"`
	IsPublished     bool      `gorm:"column:is_published;default:false;comment:版本是否已发布" json:"isPublished"`
}

var BusAppUpdateDao = new(BusAppUpdate)

func (t *BusAppUpdate) TableName() string {
	return "bus_app_update"
}
