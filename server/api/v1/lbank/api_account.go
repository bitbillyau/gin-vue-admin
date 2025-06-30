package lbank

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ApiAccountApi struct {}



// CreateApiAccount 创建做市账户管理
// @Tags ApiAccount
// @Summary 创建做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ApiAccount true "创建做市账户管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /apiAccount/createApiAccount [post]
func (apiAccountApi *ApiAccountApi) CreateApiAccount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var apiAccount lbank.ApiAccount
	err := c.ShouldBindJSON(&apiAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiAccountService.CreateApiAccount(ctx,&apiAccount)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteApiAccount 删除做市账户管理
// @Tags ApiAccount
// @Summary 删除做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ApiAccount true "删除做市账户管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /apiAccount/deleteApiAccount [delete]
func (apiAccountApi *ApiAccountApi) DeleteApiAccount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := apiAccountService.DeleteApiAccount(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteApiAccountByIds 批量删除做市账户管理
// @Tags ApiAccount
// @Summary 批量删除做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /apiAccount/deleteApiAccountByIds [delete]
func (apiAccountApi *ApiAccountApi) DeleteApiAccountByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := apiAccountService.DeleteApiAccountByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateApiAccount 更新做市账户管理
// @Tags ApiAccount
// @Summary 更新做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ApiAccount true "更新做市账户管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /apiAccount/updateApiAccount [put]
func (apiAccountApi *ApiAccountApi) UpdateApiAccount(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var apiAccount lbank.ApiAccount
	err := c.ShouldBindJSON(&apiAccount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiAccountService.UpdateApiAccount(ctx,apiAccount)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindApiAccount 用id查询做市账户管理
// @Tags ApiAccount
// @Summary 用id查询做市账户管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询做市账户管理"
// @Success 200 {object} response.Response{data=lbank.ApiAccount,msg=string} "查询成功"
// @Router /apiAccount/findApiAccount [get]
func (apiAccountApi *ApiAccountApi) FindApiAccount(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reapiAccount, err := apiAccountService.GetApiAccount(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reapiAccount, c)
}
// GetApiAccountList 分页获取做市账户管理列表
// @Tags ApiAccount
// @Summary 分页获取做市账户管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.ApiAccountSearch true "分页获取做市账户管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /apiAccount/getApiAccountList [get]
func (apiAccountApi *ApiAccountApi) GetApiAccountList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo lbankReq.ApiAccountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := apiAccountService.GetApiAccountInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetApiAccountPublic 不需要鉴权的做市账户管理接口
// @Tags ApiAccount
// @Summary 不需要鉴权的做市账户管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /apiAccount/getApiAccountPublic [get]
func (apiAccountApi *ApiAccountApi) GetApiAccountPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    apiAccountService.GetApiAccountPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的做市账户管理接口信息",
    }, "获取成功", c)
}
