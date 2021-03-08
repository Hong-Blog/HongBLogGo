package bizTags

import (
	"errors"
	"hong-blog/db"
	"hong-blog/models"
	"log"
)

func (t *BizTags) ExistByName() bool {
	dataSql := `
select ifnull((select 1
               from biz_tags
               where name = ?
               limit 1), 0) a;
`
	res := 0
	if err := db.Db.QueryRow(dataSql, t.Name).Scan(&res); err != nil {
		log.Panicln("BizTags ExistByName err: ", err.Error())
	}
	return res > 0
}

func (t *BizTags) Update() error {
	updateSql := `
update biz_tags
set name        = :name,
    description = :description,
    update_time = now()
where id = :id
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

func (t *BizTags) Delete() error {
	articleCount, err := countArticleByTagId(t.Id)
	if err != nil {
		return err
	}
	if articleCount > 0 {
		return errors.New("有关联博客，禁止删除")
	}

	deleteSql := `
delete
from biz_tags
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

func GetTagList(req models.PagedRequest) (list []BizTags, count int) {
	sql := `
select id, name, description, create_time, update_time
from biz_tags
`
	var params = make([]interface{}, 0)

	var filter string
	sql += filter + " order by id desc limit ?, ?;"
	countSql := "select count(1) from biz_tags " + filter

	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count biz_tags err: ", err.Error())
	}

	offset, limit := req.GetLimit()
	params = append(params, offset)
	params = append(params, limit)

	if err := db.Db.Select(&list, sql, params...); err != nil {
		log.Panicln("select biz_tags err: " + err.Error())
	}
	return
}

func AddBizTag(request AddBizTagRequest) (err error) {
	bizTags := BizTags{Name: request.Name}
	if bizTags.ExistByName() {
		err = errors.New("标签已存在")
		return
	}

	insertSql := `
insert into biz_tags
    (name, description, create_time)
    value
    (:name, :description, now())
`
	if _, err := db.Db.NamedExec(insertSql, request); err != nil {
		return err
	}
	return nil
}

func GetById(id int) (bizTag BizTags, err error) {
	dataSql := `
select id, name, description, create_time, update_time
from biz_tags
where id = ?
`
	err = db.Db.Get(&bizTag, dataSql, id)
	return
}

func GetByArticleId(articleId int) (list []BizTags, err error) {
	dataSql := `
select bt.id, name, description, bt.create_time, bt.update_time
from biz_tags bt
         inner join biz_article_tags bat on bt.id = bat.tag_id
where bat.article_id = ?
`
	err = db.Db.Select(&list, dataSql, articleId)
	return
}

func countArticleByTagId(tagId int) (count int, err error) {
	dataSql := `
select count(1)
from biz_article_tags
where tag_id = ?
`
	if err := db.Db.QueryRow(dataSql, tagId).Scan(&count); err != nil {
		return 0, err
	}
	return count, nil
}
