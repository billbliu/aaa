package business

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) {
	apiRouterWithoutRecord := Router.Group("base")
	{
		apiRouterWithoutRecord.GET("getEmailSMS", busBaseApi.GetEmailSMS) // 获取邮箱验证码
		apiRouterWithoutRecord.GET("getPhoneSMS", busBaseApi.GetPhoneSMS) // 获取手机验证码

		apiRouterWithoutRecord.POST("loginOrRegisterByEmailSMS", busBaseApi.LoginOrRegisterByEmailSMS) // 通过邮箱验证码登录或注册
		apiRouterWithoutRecord.POST("loginOrRegisterByPhoneSMS", busBaseApi.LoginOrRegisterByPhoneSMS) // 通过手机验证码登录或注册
		// apiRouterWithoutRecord.POST("loginByPassword", busBaseApi.LoginByPassword)                     // 通过用户名密码登录
	}
}
