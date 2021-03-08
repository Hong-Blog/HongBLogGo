package bizType

import "github.com/guregu/null"

type AddBizTypeRequest struct {
	Pid         null.Int    `json:"pid"`
	Name        null.String `json:"name"`        // 文章类型名
	Description null.String `json:"description"` // 类型介绍
	Sort        null.Int    `json:"sort"`        // 排序
	Icon        null.String `json:"icon"`        // 图标
	Available   null.Int    `json:"available"`   // 是否可用
}
