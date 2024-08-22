package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type WebsiteCategoryRouter struct{}

func (s *WebsiteCategoryRouter) InitWebsiteCategoryRouter(Router *gin.RouterGroup) {
	modeCategoryRouter := Router.Group("website/category").Use(middleware.OperationRecord())
	modeCategoryRouterWithoutRecord := Router.Group("website/category")
	{

		modeCategoryRouter.POST("createWebsiteCategory", busWebsiteCategoryApi.CreateWebsiteCategory)               // 创建网站分类
		modeCategoryRouter.POST("deleteWebsiteCategory", busWebsiteCategoryApi.DeleteWebsiteCategory)               // 删除网站分类
		modeCategoryRouter.POST("updateWebsiteCategory", busWebsiteCategoryApi.UpdateWebsiteCategory)               // 更新网站分类信息
		modeCategoryRouter.DELETE("deleteWebsiteCategorysByIds", busWebsiteCategoryApi.DeleteWebsiteCategorysByIds) // 删除选中网站分类
	}
	{
		modeCategoryRouterWithoutRecord.POST("getWebsiteCategoryById", busWebsiteCategoryApi.GetWebsiteCategoryById) // 获取单条网站分类信息
		modeCategoryRouterWithoutRecord.POST("getAllWebsiteCategorys", busWebsiteCategoryApi.GetAllWebsiteCategorys) // 获取所有网站分类(不分页)
		modeCategoryRouterWithoutRecord.POST("getWebsiteCategoryList", busWebsiteCategoryApi.GetWebsiteCategoryList) // 获取网站分类列表(分页)
	}
}
