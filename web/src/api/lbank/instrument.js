import service from '@/utils/request'
// @Tags Instrument
// @Summary 创建交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Instrument true "创建交易对管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /instrument/createInstrument [post]
export const createInstrument = (data) => {
  return service({
    url: '/instrument/createInstrument',
    method: 'post',
    data
  })
}

// @Tags Instrument
// @Summary 删除交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Instrument true "删除交易对管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /instrument/deleteInstrument [delete]
export const deleteInstrument = (params) => {
  return service({
    url: '/instrument/deleteInstrument',
    method: 'delete',
    params
  })
}

// @Tags Instrument
// @Summary 批量删除交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除交易对管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /instrument/deleteInstrument [delete]
export const deleteInstrumentByIds = (params) => {
  return service({
    url: '/instrument/deleteInstrumentByIds',
    method: 'delete',
    params
  })
}

// @Tags Instrument
// @Summary 更新交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.Instrument true "更新交易对管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /instrument/updateInstrument [put]
export const updateInstrument = (data) => {
  return service({
    url: '/instrument/updateInstrument',
    method: 'put',
    data
  })
}

// @Tags Instrument
// @Summary 用id查询交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.Instrument true "用id查询交易对管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /instrument/findInstrument [get]
export const findInstrument = (params) => {
  return service({
    url: '/instrument/findInstrument',
    method: 'get',
    params
  })
}

// @Tags Instrument
// @Summary 分页获取交易对管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取交易对管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /instrument/getInstrumentList [get]
export const getInstrumentList = (params) => {
  return service({
    url: '/instrument/getInstrumentList',
    method: 'get',
    params
  })
}

// @Tags Instrument
// @Summary 不需要鉴权的交易对管理接口
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.InstrumentSearch true "分页获取交易对管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /instrument/getInstrumentPublic [get]
export const getInstrumentPublic = () => {
  return service({
    url: '/instrument/getInstrumentPublic',
    method: 'get',
  })
}
