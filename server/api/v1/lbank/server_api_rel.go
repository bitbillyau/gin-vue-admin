package lbank

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ServerApiRelApi struct {}



// CreateServerApiRel 创建做市API操作控制管理
// @Tags ServerApiRel
// @Summary 创建做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ServerApiRel true "创建做市API操作控制管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /serverApiRel/createServerApiRel [post]
func (serverApiRelApi *ServerApiRelApi) CreateServerApiRel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var serverApiRel lbank.ServerApiRel
	err := c.ShouldBindJSON(&serverApiRel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serverApiRelService.CreateServerApiRel(ctx,&serverApiRel)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteServerApiRel 删除做市API操作控制管理
// @Tags ServerApiRel
// @Summary 删除做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ServerApiRel true "删除做市API操作控制管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /serverApiRel/deleteServerApiRel [delete]
func (serverApiRelApi *ServerApiRelApi) DeleteServerApiRel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := serverApiRelService.DeleteServerApiRel(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteServerApiRelByIds 批量删除做市API操作控制管理
// @Tags ServerApiRel
// @Summary 批量删除做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /serverApiRel/deleteServerApiRelByIds [delete]
func (serverApiRelApi *ServerApiRelApi) DeleteServerApiRelByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := serverApiRelService.DeleteServerApiRelByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateServerApiRel 更新做市API操作控制管理
// @Tags ServerApiRel
// @Summary 更新做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.ServerApiRel true "更新做市API操作控制管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /serverApiRel/updateServerApiRel [put]
func (serverApiRelApi *ServerApiRelApi) UpdateServerApiRel(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var serverApiRel lbank.ServerApiRel
	err := c.ShouldBindJSON(&serverApiRel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = serverApiRelService.UpdateServerApiRel(ctx,serverApiRel)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindServerApiRel 用id查询做市API操作控制管理
// @Tags ServerApiRel
// @Summary 用id查询做市API操作控制管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询做市API操作控制管理"
// @Success 200 {object} response.Response{data=lbank.ServerApiRel,msg=string} "查询成功"
// @Router /serverApiRel/findServerApiRel [get]
func (serverApiRelApi *ServerApiRelApi) FindServerApiRel(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reserverApiRel, err := serverApiRelService.GetServerApiRel(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reserverApiRel, c)
}
// GetServerApiRelList 分页获取做市API操作控制管理列表
// @Tags ServerApiRel
// @Summary 分页获取做市API操作控制管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.ServerApiRelSearch true "分页获取做市API操作控制管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /serverApiRel/getServerApiRelList [get]
func (serverApiRelApi *ServerApiRelApi) GetServerApiRelList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo lbankReq.ServerApiRelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := serverApiRelService.GetServerApiRelInfoList(ctx,pageInfo)
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

// GetServerApiRelPublic 不需要鉴权的做市API操作控制管理接口
// @Tags ServerApiRel
// @Summary 不需要鉴权的做市API操作控制管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /serverApiRel/getServerApiRelPublic [get]
func (serverApiRelApi *ServerApiRelApi) GetServerApiRelPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    serverApiRelService.GetServerApiRelPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的做市API操作控制管理接口信息",
    }, "获取成功", c)
}
