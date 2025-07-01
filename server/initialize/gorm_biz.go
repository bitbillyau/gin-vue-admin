package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate()
	if err != nil {
		return err
	}
	daemonDb := global.GetGlobalDBByDBName("daemon")
	daemonDb.AutoMigrate(lbank.ApiAccount{},lbank.Instrument{},lbank.ApiSubRel{}, lbank.ServerApiRel{})
	return nil
}
