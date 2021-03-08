package bizArticle

import (
	"github.com/guregu/null"
)

type BizArticle struct {
	Id          int         `json:"id"`
	Title       null.String `json:"title"`                        // 文章标题
	UserId      int         `json:"user_id" db:"user_id"`         // 用户ID
	CoverImage  null.String `json:"cover_image" db:"cover_image"` // 文章封面图片
	QrcodePath  null.String `json:"qrcode_path" db:"qrcode_path"` // 文章专属二维码地址
	IsMarkdown  null.Int    `json:"is_markdown" db:"is_markdown"`
	Content     null.String `json:"content"`                      // 文章内容
	ContentMd   null.String `json:"content_md" db:"content_md"`   // markdown版的文章内容
	Top         null.Int    `json:"top"`                          // 是否置顶
	TypeId      int         `json:"type_id" db:"type_id"`         // 类型
	Status      null.Int    `json:"status"`                       // 状态
	Recommended null.Int    `json:"recommended"`                  // 是否推荐
	Original    null.Int    `json:"original"`                     // 是否原创
	Description null.String `json:"description"`                  // 文章简介，最多200字
	Keywords    null.String `json:"keywords"`                     // 文章关键字，优化搜索
	Comment     null.Int    `json:"comment"`                      // 是否开启评论
	CreateTime  null.Time   `json:"create_time" db:"create_time"` // 添加时间
	UpdateTime  null.Time   `json:"update_time" db:"update_time"` // 更新时间
}
