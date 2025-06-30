package lbank

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
)

type RouterGroup struct {
	DbRouter
	ApiAccountRouter
}

var (
	dbPageApi     = api.ApiGroupApp.LbankApiGroup.DbPageApi
	apiAccountApi = api.ApiGroupApp.LbankApiGroup.ApiAccountApi
)
