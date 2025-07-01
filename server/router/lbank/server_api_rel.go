package lbank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ServerApiRelRouter struct {}

// InitServerApiRelRouter 初始化 做市API操作控制管理 路由信息
func (s *ServerApiRelRouter) InitServerApiRelRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	serverApiRelRouter := Router.Group("serverApiRel").Use(middleware.OperationRecord())
	serverApiRelRouterWithoutRecord := Router.Group("serverApiRel")
	serverApiRelRouterWithoutAuth := PublicRouter.Group("serverApiRel")
	{
		serverApiRelRouter.POST("createServerApiRel", serverApiRelApi.CreateServerApiRel)   // 新建做市API操作控制管理
		serverApiRelRouter.DELETE("deleteServerApiRel", serverApiRelApi.DeleteServerApiRel) // 删除做市API操作控制管理
		serverApiRelRouter.DELETE("deleteServerApiRelByIds", serverApiRelApi.DeleteServerApiRelByIds) // 批量删除做市API操作控制管理
		serverApiRelRouter.PUT("updateServerApiRel", serverApiRelApi.UpdateServerApiRel)    // 更新做市API操作控制管理
	}
	{
		serverApiRelRouterWithoutRecord.GET("findServerApiRel", serverApiRelApi.FindServerApiRel)        // 根据ID获取做市API操作控制管理
		serverApiRelRouterWithoutRecord.GET("getServerApiRelList", serverApiRelApi.GetServerApiRelList)  // 获取做市API操作控制管理列表
	}
	{
	    serverApiRelRouterWithoutAuth.GET("getServerApiRelPublic", serverApiRelApi.GetServerApiRelPublic)  // 做市API操作控制管理开放接口
	}
}
