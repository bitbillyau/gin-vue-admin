package lbank

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DbPageApi
	ApiAccountApi
	InstrumentApi
}
type DbPageApi struct{}

var (
	apiAccountService = service.ServiceGroupApp.LbankServiceGroup.ApiAccountService
	instrumentService = service.ServiceGroupApp.LbankServiceGroup.InstrumentService
)
