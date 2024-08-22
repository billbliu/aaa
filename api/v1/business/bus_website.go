package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BusWebsiteApi struct{}

// CreateWebsite
// @Tags      BusWebsiteApi
// @Summary   创建网站信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.CreateWebsiteReq    true  "网站信息"
// @Success   200   {object}  response.Response{msg=string}  "创建网站信息"
// @Router    /website/createWebsite [post]
func (bus *BusWebsiteApi) CreateWebsite(c *gin.Context) {
	var req businessReq.CreateWebsiteReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.CreateWebsite(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWebsite
// @Tags      BusWebsiteApi
// @Summary   删除网站信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                  true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除网站信息"
// @Router    /website/deleteWebsite [post]
func (s *BusWebsiteApi) DeleteWebsite(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteService.DeleteWebsite(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetWebsiteById
// @Tags      BusWebsiteApi
// @Summary   根据id获取网站信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                   true  "根据id获取网站信息"
// @Success   200   {object}  response.Response{data=business.BusWebsite}      "根据id获取网站信息"
// @Router    /website/getWebsiteById [post]
func (s *BusWebsiteApi) GetWebsiteById(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindJSON(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	website, err := websiteService.GetWebsiteById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(website, "获取成功", c)
}

// GetWebsiteList
// @Tags      BusWebsiteApi
// @Summary   分页获取网站列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.SearchWebsiteParams                              true  "分页获取网站列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取网站列表,返回包括列表,总数,页码,每页数量"
// @Router    /website/getWebsiteList [post]
func (s *BusWebsiteApi) GetWebsiteList(c *gin.Context) {
	var pageInfo businessReq.SearchWebsiteParams
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo.PageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := websiteService.GetWebsiteList(pageInfo.PageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// UpdateWebsite
// @Tags      BusWebsiteApi
// @Summary   修改网站信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.UpdateWebsiteReq   true  "修改网站信息"
// @Success   200   {object}  response.Response{msg=string}  "修改网站信息"
// @Router    /website/updateWebsite [post]
func (s *BusWebsiteApi) UpdateWebsite(c *gin.Context) {
	var req businessReq.UpdateWebsiteReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = websiteService.UpdateWebsite(req)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetAllWebsites
// @Tags      BusWebsiteApi
// @Summary   获取所有的网站 不分页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=[]business.BusWebsite,msg=string}  "获取所有的网站 不分页"
// @Router    /website/getAllWebsites [post]
func (s *BusWebsiteApi) GetAllWebsites(c *gin.Context) {
	websites, err := websiteService.GetAllWebsites()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(websites, "获取成功", c)
}

// DeleteWebsitesByIds
// @Tags      BusWebsiteApi
// @Summary   删除选中Api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除选中Api"
// @Router    /website/deleteWebsitesByIds [delete]
func (s *BusWebsiteApi) DeleteWebsitesByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = websiteService.DeleteWebsitesByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
