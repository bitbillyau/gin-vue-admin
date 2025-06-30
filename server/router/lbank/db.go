package lbank

import (
	"github.com/gin-gonic/gin"
)

type DbRouter struct{}

func (d *DbRouter) InitDbRouter(Router *gin.RouterGroup) {
	dbRouter := Router.Group("db")
	{
		dbRouter.POST("syncServerStrategySymbol", dbPageApi.SyncServerStrategySymbol)              // 同步server_strategy_symbol
		dbRouter.GET("getSymbolsForLinkedLbankAccount", dbPageApi.GetSymbolsForLinkedLbankAccount) // 获取可用于添加Lbank账号的symbol列表
		dbRouter.GET("getLbankAccounts", dbPageApi.GetLbankAccounts)                               // 获取Lbank账号列表
		dbRouter.POST("submitLinkedLbankSubAccounts", dbPageApi.SubmitLinkedLbankSubAccounts)      // 关联母账号-子账号提交
	}
}
