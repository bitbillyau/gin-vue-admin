package lbank

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ApiSubRelApi struct {}



// CreateApiSubRel 创建做市子账户
// @Tags ApiSubRel
// @Summary 创建做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ApiSubRel true "创建做市子账户"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /apiSubRel/createApiSubRel [post]
func (apiSubRelApi *ApiSubRelApi) CreateApiSubRel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var apiSubRel lbank.ApiSubRel
	err := c.ShouldBindJSON(&apiSubRel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiSubRelService.CreateApiSubRel(ctx,&apiSubRel)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteApiSubRel 删除做市子账户
// @Tags ApiSubRel
// @Summary 删除做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ApiSubRel true "删除做市子账户"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /apiSubRel/deleteApiSubRel [delete]
func (apiSubRelApi *ApiSubRelApi) DeleteApiSubRel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := apiSubRelService.DeleteApiSubRel(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteApiSubRelByIds 批量删除做市子账户
// @Tags ApiSubRel
// @Summary 批量删除做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /apiSubRel/deleteApiSubRelByIds [delete]
func (apiSubRelApi *ApiSubRelApi) DeleteApiSubRelByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := apiSubRelService.DeleteApiSubRelByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateApiSubRel 更新做市子账户
// @Tags ApiSubRel
// @Summary 更新做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ApiSubRel true "更新做市子账户"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /apiSubRel/updateApiSubRel [put]
func (apiSubRelApi *ApiSubRelApi) UpdateApiSubRel(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var apiSubRel lbank.ApiSubRel
	err := c.ShouldBindJSON(&apiSubRel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = apiSubRelService.UpdateApiSubRel(ctx,apiSubRel)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindApiSubRel 用id查询做市子账户
// @Tags ApiSubRel
// @Summary 用id查询做市子账户
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询做市子账户"
// @Success 200 {object} response.Response{data=lbank.ApiSubRel,msg=string} "查询成功"
// @Router /apiSubRel/findApiSubRel [get]
func (apiSubRelApi *ApiSubRelApi) FindApiSubRel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reapiSubRel, err := apiSubRelService.GetApiSubRel(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reapiSubRel, c)
}
// GetApiSubRelList 分页获取做市子账户列表
// @Tags ApiSubRel
// @Summary 分页获取做市子账户列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.ApiSubRelSearch true "分页获取做市子账户列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /apiSubRel/getApiSubRelList [get]
func (apiSubRelApi *ApiSubRelApi) GetApiSubRelList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo lbankReq.ApiSubRelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := apiSubRelService.GetApiSubRelInfoList(ctx,pageInfo)
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

// GetApiSubRelPublic 不需要鉴权的做市子账户接口
// @Tags ApiSubRel
// @Summary 不需要鉴权的做市子账户接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /apiSubRel/getApiSubRelPublic [get]
func (apiSubRelApi *ApiSubRelApi) GetApiSubRelPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    apiSubRelService.GetApiSubRelPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的做市子账户接口信息",
    }, "获取成功", c)
}
