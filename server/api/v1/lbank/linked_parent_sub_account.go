package lbank

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type Name struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// GetSymbolsForAddLbankAccount 获取可用于添加Lbank账号的symbol列表
// @Tags      Lbank
// @Summary   获取可用于添加Lbank账号的symbol列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=[]Name,msg=string}  "获取成功"
// @Router    /db/getSymbolsForAddLbankAccount [get]
func (s *DbPageApi) GetSymbolsForAddLbankAccount(c *gin.Context) {
	// 这里先返回模拟数据，后续可替换为数据库查询
	symbols := []Name{
		{ID: 1, Name: "BTC/USDT"},
		{ID: 2, Name: "ETH/USDT"},
		{ID: 3, Name: "LTC/USDT"},
	}
	response.OkWithData(symbols, c)
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
	accounts := []Name{
		{ID: 1, Name: "母账号1"},
		{ID: 2, Name: "母账号2"},
		{ID: 3, Name: "母账号3"},
	}
	response.OkWithData(accounts, c)
}
