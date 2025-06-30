
// 自动生成模板Instrument
package lbank
import (
)

// 交易对管理 结构体  Instrument
type Instrument struct {
  Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;size:10;"`  //序号
  ExchId  *int `json:"exchId" form:"exchId" gorm:"column:exch_id;size:19;"`  //交易所
  InstId  *int `json:"instId" form:"instId" gorm:"column:inst_id;size:19;"`  //交易对编号
  Name  *string `json:"name" form:"name" gorm:"column:name;size:21;"`  //交易对名称
  Status  *int `json:"status" form:"status" gorm:"default:1;column:status;size:10;"`  //是否有效
}


// TableName 交易对管理 Instrument自定义表名 instrument
func (Instrument) TableName() string {
    return "instrument"
}





