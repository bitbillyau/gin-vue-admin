package lbank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ApiAccountRouter struct {}

// InitApiAccountRouter 初始化 做市账户管理 路由信息
func (s *ApiAccountRouter) InitApiAccountRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	apiAccountRouter := Router.Group("apiAccount").Use(middleware.OperationRecord())
	apiAccountRouterWithoutRecord := Router.Group("apiAccount")
	apiAccountRouterWithoutAuth := PublicRouter.Group("apiAccount")
	{
		apiAccountRouter.POST("createApiAccount", apiAccountApi.CreateApiAccount)   // 新建做市账户管理
		apiAccountRouter.DELETE("deleteApiAccount", apiAccountApi.DeleteApiAccount) // 删除做市账户管理
		apiAccountRouter.DELETE("deleteApiAccountByIds", apiAccountApi.DeleteApiAccountByIds) // 批量删除做市账户管理
		apiAccountRouter.PUT("updateApiAccount", apiAccountApi.UpdateApiAccount)    // 更新做市账户管理
	}
	{
		apiAccountRouterWithoutRecord.GET("findApiAccount", apiAccountApi.FindApiAccount)        // 根据ID获取做市账户管理
		apiAccountRouterWithoutRecord.GET("getApiAccountList", apiAccountApi.GetApiAccountList)  // 获取做市账户管理列表
	}
	{
	    apiAccountRouterWithoutAuth.GET("getApiAccountPublic", apiAccountApi.GetApiAccountPublic)  // 做市账户管理开放接口
	}
}
