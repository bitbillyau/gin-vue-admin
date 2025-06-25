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
	// 获取hq数据库连接
	hqDB := global.GetGlobalDBByDBName("daemon")
	if hqDB == nil {
		global.GVA_LOG.Error("hq数据库连接不存在")
		response.FailWithMessage("数据库连接失败", c)
		return
	}

	// 执行插入SQL
	sql := "INSERT INTO test (name, date) VALUES ('gin-vue-admin', NOW())"
	err := hqDB.Exec(sql).Error
	if err != nil {
		global.GVA_LOG.Error("插入数据失败", zap.Error(err))
		response.FailWithMessage("插入数据失败: "+err.Error(), c)
		return
	}

	global.GVA_LOG.Info("成功向hq.test表插入记录")
	response.OkWithMessage("同步成功", c)
}
