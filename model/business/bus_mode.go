package business

import (
	"errors"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
)

type BusMode struct {
	global.GVA_MODEL
	Name string `gorm:"column:name;comment:模式名称" json:"name"`
}

var BusModeDao = new(BusMode)

func (t *BusMode) TableName() string {
	return "bus_mode"
}

func (t *BusMode) TableComment(db *gorm.DB) {
	db.Exec(fmt.Sprintf("ALTER TABLE %s COMMENT '模式表';", t.TableName()))
}

func (t *BusMode) GetBusModeById(db *gorm.DB, id int) (*BusMode, error) {
	busMode := BusMode{}
	if err := db.Table(t.TableName()).Where("id = ?", id).First(&busMode).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &busMode, nil
}
