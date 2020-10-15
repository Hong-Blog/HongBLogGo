package sysUser

import (
	"github.com/guregu/null"
	"log"
	"usercenter/db"
	"usercenter/models"
)

type SysUser struct {
	Id            int         `json:"id"`
	Username      null.String `json:"username"`
	Password      null.String `json:"password"`                             // 登录密码
	Nickname      null.String `json:"nickname"`                             // 昵称
	Mobile        null.String `json:"mobile"`                               // 手机号
	Email         null.String `json:"email"`                                // 邮箱地址
	Qq            null.String `json:"qq"`                                   // QQ
	Birthday      null.Time   `json:"birthday"`                             // 生日
	Gender        null.Int    `json:"gender"`                               // 性别
	Avatar        null.String `json:"avatar"`                               // 头像地址
	UserType      null.String `json:"user_type" db:"user_type"`             // 超级管理员、管理员、普通用户
	Company       null.String `json:"company"`                              // 公司
	Blog          null.String `json:"blog"`                                 // 个人博客地址
	Location      null.String `json:"location"`                             // 地址
	Source        null.String `json:"source"`                               // 用户来源
	Uuid          null.String `json:"uuid"`                                 // 用户唯一表示(第三方网站)
	Privacy       null.Int    `json:"privacy"`                              // 隐私（1：公开，0：不公开）
	Notification  null.Int    `json:"notification"`                         // 通知：(1：通知显示消息详情，2：通知不显示详情)
	Score         null.Int    `json:"score"`                                // 金币值
	Experience    null.Int    `json:"experience"`                           // 经验值
	RegIp         null.String `json:"reg_ip" db:"reg_ip"`                   // 注册IP
	LastLoginIp   null.String `json:"last_login_ip" db:"last_login_ip"`     // 最近登录IP
	LastLoginTime null.Time   `json:"last_login_time" db:"last_login_time"` // 最近登录时间
	LoginCount    null.Int    `json:"login_count" db:"login_count"`         // 登录次数
	Remark        null.String `json:"remark"`                               // 用户备注
	Status        null.Int    `json:"status"`                               // 用户状态
	CreateTime    null.Time   `json:"create_time" db:"create_time"`         // 注册时间
	UpdateTime    null.Time   `json:"update_time" db:"update_time"`         // 更新时间
}

type GetAllUserRequest struct {
	models.PagedRequest
	KeyWord string `json:"key_word"`
}

func GetAllUser(request GetAllUserRequest) (list []SysUser, count int) {
	strSql := `
select id,
       username,
       password,
       nickname,
       mobile,
       email,
       qq,
       birthday,
       gender,
       avatar,
       user_type,
       company,
       blog,
       location,
       source,
       uuid,
       privacy,
       notification,
       score,
       experience,
       reg_ip,
       last_login_ip,
       last_login_time,
       login_count,
       remark,
       status,
       create_time,
       update_time
from sys_user
`
	var params = make([]interface{}, 0)
	var filter string
	if len(request.KeyWord) != 0 {
		filter = " where username like ? "
		params = append(params, "%"+request.KeyWord+"%")
	}
	strSql += filter + " limit ?, ?;"
	countSql := "select count(1) from sys_user " + filter
	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count sys_user err: ", err.Error())
	}

	offset, limit := request.GetLimit()
	params = append(params, offset)
	params = append(params, limit)
	err = db.Db.Select(&list, strSql, params...)
	if err != nil {
		log.Panicln("select sys_user err: ", err.Error())
	}

	return
}
