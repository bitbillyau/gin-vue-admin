package lbank

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DbPageApi
	ApiAccountApi
	InstrumentApi
	ApiSubRelApi
	ServerApiRelApi
}
type DbPageApi struct{}

var (
	apiAccountService   = service.ServiceGroupApp.LbankServiceGroup.ApiAccountService
	instrumentService   = service.ServiceGroupApp.LbankServiceGroup.InstrumentService
	apiSubRelService    = service.ServiceGroupApp.LbankServiceGroup.ApiSubRelService
	serverApiRelService = service.ServiceGroupApp.LbankServiceGroup.ServerApiRelService
)
