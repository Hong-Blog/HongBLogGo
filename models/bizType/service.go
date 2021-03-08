package bizType

import (
	"errors"
	"hong-blog/db"
	"hong-blog/models"
	"log"
)

func (t *BizType) ExistByName() bool {
	dataSql := `
select ifnull((select 1
               from biz_type
               where name = ?
               limit 1), 0) a;
`
	res := 0
	if err := db.Db.QueryRow(dataSql, t.Name).Scan(&res); err != nil {
		log.Panicln("BizType ExistByName err: ", err.Error())
	}
	return res > 0
}

func GetTypeList(req models.PagedRequest) (list []BizType, count int) {
	sql := `
select id,
       pid,
       name,
       description,
       sort,
       icon,
       available,
       create_time,
       update_time
from biz_type
`

	var params = make([]interface{}, 0)

	var filter string
	sql += filter + " order by id desc limit ?, ?;"
	countSql := "select count(1) from biz_type " + filter

	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count biz_type err: ", err.Error())
	}

	offset, limit := req.GetLimit()
	params = append(params, offset)
	params = append(params, limit)

	if err := db.Db.Select(&list, sql, params...); err != nil {
		log.Panicln("select biz_type err: " + err.Error())
	}
	return
}

func AddBizType(request AddBizTypeRequest) (err error) {
	bizType := BizType{Name: request.Name}
	if bizType.ExistByName() {
		err = errors.New("分类已存在")
		return
	}

	insertSql := `
insert into biz_type
    (pid, name, description, sort, icon, available, create_time)
    value
    (:pid, :name, :description, :sort, :icon, :available, now())
`
	if _, err := db.Db.NamedExec(insertSql, request); err != nil {
		return err
	}
	return nil
}

func GetById(id int) (bizType BizType, err error) {
	dataSql := `
select id,
       pid,
       name,
       description,
       sort,
       icon,
       available,
       create_time,
       update_time
from biz_type
where id = ?;
`
	err = db.Db.Get(&bizType, dataSql, id)
	return
}

func (t *BizType) Update() error {
	updateSql := `
update biz_type
set name        = :name,
    pid         = :pid,
    description = :description,
    sort        = :sort,
    icon        = :icon,
    available   = :available,
    update_time = now()
where id = :id;
`
	result, err := db.Db.NamedExec(updateSql, t)
	if err != nil {
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func (t *BizType) Delete() error {
	articleCount, err := countArticleByTypeId(t.Id)
	if err != nil {
		return err
	}
	if articleCount > 0 {
		return errors.New("有关联博客，禁止删除")
	}

	deleteSql := `
delete
from biz_type
where id = ?
`
	result, err := db.Db.Exec(deleteSql, t.Id)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("删除失败")
	}
	return nil
}

func countArticleByTypeId(typeId int) (count int, err error) {
	dataSql := `
select count(1)
from biz_article
where type_id = ?
`
	if err := db.Db.QueryRow(dataSql, typeId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
