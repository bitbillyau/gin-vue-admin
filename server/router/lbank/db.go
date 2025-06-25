package lbank

import (
	"github.com/gin-gonic/gin"
)

type DbRouter struct{}

func (d *DbRouter) InitDbRouter(Router *gin.RouterGroup) {
	dbRouter := Router.Group("db")
	{
		dbRouter.POST("syncServerStrategySymbol", syncServerStrategySymbolApi.SyncServerStrategySymbol) // 同步server_strategy_symbol
	}
}
