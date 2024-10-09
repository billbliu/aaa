package business

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type AdService struct{}

var AdServiceApp = new(AdService)

func (apiService *AdService) CreateAd(req businessReq.CreateAdReq) (err error) {
	ad := business.BusAd{
		Title:        req.Title,
		Description:  req.Description,
		AdType:       req.AdType,
		ContentUrl:   req.ContentUrl,
		TargetUrl:    req.TargetUrl,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Enable:       req.Enable,
		BusAdLevelId: req.BusAdLevelId,
	}
	return global.GVA_DB.Create(&ad).Error
}

func (apiService *AdService) DeleteAd(id int) (err error) {
	var entity business.BusAd
	err = global.GVA_DB.First(&entity, "id = ?", id).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {            // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *AdService) GetAdById(id int) (ad business.BusAd, err error) {
	err = global.GVA_DB.First(&ad, "id = ?", id).Error
	return
}

func (apiService *AdService) UpdateAd(req businessReq.UpdateAdReq) (err error) {
	ad := business.BusAd{
		Title:        req.Title,
		Description:  req.Description,
		AdType:       req.AdType,
		ContentUrl:   req.ContentUrl,
		TargetUrl:    req.TargetUrl,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Enable:       req.Enable,
		BusAdLevelId: req.BusAdLevelId,
	}
	return global.GVA_DB.Where("id = ?", req.ID).Updates(ad).Error
}

func (apiService *AdService) DeleteAdsByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []business.BusAd
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]business.BusAd{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}

		return err
	})
}

// @author: [bill]
// @function: GetAdList
// @description: 分页获取广告数据
// @param: website business.BusAd, info request.PageInfo, order string, desc bool
// @return: list interface{}, total int64, err error
func (apiService *AdService) GetAdList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.BusAd{})
	var websiteList []business.BusAd

	err = db.Count(&total).Error

	if err != nil {
		return websiteList, total, err
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&websiteList).Error
	return websiteList, total, err
}

func (apiService *AdService) GetAllAds() (ads []business.BusAd, err error) {
	err = global.GVA_DB.Order("id desc").Find(&ads).Error
	return
}

func (apiService *AdService) CreateAdLevel(req businessReq.CreateAdLevelReq) (err error) {
	adLevel := business.BusAdLevel{
		Level:           req.Level,
		TotalFlowReword: req.TotalFlowReword,
		OneFlowReword:   req.OneFlowReword,
	}
	return global.GVA_DB.Create(&adLevel).Error
}

func (apiService *AdService) DeleteAdLevel(id int) (err error) {
	var entity business.BusAdLevel
	err = global.GVA_DB.First(&entity, "id = ?", id).Error // 根据id查询api记录
	if errors.Is(err, gorm.ErrRecordNotFound) {            // api记录不存在
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	return nil
}

func (apiService *AdService) GetAdLevelById(id int) (ad business.BusAdLevel, err error) {
	err = global.GVA_DB.First(&ad, "id = ?", id).Error
	return
}

func (apiService *AdService) UpdateAdLevel(req businessReq.UpdateAdLevelReq) (err error) {
	ad := business.BusAdLevel{
		Level:           req.Level,
		TotalFlowReword: req.TotalFlowReword,
		OneFlowReword:   req.OneFlowReword,
	}
	return global.GVA_DB.Where("id = ?", req.ID).Updates(ad).Error
}

func (apiService *AdService) DeleteAdLevelsByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []business.BusAdLevel
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]business.BusAdLevel{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}

		return err
	})
}

func (apiService *AdService) GetAllAdLevels() (ads []business.BusAdLevel, err error) {
	err = global.GVA_DB.Order("id desc").Find(&ads).Error
	return
}
