package business

import (
	"github.com/gin-gonic/gin"
)

type OperationRouter struct{}

func (s *OperationRouter) InitOperationRouter(Router *gin.RouterGroup) {
	operationRouter := Router.Group("operation")
	{
		operationRouter.GET("statistics/count", operationApi.GetStatisticsCount) // 获取会员、充值、金币等统计信息
		// operationRouter.GET("deposit/amount", operationApi.GetDepositAmount) // 获取充值信息

		operationRouter.POST("update/all/user/asset", operationApi.UpdateAllUserAssetTotalInfo) // 更新所有用户资产总信息
		operationRouter.POST("update/report", operationApi.UpdateOperationReport)               // 更新运营报表

		operationRouter.GET("chart/member/count", operationApi.GetMemberCountChart) // 获取会员数量图表数据
		operationRouter.GET("chart/depositusdt", operationApi.GetDepositUsdtChart)  // 获取充值图表数据
		operationRouter.GET("chart/coinuse", operationApi.GetCoinUseChart)          // 获取金币消费图表数据

		operationRouter.GET("chart/country/member", operationApi.GetCountryMemberChart)           // 获取国家会员数图表数据
		operationRouter.GET("chart/allcountry/member", operationApi.GetAllCountryMemberChart)     // 获取所有国家会员数图表数据
		operationRouter.GET("chart/member/depositusdt", operationApi.GetMemberDepositUsdtChart)   // 获取会员充值Usdt图表数据
		operationRouter.GET("chart/member/invite", operationApi.GetMemberInviteChart)             // 获取用户邀请图表数据
		operationRouter.GET("chart/member/rebate", operationApi.GetMemberRebateChart)             // 获取用户返利图表数据
		operationRouter.GET("chart/country/depositusdt", operationApi.GetCountryDepositUsdtChart) // 获取国家充值图表数据

		operationRouter.GET("chart/divination/percent", operationApi.GetDivinationPercentChart) // 获取塔罗占卜百分比

	}
}
