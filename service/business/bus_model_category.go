package business

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type WebsiteCategoryService struct{}

var WebsiteCategoryServiceApp = new(WebsiteCategoryService)

func (apiService *WebsiteCategoryService) CreateWebsiteCategory(req businessReq.CreateWebsiteCategoryReq) (err error) {
	website := business.BusWebsiteCategory{
		Name: req.Name,
	}
	return global.GVA_DB.Create(&website).Error
}

func (apiService *WebsiteCategoryService) DeleteWebsiteCategory(id int) (err error) {
	var entity business.BusWebsiteCategory
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

func (apiService *WebsiteCategoryService) GetWebsiteCategoryById(id int) (website business.BusWebsiteCategory, err error) {
	err = global.GVA_DB.First(&website, "id = ?", id).Error
	return
}

func (apiService *WebsiteCategoryService) UpdateWebsiteCategory(req businessReq.UpdateWebsiteCategoryReq) (err error) {
	website := business.BusWebsiteCategory{
		Name: req.Name,
	}
	return global.GVA_DB.Where("id = ?", req.ID).Updates(website).Error
}

func (apiService *WebsiteCategoryService) DeleteWebsiteCategorysByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []business.BusWebsiteCategory
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]business.BusWebsiteCategory{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}

		return err
	})
}

// @author: [bill]
// @function: GetWebsiteCategoryList
// @description: 分页获取模式分类
// @param: website business.BusWebsiteCategory, info request.PageInfo, order string, desc bool
// @return: list interface{}, total int64, err error
func (apiService *WebsiteCategoryService) GetWebsiteCategoryList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.BusWebsiteCategory{})
	var websiteList []business.BusWebsiteCategory

	err = db.Count(&total).Error

	if err != nil {
		return websiteList, total, err
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&websiteList).Error
	return websiteList, total, err
}

func (apiService *WebsiteCategoryService) GetAllWebsiteCategorys() (websites []business.BusWebsiteCategory, err error) {
	err = global.GVA_DB.Order("id desc").Find(&websites).Error
	return
}
