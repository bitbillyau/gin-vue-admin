package lbank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InstrumentRouter struct {}

// InitInstrumentRouter 初始化 交易对管理 路由信息
func (s *InstrumentRouter) InitInstrumentRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	instrumentRouter := Router.Group("instrument").Use(middleware.OperationRecord())
	instrumentRouterWithoutRecord := Router.Group("instrument")
	instrumentRouterWithoutAuth := PublicRouter.Group("instrument")
	{
		instrumentRouter.POST("createInstrument", instrumentApi.CreateInstrument)   // 新建交易对管理
		instrumentRouter.DELETE("deleteInstrument", instrumentApi.DeleteInstrument) // 删除交易对管理
		instrumentRouter.DELETE("deleteInstrumentByIds", instrumentApi.DeleteInstrumentByIds) // 批量删除交易对管理
		instrumentRouter.PUT("updateInstrument", instrumentApi.UpdateInstrument)    // 更新交易对管理
	}
	{
		instrumentRouterWithoutRecord.GET("findInstrument", instrumentApi.FindInstrument)        // 根据ID获取交易对管理
		instrumentRouterWithoutRecord.GET("getInstrumentList", instrumentApi.GetInstrumentList)  // 获取交易对管理列表
	}
	{
	    instrumentRouterWithoutAuth.GET("getInstrumentPublic", instrumentApi.GetInstrumentPublic)  // 交易对管理开放接口
	}
}
