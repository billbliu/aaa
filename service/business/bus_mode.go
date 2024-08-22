package business

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	businessRes "github.com/flipped-aurora/gin-vue-admin/server/model/business/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type ModeService struct{}

var ModeServiceApp = new(ModeService)

func (apiService *ModeService) CreateMode(req businessReq.CreateModeReq) (err error) {
	website := business.BusMode{
		Name: req.Name,
	}
	return global.GVA_DB.Create(&website).Error
}

func (apiService *ModeService) DeleteMode(id int) (err error) {
	var entity business.BusMode
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

func (apiService *ModeService) GetModeById(id int) (website business.BusMode, err error) {
	err = global.GVA_DB.First(&website, "id = ?", id).Error
	return
}

func (apiService *ModeService) UpdateMode(req businessReq.UpdateModeReq) (err error) {
	website := business.BusMode{
		Name: req.Name,
	}
	return global.GVA_DB.Where("id = ?", req.ID).Updates(website).Error
}

func (apiService *ModeService) DeleteModesByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []business.BusMode
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]business.BusMode{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}

		return err
	})
}

// @author: [bill]
// @function: GetModeList
// @description: 分页获取网站数据
// @param: website business.BusMode, info request.PageInfo, order string, desc bool
// @return: list interface{}, total int64, err error
func (apiService *ModeService) GetModeList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.BusMode{})
	var websiteList []business.BusMode

	err = db.Count(&total).Error

	if err != nil {
		return websiteList, total, err
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&websiteList).Error
	return websiteList, total, err
}

func (apiService *ModeService) GetAllModes() (websites []business.BusMode, err error) {
	err = global.GVA_DB.Order("id desc").Find(&websites).Error
	return
}

func (apiService *ModeService) ModeBindingCategoryWebsites(req businessReq.ModeBindingCategoryWebsitesReq) (err error) {
	global.GVA_DB.Model(&business.BusModeCategoryWebsite{}).Where("bus_mode_id = ?", req.ModeID).Delete(&[]business.BusModeCategoryWebsite{}, "bus_mode_id = ?", req.ModeID)
	busModeCategoryWebsites := make([]business.BusModeCategoryWebsite, 0)
	for _, category := range req.CategoryList {
		for _, websiteID := range category.WebsiteIDs {
			busModeCategoryWebsites = append(busModeCategoryWebsites, business.BusModeCategoryWebsite{
				BusModeID:            req.ModeID,
				BusWebsiteCategoryID: category.CategoryID,
				BusWebsiteID:         websiteID,
			})
		}
	}
	return global.GVA_DB.Create(&busModeCategoryWebsites).Error
}

func (apiService *ModeService) GetModeBindingCategoryWebsites(id int) (res businessRes.GetModeBindingCategoryWebsitesRes, err error) {
	busMode, err := business.BusModeDao.GetBusModeById(global.GVA_DB, id)
	if err != nil {
		return res, err
	}

	res.Mode = *busMode

	busModeCategoryWebsiteList := []business.BusModeCategoryWebsite{}
	global.GVA_DB.Model(&business.BusModeCategoryWebsite{}).Where("bus_mode_id = ?", id).Find(&busModeCategoryWebsiteList)
	// maps := make(map[uint][]uint, 0)
	// for _, busModeCategoryWebsite := range busModeCategoryWebsiteList {

	// }
	return
}
