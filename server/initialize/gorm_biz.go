package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(lbank.ApiAccount{}, lbank.Instrument{})
	if err != nil {
		return err
	}
	return nil
}
