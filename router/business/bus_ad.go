package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdRouter struct{}

func (s *AdRouter) InitAdRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("ad").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("ad")
	{
		apiRouter.POST("", busAdApi.CreateAd)                       // 创建广告
		apiRouter.DELETE(":id", busAdApi.DeleteAd)                  // 删除广告
		apiRouter.POST("updateAd", busAdApi.UpdateAd)               // 更新广告信息
		apiRouter.DELETE("deleteAdsByIds", busAdApi.DeleteAdsByIds) // 删除选中广告

		apiRouter.POST("", busAdApi.CreateAdLevel)                       // 创建广告等级
		apiRouter.DELETE(":id", busAdApi.DeleteAdLevel)                  // 删除广告等级
		apiRouter.POST("updateAd", busAdApi.UpdateAdLevel)               // 更新广告等级
		apiRouter.DELETE("deleteAdsByIds", busAdApi.DeleteAdLevelsByIds) // 删除选中广告等级
	}
	{
		apiRouterWithoutRecord.GET(":id", busAdApi.GetAdById)  // 获取单条广告信息
		apiRouterWithoutRecord.GET("all", busAdApi.GetAllAds)  // 获取所有广告(不分页)
		apiRouterWithoutRecord.GET("list", busAdApi.GetAdList) // 获取广告列表(分页)

		apiRouterWithoutRecord.GET(":id", busAdApi.GetAdLevelById) // 获取单条广告等级
		apiRouterWithoutRecord.GET("all", busAdApi.GetAllAdLevels) // 获取所有广告等级(不分页)
	}
}
