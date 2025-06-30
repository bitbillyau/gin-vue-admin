package lbank

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DbPageApi
	ApiAccountApi
}
type DbPageApi struct{}

var apiAccountService = service.ServiceGroupApp.LbankServiceGroup.ApiAccountService
