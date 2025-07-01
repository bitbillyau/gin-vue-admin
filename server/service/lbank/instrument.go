
package lbank

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
)

type InstrumentService struct {}
// CreateInstrument 创建交易对管理记录
// Author [yourname](https://github.com/yourname)
func (instrumentService *InstrumentService) CreateInstrument(ctx context.Context, instrument *lbank.Instrument) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Create(instrument).Error
	return err
}

// DeleteInstrument 删除交易对管理记录
// Author [yourname](https://github.com/yourname)
func (instrumentService *InstrumentService)DeleteInstrument(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Delete(&lbank.Instrument{},"id = ?",id).Error
	return err
}

// DeleteInstrumentByIds 批量删除交易对管理记录
// Author [yourname](https://github.com/yourname)
func (instrumentService *InstrumentService)DeleteInstrumentByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Delete(&[]lbank.Instrument{},"id in ?",ids).Error
	return err
}

// UpdateInstrument 更新交易对管理记录
// Author [yourname](https://github.com/yourname)
func (instrumentService *InstrumentService)UpdateInstrument(ctx context.Context, instrument lbank.Instrument) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Model(&lbank.Instrument{}).Where("id = ?",instrument.Id).Updates(&instrument).Error
	return err
}

// GetInstrument 根据id获取交易对管理记录
// Author [yourname](https://github.com/yourname)
func (instrumentService *InstrumentService)GetInstrument(ctx context.Context, id string) (instrument lbank.Instrument, err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Where("id = ?", id).First(&instrument).Error
	return
}
// GetInstrumentInfoList 分页获取交易对管理记录
// Author [yourname](https://github.com/yourname)
func (instrumentService *InstrumentService)GetInstrumentInfoList(ctx context.Context, info lbankReq.InstrumentSearch) (list []lbank.Instrument, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("daemon").Model(&lbank.Instrument{})
    var instruments []lbank.Instrument
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&instruments).Error
	return  instruments, total, err
}
func (instrumentService *InstrumentService)GetInstrumentPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
