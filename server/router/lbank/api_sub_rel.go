package lbank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiSubRelRouter struct {}

// InitApiSubRelRouter 初始化 做市子账户 路由信息
func (s *ApiSubRelRouter) InitApiSubRelRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	apiSubRelRouter := Router.Group("apiSubRel").Use(middleware.OperationRecord())
	apiSubRelRouterWithoutRecord := Router.Group("apiSubRel")
	apiSubRelRouterWithoutAuth := PublicRouter.Group("apiSubRel")
	{
		apiSubRelRouter.POST("createApiSubRel", apiSubRelApi.CreateApiSubRel)   // 新建做市子账户
		apiSubRelRouter.DELETE("deleteApiSubRel", apiSubRelApi.DeleteApiSubRel) // 删除做市子账户
		apiSubRelRouter.DELETE("deleteApiSubRelByIds", apiSubRelApi.DeleteApiSubRelByIds) // 批量删除做市子账户
		apiSubRelRouter.PUT("updateApiSubRel", apiSubRelApi.UpdateApiSubRel)    // 更新做市子账户
	}
	{
		apiSubRelRouterWithoutRecord.GET("findApiSubRel", apiSubRelApi.FindApiSubRel)        // 根据ID获取做市子账户
		apiSubRelRouterWithoutRecord.GET("getApiSubRelList", apiSubRelApi.GetApiSubRelList)  // 获取做市子账户列表
	}
	{
	    apiSubRelRouterWithoutAuth.GET("getApiSubRelPublic", apiSubRelApi.GetApiSubRelPublic)  // 做市子账户开放接口
	}
}
