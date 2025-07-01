
// 自动生成模板ApiAccount
package lbank
import (
)

// 做市账户管理 结构体  ApiAccount
type ApiAccount struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;"`  //序号
  Name  *string `json:"name" form:"name" gorm:"column:name;size:64;"`  //邮箱
  ApiKey  *string `json:"apiKey" form:"apiKey" gorm:"column:api_key;size:255;"`  //apiKey
  ApiSecret  *string `json:"apiSecret" form:"apiSecret" gorm:"column:api_secret;"`  //apiSecret
  Uid  *string `json:"uid" form:"uid" gorm:"column:uid;size:64;"`  //uid
  Passphrase  *string `json:"passphrase" form:"passphrase" gorm:"comment:okx等交易所需要;column:passphrase;size:64;"`  //passphrase
  ExchId  *int `json:"exchId" form:"exchId" gorm:"comment:1-lbk,2-okx,3-lbank;column:exch_id;size:10;"`  //交易所ID
  Status  *int `json:"status" form:"status" gorm:"default:1;column:status;size:10;"`  //是否生效
}


// TableName 做市账户管理 ApiAccount自定义表名 api_account
func (ApiAccount) TableName() string {
    return "api_account"
}





