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

type BusWebsiteCategoryApi struct{}

// CreateWebsiteCategory
// @Tags      BusWebsiteCategoryApi
// @Summary   创建网站分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.CreateWebsiteCategoryReq    true  "网站分类"
// @Success   200   {object}  response.Response{msg=string}  "创建网站分类"
// @Router    /website/category/createWebsiteCategory [post]
func (bus *BusWebsiteCategoryApi) CreateWebsiteCategory(c *gin.Context) {
	var req businessReq.CreateWebsiteCategoryReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = websiteCategoryService.CreateWebsiteCategory(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteWebsiteCategory
// @Tags      BusWebsiteCategoryApi
// @Summary   删除网站分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                  true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除网站分类"
// @Router    /website/category/deleteWebsiteCategory [post]
func (s *BusWebsiteCategoryApi) DeleteWebsiteCategory(c *gin.Context) {
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

	err = websiteCategoryService.DeleteWebsiteCategory(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetWebsiteCategoryById
// @Tags      BusWebsiteCategoryApi
// @Summary   根据id获取网站分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                   true  "根据id获取网站分类"
// @Success   200   {object}  response.Response{data=business.BusWebsiteCategory}      "根据id获取网站分类"
// @Router    /website/category/getWebsiteCategoryById [post]
func (s *BusWebsiteCategoryApi) GetWebsiteCategoryById(c *gin.Context) {
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
	website, err := websiteCategoryService.GetWebsiteCategoryById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(website, "获取成功", c)
}

// GetWebsiteCategoryList
// @Tags      BusWebsiteCategoryApi
// @Summary   分页获取网站分类列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.SearchWebsiteCategoryParams                               true  "分页获取网站分类列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取网站分类列表,返回包括列表,总数,页码,每页数量"
// @Router    /website/category/getWebsiteCategoryList [post]
func (s *BusWebsiteCategoryApi) GetWebsiteCategoryList(c *gin.Context) {
	var pageInfo businessReq.SearchWebsiteCategoryParams
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
	list, total, err := websiteCategoryService.GetWebsiteCategoryList(pageInfo.PageInfo)
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

// UpdateWebsiteCategory
// @Tags      BusWebsiteCategoryApi
// @Summary   修改网站分类
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.UpdateWebsiteCategoryReq   true  "修改网站分类"
// @Success   200   {object}  response.Response{msg=string}  "修改网站分类"
// @Router    /website/category/updateWebsiteCategory [post]
func (s *BusWebsiteCategoryApi) UpdateWebsiteCategory(c *gin.Context) {
	var req businessReq.UpdateWebsiteCategoryReq
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
	err = websiteCategoryService.UpdateWebsiteCategory(req)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetAllWebsiteCategorys
// @Tags      BusWebsiteCategoryApi
// @Summary   获取所有的网站分类 不分页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=[]business.BusWebsiteCategory,msg=string}  "获取所有的网站分类 不分页"
// @Router    /website/category/getAllWebsiteCategorys [post]
func (s *BusWebsiteCategoryApi) GetAllWebsiteCategorys(c *gin.Context) {
	websites, err := websiteCategoryService.GetAllWebsiteCategorys()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(websites, "获取成功", c)
}

// DeleteWebsiteCategorysByIds
// @Tags      BusWebsiteCategoryApi
// @Summary   删除选中Api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除选中Api"
// @Router    /website/category/deleteWebsiteCategorysByIds [delete]
func (s *BusWebsiteCategoryApi) DeleteWebsiteCategorysByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = websiteCategoryService.DeleteWebsiteCategorysByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
