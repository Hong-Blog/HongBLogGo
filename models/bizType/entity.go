package bizType

import (
	"github.com/guregu/null"
)

type BizType struct {
	Id          int         `json:"id"`
	Pid         null.Int    `json:"pid"`
	Name        null.String `json:"name"`                         // 文章类型名
	Description null.String `json:"description"`                  // 类型介绍
	Sort        null.Int    `json:"sort"`                         // 排序
	Icon        null.String `json:"icon"`                         // 图标
	Available   null.Int    `json:"available"`                    // 是否可用
	CreateTime  null.Time   `json:"create_time" db:"create_time"` // 添加时间
	UpdateTime  null.Time   `json:"update_time" db:"update_time"` // 更新时间
}
