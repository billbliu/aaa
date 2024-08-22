package business

import (
	"errors"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"gorm.io/gorm"
)

type WebsiteService struct{}

var WebsiteServiceApp = new(WebsiteService)

func (apiService *WebsiteService) CreateWebsite(req businessReq.CreateWebsiteReq) (err error) {
	website := business.BusWebsite{
		Title:       req.Title,
		Url:         req.Url,
		Description: req.Description,
		FaviconUrl:  req.FaviconUrl,
	}
	return global.GVA_DB.Create(&website).Error
}

func (apiService *WebsiteService) DeleteWebsite(id int) (err error) {
	var entity business.BusWebsite
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

func (apiService *WebsiteService) GetWebsiteById(id int) (website business.BusWebsite, err error) {
	err = global.GVA_DB.First(&website, "id = ?", id).Error
	return
}

func (apiService *WebsiteService) UpdateWebsite(req businessReq.UpdateWebsiteReq) (err error) {
	website := business.BusWebsite{
		Title:       req.Title,
		Url:         req.Url,
		Description: req.Description,
		FaviconUrl:  req.FaviconUrl,
	}
	return global.GVA_DB.Where("id = ?", req.ID).Updates(website).Error
}

func (apiService *WebsiteService) DeleteWebsitesByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []business.BusWebsite
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]business.BusWebsite{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}

		return err
	})
}

// @author: [bill]
// @function: GetWebsiteList
// @description: 分页获取网站数据
// @param: website business.BusWebsite, info request.PageInfo, order string, desc bool
// @return: list interface{}, total int64, err error
func (apiService *WebsiteService) GetWebsiteList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&business.BusWebsite{})
	var websiteList []business.BusWebsite

	err = db.Count(&total).Error

	if err != nil {
		return websiteList, total, err
	}

	err = db.Limit(limit).Offset(offset).Order("id desc").Find(&websiteList).Error
	return websiteList, total, err
}

func (apiService *WebsiteService) GetAllWebsites() (websites []business.BusWebsite, err error) {
	err = global.GVA_DB.Order("id desc").Find(&websites).Error
	return
}
