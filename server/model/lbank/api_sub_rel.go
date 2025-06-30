
// 自动生成模板ApiSubRel
package lbank
import (
)

// 做市子账户 结构体  ApiSubRel
type ApiSubRel struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:10;"`  //id字段
  ApiId  *int `json:"apiId" form:"apiId" gorm:"column:api_id;size:10;"`  //账户序号
  MarketuNid  *string `json:"marketuNid" form:"marketuNid" gorm:"column:marketu_nid;size:64;"`  //子账户编号
  InstId  *int `json:"instId" form:"instId" gorm:"column:inst_id;size:10;"`  //交易对序号
  ShowName  *string `json:"showName" form:"showName" gorm:"column:show_name;size:32;"`  //子账户备注
  Status  *int `json:"status" form:"status" gorm:"column:status;size:10;"`  //status字段
}


// TableName 做市子账户 ApiSubRel自定义表名 api_sub_rel
func (ApiSubRel) TableName() string {
    return "api_sub_rel"
}





