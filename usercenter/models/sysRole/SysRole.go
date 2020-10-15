package sysRole

import (
	"github.com/guregu/null"
	"log"
	"usercenter/db"
	"usercenter/models"
)

type SysRole struct {
	Id          int64       `json:"id"`
	Name        null.String `json:"name"` // 角色名
	Description null.String `json:"description"`
	Available   null.Int    `json:"available"`
	CreateTime  null.Time   `json:"create_time" db:"create_time"` // 添加时间
	UpdateTime  null.Time   `json:"update_time" db:"update_time"` // 更新时间
}

type GetAllRoleRequest struct {
	models.PagedRequest
	KeyWord string `json:"key_word"`
}

func GetAllRole(request GetAllRoleRequest) (list []SysRole, count int) {
	dataSql := `
select id, name, description, available, create_time, update_time
from sys_role
`
	countSql := "select count(1) from sys_role "
	var params = make([]interface{}, 0)
	var filter string
	if len(request.KeyWord) > 0 {
		filter = " where name like ? "
		params = append(params, "%"+request.KeyWord+"%")
	}
	dataSql += filter + " limit ?,? "
	countSql += filter

	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count sys_role err: ", err.Error())
	}

	offset, limit := request.GetLimit()
	params = append(params, offset)
	params = append(params, limit)
	err = db.Db.Select(&list, dataSql, params...)
	if err != nil {
		log.Panicln("select sys_role err: ", err.Error())
	}

	return
}
