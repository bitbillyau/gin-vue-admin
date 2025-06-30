
// 自动生成模板ServerApiRel
package lbank
import (
)

// 做市API操作控制管理 结构体  ServerApiRel
type ServerApiRel struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:10;"`  //控制序号
  ServerId  *int `json:"serverId" form:"serverId" gorm:"column:server_id;size:10;"`  //服务器序号
  ApiId  *int `json:"apiId" form:"apiId" gorm:"column:api_id;size:10;"`  //账户序号
  InstId  *int `json:"instId" form:"instId" gorm:"column:inst_id;size:10;"`  //交易对序号
  Status  *int `json:"status" form:"status" gorm:"column:status;size:10;"`  //是否有效
}


// TableName 做市API操作控制管理 ServerApiRel自定义表名 server_api_rel
func (ServerApiRel) TableName() string {
    return "server_api_rel"
}





