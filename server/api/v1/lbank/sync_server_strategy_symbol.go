package lbank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SyncServerStrategySymbolApi struct{}

// SyncServerStrategySymbol 同步server_strategy_symbol
// @Tags      Lbank
// @Summary   同步server_strategy_symbol
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "同步成功"
// @Router    /db/syncServerStrategySymbol [post]
func (s *SyncServerStrategySymbolApi) SyncServerStrategySymbol(c *gin.Context) {
	// 获取数据库连接
	db := global.GetGlobalDBByDBName("daemon")
	if db == nil {
		global.GVA_LOG.Error("daemonDB连接不存在")
		response.FailWithMessage("daemonDB连接失败", c)
		return
	}

	// 执行插入SQL
	err := db.Exec(`
	INSERT INTO server_strategy_symbol 
		(server_id, strategy_symbol_id, status)
	SELECT 
		s.server_id,
		ssr.id,
		ssr.status
	FROM server_strategy_rel ssr
	CROSS JOIN (SELECT 1 AS server_id UNION SELECT 2) s
	WHERE ssr.status IS NOT NULL
		AND NOT EXISTS (
			SELECT 1
			FROM server_strategy_symbol sss
			WHERE sss.strategy_symbol_id = ssr.id AND sss.server_id = s.server_id
    )`,
	).Error

	if err != nil {
		global.GVA_LOG.Error("插入数据失败", zap.Error(err))
		response.FailWithMessage("插入数据失败: "+err.Error(), c)
		return
	}

	global.GVA_LOG.Info("向server_strategy_symbol表插入数据成功")
	response.OkWithMessage("同步成功", c)
}
