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

type BusModeApi struct{}

// CreateMode
// @Tags      BusModeApi
// @Summary   创建模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.CreateModeReq    true  "模式信息"
// @Success   200   {object}  response.Response{msg=string}  "创建模式信息"
// @Router    /mode/createMode [post]
func (bus *BusModeApi) CreateMode(c *gin.Context) {
	var req businessReq.CreateModeReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = modeService.CreateMode(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteMode
// @Tags      BusModeApi
// @Summary   删除模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                  true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除模式信息"
// @Router    /mode/deleteMode [post]
func (s *BusModeApi) DeleteMode(c *gin.Context) {
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

	err = modeService.DeleteMode(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetModeById
// @Tags      BusModeApi
// @Summary   根据id获取模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                   true  "根据id获取模式信息"
// @Success   200   {object}  response.Response{data=business.BusMode}      "根据id获取模式信息"
// @Router    /mode/getModeById [post]
func (s *BusModeApi) GetModeById(c *gin.Context) {
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
	website, err := modeService.GetModeById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(website, "获取成功", c)
}

// GetModeList
// @Tags      BusModeApi
// @Summary   分页获取模式列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.SearchModeParams                              true  "分页获取模式列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取模式列表,返回包括列表,总数,页码,每页数量"
// @Router    /mode/getModeList [post]
func (s *BusModeApi) GetModeList(c *gin.Context) {
	var pageInfo businessReq.SearchModeParams
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
	list, total, err := modeService.GetModeList(pageInfo.PageInfo)
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

// UpdateMode
// @Tags      BusModeApi
// @Summary   修改模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.UpdateModeReq   true  "修改模式信息"
// @Success   200   {object}  response.Response{msg=string}  "修改模式信息"
// @Router    /mode/updateMode [post]
func (s *BusModeApi) UpdateMode(c *gin.Context) {
	var req businessReq.UpdateModeReq
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
	err = modeService.UpdateMode(req)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetAllModes
// @Tags      BusModeApi
// @Summary   获取所有的模式 不分页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=[]business.BusMode,msg=string}  "获取所有的模式 不分页"
// @Router    /mode/getAllModes [post]
func (s *BusModeApi) GetAllModes(c *gin.Context) {
	websites, err := modeService.GetAllModes()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(websites, "获取成功", c)
}

// DeleteModesByIds
// @Tags      BusModeApi
// @Summary   删除选中Api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除选中Api"
// @Router    /mode/deleteModesByIds [delete]
func (s *BusModeApi) DeleteModesByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = modeService.DeleteModesByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// ModeBindingCategoryWebsites
// @Tags      BusModeApi
// @Summary   模式绑定分类、网站
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.ModeBindingCategoryWebsitesReq   true  "模式绑定分类、网站"
// @Success   200   {object}  response.Response{msg=string}  "模式绑定分类、网站"
// @Router    /mode/modeBindingCategoryWebsites [post]
func (s *BusModeApi) ModeBindingCategoryWebsites(c *gin.Context) {
	var req businessReq.ModeBindingCategoryWebsitesReq
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
	err = modeService.ModeBindingCategoryWebsites(req)
	if err != nil {
		global.GVA_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithMessage("绑定失败", c)
		return
	}
	response.OkWithMessage("绑定成功", c)
}

// GetModeBindingCategoryWebsites
// @Tags      BusModeApi
// @Summary    获取模式已绑定分类、网站
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=[]business.BusMode,msg=string}  " 获取模式已绑定分类、网站"
// @Router    /mode/getModeBindingWebsites [get]
func (s *BusModeApi) GetModeBindingCategoryWebsites(c *gin.Context) {
	var idInfo request.GetById
	err := c.ShouldBindQuery(&idInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(idInfo, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	websites, err := modeService.GetModeBindingCategoryWebsites(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(websites, "获取成功", c)
}
