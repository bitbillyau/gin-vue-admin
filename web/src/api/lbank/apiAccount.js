import service from '@/utils/request'
// @Tags ApiAccount
// @Summary 创建做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApiAccount true "创建做市账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /apiAccount/createApiAccount [post]
export const createApiAccount = (data) => {
  return service({
    url: '/apiAccount/createApiAccount',
    method: 'post',
    data
  })
}

// @Tags ApiAccount
// @Summary 删除做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApiAccount true "删除做市账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiAccount/deleteApiAccount [delete]
export const deleteApiAccount = (params) => {
  return service({
    url: '/apiAccount/deleteApiAccount',
    method: 'delete',
    params
  })
}

// @Tags ApiAccount
// @Summary 批量删除做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除做市账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /apiAccount/deleteApiAccount [delete]
export const deleteApiAccountByIds = (params) => {
  return service({
    url: '/apiAccount/deleteApiAccountByIds',
    method: 'delete',
    params
  })
}

// @Tags ApiAccount
// @Summary 更新做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body model.ApiAccount true "更新做市账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /apiAccount/updateApiAccount [put]
export const updateApiAccount = (data) => {
  return service({
    url: '/apiAccount/updateApiAccount',
    method: 'put',
    data
  })
}

// @Tags ApiAccount
// @Summary 用id查询做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query model.ApiAccount true "用id查询做市账户管理"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /apiAccount/findApiAccount [get]
export const findApiAccount = (params) => {
  return service({
    url: '/apiAccount/findApiAccount',
    method: 'get',
    params
  })
}

// @Tags ApiAccount
// @Summary 分页获取做市账户管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取做市账户管理列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /apiAccount/getApiAccountList [get]
export const getApiAccountList = (params) => {
  return service({
    url: '/apiAccount/getApiAccountList',
    method: 'get',
    params
  })
}

// @Tags ApiAccount
// @Summary 不需要鉴权的做市账户管理接口
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.ApiAccountSearch true "分页获取做市账户管理列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /apiAccount/getApiAccountPublic [get]
export const getApiAccountPublic = () => {
  return service({
    url: '/apiAccount/getApiAccountPublic',
    method: 'get',
  })
}
