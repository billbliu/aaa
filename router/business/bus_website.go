package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WebsiteRouter struct{}

func (s *WebsiteRouter) InitWebsiteRouter(Router *gin.RouterGroup) {
	apiRouter := Router.Group("website").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("website")
	{

		apiRouter.POST("createWebsite", busWebsiteApi.CreateWebsite)               // 创建网站
		apiRouter.POST("deleteWebsite", busWebsiteApi.DeleteWebsite)               // 删除网站
		apiRouter.POST("updateWebsite", busWebsiteApi.UpdateWebsite)               // 更新网站信息
		apiRouter.DELETE("deleteWebsitesByIds", busWebsiteApi.DeleteWebsitesByIds) // 删除选中网站
	}
	{
		apiRouterWithoutRecord.POST("getWebsiteById", busWebsiteApi.GetWebsiteById) // 获取单条网站信息
		apiRouterWithoutRecord.POST("getAllWebsites", busWebsiteApi.GetAllWebsites) // 获取所有网站(不分页)
		apiRouterWithoutRecord.POST("getWebsiteList", busWebsiteApi.GetWebsiteList) // 获取网站列表(分页)
	}
}
