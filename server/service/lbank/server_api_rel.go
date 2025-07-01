
package lbank

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
)

type ServerApiRelService struct {}
// CreateServerApiRel 创建做市API操作控制管理记录
// Author [yourname](https://github.com/yourname)
func (serverApiRelService *ServerApiRelService) CreateServerApiRel(ctx context.Context, serverApiRel *lbank.ServerApiRel) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Create(serverApiRel).Error
	return err
}

// DeleteServerApiRel 删除做市API操作控制管理记录
// Author [yourname](https://github.com/yourname)
func (serverApiRelService *ServerApiRelService)DeleteServerApiRel(ctx context.Context, id string) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Delete(&lbank.ServerApiRel{},"id = ?",id).Error
	return err
}

// DeleteServerApiRelByIds 批量删除做市API操作控制管理记录
// Author [yourname](https://github.com/yourname)
func (serverApiRelService *ServerApiRelService)DeleteServerApiRelByIds(ctx context.Context, ids []string) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Delete(&[]lbank.ServerApiRel{},"id in ?",ids).Error
	return err
}

// UpdateServerApiRel 更新做市API操作控制管理记录
// Author [yourname](https://github.com/yourname)
func (serverApiRelService *ServerApiRelService)UpdateServerApiRel(ctx context.Context, serverApiRel lbank.ServerApiRel) (err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Model(&lbank.ServerApiRel{}).Where("id = ?",serverApiRel.Id).Updates(&serverApiRel).Error
	return err
}

// GetServerApiRel 根据id获取做市API操作控制管理记录
// Author [yourname](https://github.com/yourname)
func (serverApiRelService *ServerApiRelService)GetServerApiRel(ctx context.Context, id string) (serverApiRel lbank.ServerApiRel, err error) {
	err = global.MustGetGlobalDBByDBName("daemon").Where("id = ?", id).First(&serverApiRel).Error
	return
}
// GetServerApiRelInfoList 分页获取做市API操作控制管理记录
// Author [yourname](https://github.com/yourname)
func (serverApiRelService *ServerApiRelService)GetServerApiRelInfoList(ctx context.Context, info lbankReq.ServerApiRelSearch) (list []lbank.ServerApiRel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("daemon").Model(&lbank.ServerApiRel{})
    var serverApiRels []lbank.ServerApiRel
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&serverApiRels).Error
	return  serverApiRels, total, err
}
func (serverApiRelService *ServerApiRelService)GetServerApiRelPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
