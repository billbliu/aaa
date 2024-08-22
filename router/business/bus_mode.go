package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ModeRouter struct{}

func (s *ModeRouter) InitModeRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("mode").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("mode")
	{
		apiRouter.POST("createMode", busModeApi.CreateMode)                                   // 创建模式
		apiRouter.POST("deleteMode", busModeApi.DeleteMode)                                   // 删除模式
		apiRouter.POST("updateMode", busModeApi.UpdateMode)                                   // 更新模式信息
		apiRouter.DELETE("deleteModesByIds", busModeApi.DeleteModesByIds)                     // 删除选中模式
		apiRouter.POST("modeBindingCategoryWebsites", busModeApi.ModeBindingCategoryWebsites) // 模式绑定分类、网站
	}
	{
		apiRouterWithoutRecord.POST("getModeById", busModeApi.GetModeById)                                      // 获取单条模式信息
		apiRouterWithoutRecord.POST("getAllModes", busModeApi.GetAllModes)                                      // 获取所有模式(不分页)
		apiRouterWithoutRecord.POST("getModeList", busModeApi.GetModeList)                                      // 获取模式列表(分页)
		apiRouterWithoutRecord.GET("getModeBindingCategoryWebsites", busModeApi.GetModeBindingCategoryWebsites) // 获取模式已绑定分类、网站
	}
}
