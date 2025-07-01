package lbank

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Name struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// getSymbolsForLinkedLbankAccount 获取可用于添加Lbank账号的symbol列表
// @Tags      Lbank
// @Summary   获取可用于添加Lbank账号的symbol列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]Name,msg=string}  "获取成功"
// @Router    /db/getSymbolsForLinkedLbankAccount [get]
func (s *DbPageApi) GetSymbolsForLinkedLbankAccount(c *gin.Context) {

	var names []Name

	db := global.GetGlobalDBByDBName("daemon")
	if db == nil {
		global.GVA_LOG.Error("daemonDB连接不存在")
		response.FailWithMessage("daemonDB连接失败", c)
		return
	}

	result := db.Raw(`
		SELECT
			ins.id, ins.name 
		FROM instrument ins
		LEFT JOIN api_sub_rel sub
			ON ins.id = sub.inst_id AND sub.sub_type = '1'
		WHERE
			ins.exch_id = 3 AND sub.id IS NULL`,
	).Scan(&names)
	err := result.Error
	rowsAffected := result.RowsAffected

	if err != nil {
		global.GVA_LOG.Error("获取symbol列表失败", zap.Error(err))
		response.FailWithMessage("获取symbol列表失败: "+err.Error(), c)
		return
	}

	global.GVA_LOG.Info(fmt.Sprintf("获取symbol列表行数: %d", rowsAffected))
	response.OkWithData(names, c)
}

// GetLbankAccounts 获取Lbank账号列表
// @Tags      Lbank
// @Summary   获取Lbank账号列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]Name,msg=string}  "获取成功"
// @Router    /db/getLbankAccounts [get]
func (s *DbPageApi) GetLbankAccounts(c *gin.Context) {
	var names []Name

	db := global.GetGlobalDBByDBName("daemon")
	if db == nil {
		global.GVA_LOG.Error("daemonDB连接不存在")
		response.FailWithMessage("daemonDB连接失败", c)
		return
	}

	result := db.Raw(`
		SELECT
			act.id, act.name
		FROM api_account act
		INNER JOIN exchange_info ex
			ON act.exch_id = ex.id AND ex.exch_name = 'lbank'`,
	).Scan(&names)
	err := result.Error
	rowsAffected := result.RowsAffected

	if err != nil {
		global.GVA_LOG.Error("获取Lbank账号列表失败", zap.Error(err))
		response.FailWithMessage("获取Lbank账号列表失败: "+err.Error(), c)
		return
	}

	global.GVA_LOG.Info(fmt.Sprintf("获取Lbank账号列表行数: %d", rowsAffected))
	response.OkWithData(names, c)
}

// SubmitLinkedLbankSubAccounts 关联母账号-子账号提交接口
// @Tags      Lbank
// @Summary   关联母账号-子账号提交
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data body SubmitLinkedLbankSubAccountsRequest true "提交数据"
// @Success   200   {object}  response.Response{msg=string}  "提交成功"
// @Router    /db/submitLinkedLbankSubAccounts [post]
type SubmitLinkedLbankSubAccountsRequest struct {
	ParentID      uint   `json:"parentId"`
	SubAccountIDs []uint `json:"subAccountIds"`
}

func (s *DbPageApi) SubmitLinkedLbankSubAccounts(c *gin.Context) {
	var req SubmitLinkedLbankSubAccountsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	// 打印参数内容
	global.GVA_LOG.Info("SubmitLinkedLbankSubAccounts 参数", zap.Any("req", req))
	//print log: {"req": {"parentId":5,"subAccountIds":[1]}}

	// 这里可以处理实际的保存逻辑
	response.OkWithMessage("提交成功", c)
}
