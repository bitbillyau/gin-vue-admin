
package lbank

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/lbank"
    lbankReq "github.com/flipped-aurora/gin-vue-admin/server/model/lbank/request"
)

type ApiAccountService struct {}
// CreateApiAccount 创建做市账户管理记录
// Author [yourname](https://github.com/yourname)
func (apiAccountService *ApiAccountService) CreateApiAccount(ctx context.Context, apiAccount *lbank.ApiAccount) (err error) {
	err = global.GVA_DB.Create(apiAccount).Error
	return err
}

// DeleteApiAccount 删除做市账户管理记录
// Author [yourname](https://github.com/yourname)
func (apiAccountService *ApiAccountService)DeleteApiAccount(ctx context.Context, id string) (err error) {
	err = global.GVA_DB.Delete(&lbank.ApiAccount{},"id = ?",id).Error
	return err
}

// DeleteApiAccountByIds 批量删除做市账户管理记录
// Author [yourname](https://github.com/yourname)
func (apiAccountService *ApiAccountService)DeleteApiAccountByIds(ctx context.Context, ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]lbank.ApiAccount{},"id in ?",ids).Error
	return err
}

// UpdateApiAccount 更新做市账户管理记录
// Author [yourname](https://github.com/yourname)
func (apiAccountService *ApiAccountService)UpdateApiAccount(ctx context.Context, apiAccount lbank.ApiAccount) (err error) {
	err = global.GVA_DB.Model(&lbank.ApiAccount{}).Where("id = ?",apiAccount.Id).Updates(&apiAccount).Error
	return err
}

// GetApiAccount 根据id获取做市账户管理记录
// Author [yourname](https://github.com/yourname)
func (apiAccountService *ApiAccountService)GetApiAccount(ctx context.Context, id string) (apiAccount lbank.ApiAccount, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&apiAccount).Error
	return
}
// GetApiAccountInfoList 分页获取做市账户管理记录
// Author [yourname](https://github.com/yourname)
func (apiAccountService *ApiAccountService)GetApiAccountInfoList(ctx context.Context, info lbankReq.ApiAccountSearch) (list []lbank.ApiAccount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&lbank.ApiAccount{})
    var apiAccounts []lbank.ApiAccount
    // 如果有条件搜索 下方会自动创建搜索语句
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&apiAccounts).Error
	return  apiAccounts, total, err
}
func (apiAccountService *ApiAccountService)GetApiAccountPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
