
package lbank

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
)

type ApiSubRelService struct {}
// CreateApiSubRel 创建做市子账户记录
// Author [yourname](https://github.com/yourname)
func (apiSubRelService *ApiSubRelService) CreateApiSubRel(ctx context.Context, apiSubRel *lbank.ApiSubRel) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Create(apiSubRel).Error
	return err
}

// DeleteApiSubRel 删除做市子账户记录
// Author [yourname](https://github.com/yourname)
func (apiSubRelService *ApiSubRelService)DeleteApiSubRel(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Delete(&lbank.ApiSubRel{},"id = ?",id).Error
	return err
}

// DeleteApiSubRelByIds 批量删除做市子账户记录
// Author [yourname](https://github.com/yourname)
func (apiSubRelService *ApiSubRelService)DeleteApiSubRelByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Delete(&[]lbank.ApiSubRel{},"id in ?",ids).Error
	return err
}

// UpdateApiSubRel 更新做市子账户记录
// Author [yourname](https://github.com/yourname)
func (apiSubRelService *ApiSubRelService)UpdateApiSubRel(ctx context.Context, apiSubRel lbank.ApiSubRel) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Model(&lbank.ApiSubRel{}).Where("id = ?",apiSubRel.Id).Updates(&apiSubRel).Error
	return err
}

// GetApiSubRel 根据id获取做市子账户记录
// Author [yourname](https://github.com/yourname)
func (apiSubRelService *ApiSubRelService)GetApiSubRel(ctx context.Context, id string) (apiSubRel lbank.ApiSubRel, err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Where("id = ?", id).First(&apiSubRel).Error
	return
}
// GetApiSubRelInfoList 分页获取做市子账户记录
// Author [yourname](https://github.com/yourname)
func (apiSubRelService *ApiSubRelService)GetApiSubRelInfoList(ctx context.Context, info lbankReq.ApiSubRelSearch) (list []lbank.ApiSubRel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("daemon").Model(&lbank.ApiSubRel{})
    var apiSubRels []lbank.ApiSubRel
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&apiSubRels).Error
	return  apiSubRels, total, err
}
func (apiSubRelService *ApiSubRelService)GetApiSubRelPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
