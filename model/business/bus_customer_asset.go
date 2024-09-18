package business

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BusCustomerAsset struct {
	global.GVA_MODEL
	Total  decimal.Decimal `gorm:"column:total;type:decimal(14,2);comment:总金额" db:"total" json:"total"  form:"total"`
	Freeze decimal.Decimal `gorm:"column:freeze;type:decimal(14,2);comment:冻结金额" db:"freeze" json:"freeze"  form:"freeze"`
	// Avaiable decimal.Decimal `gorm:"column:avaiable;type:decimal(12,2);comment:可用金额" db:"avaiable" json:"avaiable"  form:"avaiable"`
	Enable bool `gorm:"column:enable;comment:是否启用;default:true" db:"enable" json:"enable"  form:"enable"`

	BusCustomerId uint `gorm:"column:bus_customer_id;comment:客户ID" db:"bus_customer_id"  json:"bus_customer_id" form:"bus_customer_id"`
}

var BusCustomerAssetDao = new(BusCustomerAsset)

func (t *BusCustomerAsset) TableName() string {
	return "bus_customer_asset"
}

// GetCustomerAssetByID 通过id查询用户资产信息
func (t *BusCustomerAsset) GetCustomerAssetByID(db *gorm.DB, busCustomerId uint) (*BusCustomerAsset, error) {
	tUserAsset := BusCustomerAsset{}
	if err := db.Table(t.TableName()).Where(&BusCustomerAsset{BusCustomerId: busCustomerId}).First(&tUserAsset).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 不存在则创建
			tUserAsset.BusCustomerId = busCustomerId
			if err = db.Create(&tUserAsset).Error; err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	return &tUserAsset, nil
}
