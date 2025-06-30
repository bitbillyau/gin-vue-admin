package lbank

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type InstrumentApi struct {}



// CreateInstrument 创建交易对管理
// @Tags Instrument
// @Summary 创建交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.Instrument true "创建交易对管理"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /instrument/createInstrument [post]
func (instrumentApi *InstrumentApi) CreateInstrument(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var instrument lbank.Instrument
	err := c.ShouldBindJSON(&instrument)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = instrumentService.CreateInstrument(ctx,&instrument)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteInstrument 删除交易对管理
// @Tags Instrument
// @Summary 删除交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.Instrument true "删除交易对管理"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /instrument/deleteInstrument [delete]
func (instrumentApi *InstrumentApi) DeleteInstrument(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	err := instrumentService.DeleteInstrument(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteInstrumentByIds 批量删除交易对管理
// @Tags Instrument
// @Summary 批量删除交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /instrument/deleteInstrumentByIds [delete]
func (instrumentApi *InstrumentApi) DeleteInstrumentByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ids := c.QueryArray("ids[]")
	err := instrumentService.DeleteInstrumentByIds(ctx,ids)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateInstrument 更新交易对管理
// @Tags Instrument
// @Summary 更新交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body lbank.Instrument true "更新交易对管理"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /instrument/updateInstrument [put]
func (instrumentApi *InstrumentApi) UpdateInstrument(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var instrument lbank.Instrument
	err := c.ShouldBindJSON(&instrument)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = instrumentService.UpdateInstrument(ctx,instrument)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindInstrument 用id查询交易对管理
// @Tags Instrument
// @Summary 用id查询交易对管理
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param id query int true "用id查询交易对管理"
// @Success 200 {object} response.Response{data=lbank.Instrument,msg=string} "查询成功"
// @Router /instrument/findInstrument [get]
func (instrumentApi *InstrumentApi) FindInstrument(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	id := c.Query("id")
	reinstrument, err := instrumentService.GetInstrument(ctx,id)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reinstrument, c)
}
// GetInstrumentList 分页获取交易对管理列表
// @Tags Instrument
// @Summary 分页获取交易对管理列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query lbankReq.InstrumentSearch true "分页获取交易对管理列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /instrument/getInstrumentList [get]
func (instrumentApi *InstrumentApi) GetInstrumentList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo lbankReq.InstrumentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := instrumentService.GetInstrumentInfoList(ctx,pageInfo)
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

// GetInstrumentPublic 不需要鉴权的交易对管理接口
// @Tags Instrument
// @Summary 不需要鉴权的交易对管理接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /instrument/getInstrumentPublic [get]
func (instrumentApi *InstrumentApi) GetInstrumentPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    instrumentService.GetInstrumentPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的交易对管理接口信息",
    }, "获取成功", c)
}
