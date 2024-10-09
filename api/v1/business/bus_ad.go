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

type BusAdApi struct{}

// CreateAd
// @Tags      BusAdApi
// @Summary   创建模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.CreateAdReq    true  "模式信息"
// @Success   200   {object}  response.Response{msg=string}  "创建模式信息"
// @Router    /ad [post]
func (bus *BusAdApi) CreateAd(c *gin.Context) {
	var req businessReq.CreateAdReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = adService.CreateAd(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAd
// @Tags      BusAdApi
// @Summary   删除模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                  true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除模式信息"
// @Router    /mode/deleteAd [post]
func (s *BusAdApi) DeleteAd(c *gin.Context) {
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

	err = adService.DeleteAd(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetAdById
// @Tags      BusAdApi
// @Summary   根据id获取模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                   true  "根据id获取模式信息"
// @Success   200   {object}  response.Response{data=business.BusAd}      "根据id获取模式信息"
// @Router    /mode/getAdById [post]
func (s *BusAdApi) GetAdById(c *gin.Context) {
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
	website, err := adService.GetAdById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(website, "获取成功", c)
}

// GetAdList
// @Tags      BusAdApi
// @Summary   分页获取模式列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.SearchAdParams                              true  "分页获取模式列表"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取模式列表,返回包括列表,总数,页码,每页数量"
// @Router    /mode/getAdList [post]
func (s *BusAdApi) GetAdList(c *gin.Context) {
	var pageInfo businessReq.SearchAdParams
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
	list, total, err := adService.GetAdList(pageInfo.PageInfo)
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

// UpdateAd
// @Tags      BusAdApi
// @Summary   修改模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.UpdateAdReq   true  "修改模式信息"
// @Success   200   {object}  response.Response{msg=string}  "修改模式信息"
// @Router    /mode/updateAd [post]
func (s *BusAdApi) UpdateAd(c *gin.Context) {
	var req businessReq.UpdateAdReq
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
	err = adService.UpdateAd(req)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetAllAds
// @Tags      BusAdApi
// @Summary   获取所有的模式 不分页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=[]business.BusAd,msg=string}  "获取所有的模式 不分页"
// @Router    /mode/getAllAds [post]
func (s *BusAdApi) GetAllAds(c *gin.Context) {
	websites, err := adService.GetAllAds()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(websites, "获取成功", c)
}

// DeleteAdsByIds
// @Tags      BusAdApi
// @Summary   删除选中Api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除选中Api"
// @Router    /mode/deleteAdsByIds [delete]
func (s *BusAdApi) DeleteAdsByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = adService.DeleteAdsByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// CreateAdLevel
// @Tags      BusAdApi
// @Summary   创建模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.CreateAdReq    true  "模式信息"
// @Success   200   {object}  response.Response{msg=string}  "创建模式信息"
// @Router    /ad [post]
func (bus *BusAdApi) CreateAdLevel(c *gin.Context) {
	var req businessReq.CreateAdLevelReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = adService.CreateAdLevel(req)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAdLevel
// @Tags      BusAdApi
// @Summary   删除模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                  true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除模式信息"
// @Router    /mode/deleteAd [post]
func (s *BusAdApi) DeleteAdLevel(c *gin.Context) {
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

	err = adService.DeleteAdLevel(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetAdLevelById
// @Tags      BusAdApi
// @Summary   根据id获取模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById                                   true  "根据id获取模式信息"
// @Success   200   {object}  response.Response{data=business.BusAd}      "根据id获取模式信息"
// @Router    /mode/getAdById [post]
func (s *BusAdApi) GetAdLevelById(c *gin.Context) {
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
	website, err := adService.GetAdLevelById(idInfo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(website, "获取成功", c)
}

// UpdateAdLevel
// @Tags      BusAdApi
// @Summary   修改模式信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      businessReq.UpdateAdReq   true  "修改模式信息"
// @Success   200   {object}  response.Response{msg=string}  "修改模式信息"
// @Router    /mode/updateAd [post]
func (s *BusAdApi) UpdateAdLevel(c *gin.Context) {
	var req businessReq.UpdateAdLevelReq
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
	err = adService.UpdateAdLevel(req)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// GetAllAdLevels
// @Tags      BusAdApi
// @Summary   获取所有的模式 不分页
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200  {object}  response.Response{data=[]business.BusAd,msg=string}  "获取所有的模式 不分页"
// @Router    /mode/getAllAds [post]
func (s *BusAdApi) GetAllAdLevels(c *gin.Context) {
	websites, err := adService.GetAllAdLevels()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(websites, "获取成功", c)
}

// DeleteAdLevelsByIds
// @Tags      BusAdApi
// @Summary   删除选中Api
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.IdsReq                 true  "ID"
// @Success   200   {object}  response.Response{msg=string}  "删除选中Api"
// @Router    /mode/deleteAdsByIds [delete]
func (s *BusAdApi) DeleteAdLevelsByIds(c *gin.Context) {
	var ids request.IdsReq
	err := c.ShouldBindJSON(&ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = adService.DeleteAdLevelsByIds(ids)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
