package bizTags

import "github.com/guregu/null"

type AddBizTagRequest struct {
	Name        string      `json:"name"`        // 书签名
	Description null.String `json:"description"` // 描述
}
