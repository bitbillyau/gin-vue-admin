import service from '@/utils/request'
// @Tags ApiSubRel
// @Summary 创建做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApiSubRel true "创建做市子账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /apiSubRel/createApiSubRel [post]
export const createApiSubRel = (data) => {
  return service({
    url: '/apiSubRel/createApiSubRel',
    method: 'post',
    data
  })
}

// @Tags ApiSubRel
// @Summary 删除做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApiSubRel true "删除做市子账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiSubRel/deleteApiSubRel [delete]
export const deleteApiSubRel = (params) => {
  return service({
    url: '/apiSubRel/deleteApiSubRel',
    method: 'delete',
    params
  })
}

// @Tags ApiSubRel
// @Summary 批量删除做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除做市子账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiSubRel/deleteApiSubRel [delete]
export const deleteApiSubRelByIds = (params) => {
  return service({
    url: '/apiSubRel/deleteApiSubRelByIds',
    method: 'delete',
    params
  })
}

// @Tags ApiSubRel
// @Summary 更新做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApiSubRel true "更新做市子账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiSubRel/updateApiSubRel [put]
export const updateApiSubRel = (data) => {
  return service({
    url: '/apiSubRel/updateApiSubRel',
    method: 'put',
    data
  })
}

// @Tags ApiSubRel
// @Summary 用id查询做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ApiSubRel true "用id查询做市子账户"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiSubRel/findApiSubRel [get]
export const findApiSubRel = (params) => {
  return service({
    url: '/apiSubRel/findApiSubRel',
    method: 'get',
    params
  })
}

// @Tags ApiSubRel
// @Summary 分页获取做市子账户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取做市子账户列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiSubRel/getApiSubRelList [get]
export const getApiSubRelList = (params) => {
  return service({
    url: '/apiSubRel/getApiSubRelList',
    method: 'get',
    params
  })
}

// @Tags ApiSubRel
// @Summary 不需要鉴权的做市子账户接口
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.ApiSubRelSearch true "分页获取做市子账户列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /apiSubRel/getApiSubRelPublic [get]
export const getApiSubRelPublic = () => {
  return service({
    url: '/apiSubRel/getApiSubRelPublic',
    method: 'get',
  })
}
