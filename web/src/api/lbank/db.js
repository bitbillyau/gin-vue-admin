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

// @Tags Lbank
// @Summary 获取symbol列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]object,msg=string} "获取成功"
// @Router /db/getSymbolsForAddLbankAccount [get]
export const getSymbolsForAddLbankAccount = () => {
  return service({
    url: '/db/getSymbolsForAddLbankAccount',
    method: 'get'
  })
}

// @Tags Lbank
// @Summary 获取Lbank账号列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=[]object,msg=string} "获取成功"
// @Router /db/getLbankAccounts [get]
export const getLbankAccounts = () => {
  return service({
    url: '/db/getLbankAccounts',
    method: 'get'
  })
} 