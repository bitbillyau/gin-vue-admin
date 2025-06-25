import service from '@/utils/request'

// @Tags Lbank
// @Summary 同步server_strategy_symbol
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "同步成功"
// @Router /db/syncServerStrategySymbol [post]
export const syncServerStrategySymbol = () => {
  return service({
    url: '/db/syncServerStrategySymbol',
    method: 'post'
  })
} 