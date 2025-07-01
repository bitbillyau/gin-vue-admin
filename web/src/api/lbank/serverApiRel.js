import service from '@/utils/request'
// @Tags ServerApiRel
// @Summary 创建做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerApiRel true "创建做市API操作控制管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /serverApiRel/createServerApiRel [post]
export const createServerApiRel = (data) => {
  return service({
    url: '/serverApiRel/createServerApiRel',
    method: 'post',
    data
  })
}

// @Tags ServerApiRel
// @Summary 删除做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerApiRel true "删除做市API操作控制管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverApiRel/deleteServerApiRel [delete]
export const deleteServerApiRel = (params) => {
  return service({
    url: '/serverApiRel/deleteServerApiRel',
    method: 'delete',
    params
  })
}

// @Tags ServerApiRel
// @Summary 批量删除做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除做市API操作控制管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /serverApiRel/deleteServerApiRel [delete]
export const deleteServerApiRelByIds = (params) => {
  return service({
    url: '/serverApiRel/deleteServerApiRelByIds',
    method: 'delete',
    params
  })
}

// @Tags ServerApiRel
// @Summary 更新做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ServerApiRel true "更新做市API操作控制管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /serverApiRel/updateServerApiRel [put]
export const updateServerApiRel = (data) => {
  return service({
    url: '/serverApiRel/updateServerApiRel',
    method: 'put',
    data
  })
}

// @Tags ServerApiRel
// @Summary 用id查询做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ServerApiRel true "用id查询做市API操作控制管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /serverApiRel/findServerApiRel [get]
export const findServerApiRel = (params) => {
  return service({
    url: '/serverApiRel/findServerApiRel',
    method: 'get',
    params
  })
}

// @Tags ServerApiRel
// @Summary 分页获取做市API操作控制管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取做市API操作控制管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /serverApiRel/getServerApiRelList [get]
export const getServerApiRelList = (params) => {
  return service({
    url: '/serverApiRel/getServerApiRelList',
    method: 'get',
    params
  })
}

// @Tags ServerApiRel
// @Summary 不需要鉴权的做市API操作控制管理接口
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.ServerApiRelSearch true "分页获取做市API操作控制管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /serverApiRel/getServerApiRelPublic [get]
export const getServerApiRelPublic = () => {
  return service({
    url: '/serverApiRel/getServerApiRelPublic',
    method: 'get',
  })
}
