package business

import (
	"github.com/billbliu/tarot-admin/global"
	"github.com/billbliu/tarot-admin/model/business/request"
	"github.com/billbliu/tarot-admin/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OperationApi struct{}

// GetStatisticsCount
// @Tags      OperationApi
// @Summary   会员、充值、金币等统计信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "会员、充值、金币等统计信息"
// @Router    /operation/statistics/count [get]
func (s *OperationApi) GetStatisticsCount(c *gin.Context) {
	res, err := operationService.GetStatisticsCount(c)
	if err != nil {
		global.GVA_LOG.Error("GetStatisticsCount err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

/*
// GetDepositAmount
// @Tags      OperationApi
// @Summary   tarot充值数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "tarot充值数据"
// @Router    /operation/deposit/amount [get]
*/
func (s *OperationApi) GetDepositAmount(c *gin.Context) {
	res, err := operationService.GetDepositAmount()
	if err != nil {
		global.GVA_LOG.Error("GetDepositAmount err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

/*
// UpdateAllUserAssetTotalInfo
// @Tags      OperationApi
// @Summary   更新所有用户资产总值
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "更新所有用户资产总值"
// @Router    /operation/update/all/user/asset [post]
*/
func (s *OperationApi) UpdateAllUserAssetTotalInfo(c *gin.Context) {
	err := operationService.UpdateAllUserAssetTotalInfo()
	if err != nil {
		global.GVA_LOG.Error("GetDepositAmount err", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

/*
// UpdateOperationReport
// @Tags      OperationApi
// @Summary   更新运营报表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param start query string true "开始时间"
// @Param end query string true "结束时间"
// @Success   200   {object}  response.Response{msg=string}  "更新运营报表"
// @Router    /operation/update/report [post]
*/
func (s *OperationApi) UpdateOperationReport(c *gin.Context) {
	var req request.UpdateOperationReportRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = operationService.UpdateOperationReport(req)
	if err != nil {
		global.GVA_LOG.Error("UpdateOperationReport err", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetMemberCountChart
// @Tags      OperationApi
// @Summary   会员数量图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     days query int false "取多少天数据"
// @Success   200   {object}  response.Response{msg=string}  "会员数量图表数据"
// @Router    /operation/chart/member/count [get]
func (s *OperationApi) GetMemberCountChart(c *gin.Context) {
	var req request.GetMemberCountChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetMemberCountChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetStatisticsCount err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetDepositUsdtChart
// @Tags      OperationApi
// @Summary   充值图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     days query int false "取多少天数据"
// @Success   200   {object}  response.Response{msg=string}  "充值图表数据"
// @Router    /operation/chart/depositusdt [get]
func (s *OperationApi) GetDepositUsdtChart(c *gin.Context) {
	var req request.GetDepositUsdtChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetDepositUsdtChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetDepositUsdtChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetCoinUseChart
// @Tags      OperationApi
// @Summary   金币消费图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     days query int false "取多少天数据"
// @Success   200   {object}  response.Response{msg=string}  "金币消费图表数据"
// @Router    /operation/chart/coinuse [get]
func (s *OperationApi) GetCoinUseChart(c *gin.Context) {
	var req request.GetCoinUseChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetCoinUseChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetCoinUseChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetCountryMemberChart
// @Tags      OperationApi
// @Summary   国家会员数图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     top_several query int false "取top多少数据"
// @Success   200   {object}  response.Response{msg=string}  "国家会员数图表数据"
// @Router    /operation/chart/country/member [get]
func (s *OperationApi) GetCountryMemberChart(c *gin.Context) {
	var req request.GetCountryMemberChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetCountryMemberChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetCountryMemberChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetAllCountryMemberChart
// @Tags      OperationApi
// @Summary   获取所有国家会员数图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "获取所有国家会员数图表数据"
// @Router    /operation/chart/allcountry/member [get]
func (s *OperationApi) GetAllCountryMemberChart(c *gin.Context) {
	res, err := operationService.GetAllCountryMemberChart()
	if err != nil {
		global.GVA_LOG.Error("GetAllCountryMemberChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetMemberDepositUsdtChart
// @Tags      OperationApi
// @Summary   会员充值Usdt图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     top_several query int false "取top多少数据"
// @Success   200   {object}  response.Response{msg=string}  "会员充值Usdt图表数据"
// @Router    /operation/chart/member/depositusdt [get]
func (s *OperationApi) GetMemberDepositUsdtChart(c *gin.Context) {
	var req request.GetMemberDepositUsdtChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetMemberDepositUsdtChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetMemberDepositUsdtChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetMemberInviteChart
// @Tags      OperationApi
// @Summary   获取用户邀请图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     top_several query int false "取top多少数据"
// @Success   200   {object}  response.Response{msg=string}  "获取用户邀请图表数据"
// @Router    /operation/chart/member/invite [get]
func (s *OperationApi) GetMemberInviteChart(c *gin.Context) {
	var req request.GetMemberInviteChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetMemberInviteChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetMemberInviteChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetMemberRebateChart
// @Tags      OperationApi
// @Summary   获取用户返利图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     top_several query int false "取top多少数据"
// @Success   200   {object}  response.Response{msg=string}  "获取用户返利图表数据"
// @Router    /operation/chart/member/rebate [get]
func (s *OperationApi) GetMemberRebateChart(c *gin.Context) {
	var req request.GetMemberRebateChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetMemberRebateChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetMemberRebateChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetCountryDepositUsdtChart
// @Tags      OperationApi
// @Summary   获取国家充值图表数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     top_several query int false "取top多少数据"
// @Success   200   {object}  response.Response{msg=string}  "获取国家充值图表数据"
// @Router    /operation/chart/country/depositusdt [get]
func (s *OperationApi) GetCountryDepositUsdtChart(c *gin.Context) {
	var req request.GetCountryDepositUsdtChartRequest
	err := c.ShouldBind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res, err := operationService.GetCountryDepositUsdtChart(req)
	if err != nil {
		global.GVA_LOG.Error("GetCountryDepositUsdtChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}

// GetDivinationPercentChart
// @Tags      OperationApi
// @Summary   获取塔罗占卜百分比
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "获取塔罗占卜百分比"
// @Router    /operation/chart/divination/percent [get]
func (s *OperationApi) GetDivinationPercentChart(c *gin.Context) {
	res, err := operationService.GetDivinationPercentChart()
	if err != nil {
		global.GVA_LOG.Error("GetDivinationPercentChart err", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(res, c)
}
