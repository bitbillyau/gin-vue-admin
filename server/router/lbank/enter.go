package lbank

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	DbRouter
	ApiAccountRouter
	InstrumentRouter
	ApiSubRelRouter
	ServerApiRelRouter
}

var (
	dbPageApi       = api.ApiGroupApp.LbankApiGroup.DbPageApi
	apiAccountApi   = api.ApiGroupApp.LbankApiGroup.ApiAccountApi
	instrumentApi   = api.ApiGroupApp.LbankApiGroup.InstrumentApi
	apiSubRelApi    = api.ApiGroupApp.LbankApiGroup.ApiSubRelApi
	serverApiRelApi = api.ApiGroupApp.LbankApiGroup.ServerApiRelApi
)
