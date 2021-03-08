package bizTags

import (
	"github.com/guregu/null"
)

type BizTags struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`                         // 书签名
	Description null.String `json:"description"`                  // 描述
	CreateTime  null.Time   `json:"create_time" db:"create_time"` // 添加时间
	UpdateTime  null.Time   `json:"update_time" db:"update_time"` // 更新时间
}
