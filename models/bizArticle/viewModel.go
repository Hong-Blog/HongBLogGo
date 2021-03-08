package bizArticle

import (
	"github.com/guregu/null"
	"hong-blog/models"
	"hong-blog/models/bizTags"
	"hong-blog/models/bizType"
)

type GetArticleListRequest struct {
	models.PagedRequest
	Keyword string `json:"keyword"`
	TypeId  int    `json:"type_id"`
}

type ArticleModel struct {
	Id           int         `json:"id"`
	Title        null.String `json:"title"`                        // 文章标题
	UserId       int         `json:"user_id" db:"user_id"`         // 用户ID
	CoverImage   null.String `json:"cover_image" db:"cover_image"` // 文章封面图片
	QrcodePath   null.String `json:"qrcode_path" db:"qrcode_path"` // 文章专属二维码地址
	IsMarkdown   null.Int    `json:"is_markdown" db:"is_markdown"`
	Top          null.Int    `json:"top"`                          // 是否置顶
	TypeId       int8        `json:"type_id" db:"type_id"`         // 类型
	TypeName     null.String `json:"type_name" db:"type_name"`     // 类型名称
	Status       null.Int    `json:"status"`                       // 状态
	StatusString null.String `json:"status_string"`                // 状态
	Recommended  null.Int    `json:"recommended"`                  // 是否推荐
	Original     null.Int    `json:"original"`                     // 是否原创
	Comment      null.Int    `json:"comment"`                      // 是否开启评论
	CreateTime   null.Time   `json:"create_time" db:"create_time"` // 添加时间
	UpdateTime   null.Time   `json:"update_time" db:"update_time"` // 更新时间
}

type ArticleDetail struct {
	BizArticle
	BizType bizType.BizType
	BizTags []bizTags.BizTags
}

type AddArticleModel struct {
	Title       null.String `json:"title"`                        // 文章标题
	UserId      int         `db:"user_id"`                        // 用户ID
	CoverImage  null.String `json:"cover_image" db:"cover_image"` // 文章封面图片
	IsMarkdown  null.Int    `json:"is_markdown" db:"is_markdown"`
	Content     null.String `json:"content"`                    // 文章内容
	ContentMd   null.String `json:"content_md" db:"content_md"` // markdown版的文章内容
	TypeId      int         `json:"type_id" db:"type_id"`       // 类型
	Status      null.Int    `json:"status"`                     // 状态
	Original    null.Int    `json:"original"`                   // 是否原创
	Description null.String `json:"description"`                // 文章简介，最多200字
	Comment     null.Int    `json:"comment"`                    // 是否开启评论
	TagIds      []int       `json:"tag_ids"`                    // 标签
}

type EditArticleModel struct {
	Id int `json:"id"`
	AddArticleModel
}
