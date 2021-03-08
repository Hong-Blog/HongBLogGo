package bizArticle

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"hong-blog/db"
	"hong-blog/models/bizTags"
	"hong-blog/models/bizType"
	"log"
)

func (r *BizArticle) Delete() (err error) {
	sql := `
delete
from biz_article
where id = ?
`
	affected, affectedErr := db.Db.MustExec(sql, r.Id).RowsAffected()
	if affectedErr != nil {
		return affectedErr
	}
	if affected == 0 {
		return errors.New("删除失败")
	}
	return nil
}

func GetArticleList(req GetArticleListRequest) (list []ArticleModel, count int) {
	dataSql := `
select a.id,
       title,
       user_id,
       cover_image,
       qrcode_path,
       is_markdown,
       top,
       type_id,
       bt.name type_name,
       status,
       recommended,
       original,
       comment,
       a.create_time,
       a.update_time
from biz_article a
inner join biz_type bt on a.type_id = bt.id
where 1=1
`

	var params = make([]interface{}, 0)
	var filter string

	if len(req.Keyword) != 0 {
		filter += " and a.title like ? "
		params = append(params, "%"+req.Keyword+"%")
	}

	if req.TypeId > 0 {
		filter += " and a.type_id = ? "
		params = append(params, req.TypeId)
	}

	dataSql += filter + " order by a.id desc limit ?, ?;"
	countSql := `
select count(1) from biz_article a
inner join biz_type bt on a.type_id = bt.id
where 1=1
` + filter

	err := db.Db.Get(&count, countSql, params...)
	if err != nil {
		log.Panicln("count biz_article err: ", err.Error())
	}

	offset, limit := req.GetLimit()
	params = append(params, offset)
	params = append(params, limit)

	if err := db.Db.Select(&list, dataSql, params...); err != nil {
		log.Panicln("select biz_article err: " + err.Error())
	}

	return
}

func GetById(id int) (article BizArticle) {
	dataSql := `
select id,
       title,
       user_id,
       cover_image,
       qrcode_path,
       is_markdown,
       content,
       content_md,
       top,
       type_id,
       status,
       recommended,
       original,
       description,
       keywords,
       comment,
       create_time,
       update_time
from biz_article
where id = ?
`
	if err := db.Db.Get(&article, dataSql, id); err != nil {
		log.Panicln("select biz_article err: " + err.Error())
	}
	return
}

func GetDetailById(id int) (detail ArticleDetail, err error) {
	article := GetById(id)
	detail = ArticleDetail{BizArticle: article}

	curType, typeErr := bizType.GetById(article.TypeId)
	if typeErr != nil {
		return ArticleDetail{}, typeErr
	}
	tags, tagErr := bizTags.GetByArticleId(id)
	if tagErr != nil {
		return ArticleDetail{}, tagErr
	}

	detail.BizType = curType
	detail.BizTags = tags
	return
}

func AddArticle(model AddArticleModel) (err error) {
	// todo 获取当前用户ID
	model.UserId = 1
	tx := db.Db.MustBegin()
	defer tx.Rollback()

	newId, insertErr := insertArticle(model, tx)
	if insertErr != nil {
		return insertErr
	}

	insertTagErr := setupArticleTag(newId, model.TagIds, tx)
	if insertTagErr != nil {
		return insertTagErr
	}

	_ = tx.Commit()
	return nil
}

func EditArticle(model EditArticleModel) (err error) {
	// todo 获取当前用户ID
	model.UserId = 1

	tx := db.Db.MustBegin()
	defer tx.Rollback()

	updateErr := updateArticle(model, tx)
	if updateErr != nil {
		return updateErr
	}

	insertTagErr := setupArticleTag(model.Id, model.TagIds, tx)
	if insertTagErr != nil {
		return insertTagErr
	}
	_ = tx.Commit()
	return nil
}

func updateArticle(model EditArticleModel, tx *sqlx.Tx) (err error) {
	updateSql := `
UPDATE biz_article
SET title       = :title,
    user_id     = :user_id,
    cover_image = :cover_image,
    is_markdown = :is_markdown,
    content     = :content,
    content_md  = :content_md,
    type_id     = :type_id,
    status      = :status,
    original    = :original,
    description = :description,
    comment     = :comment,
    update_time = now()
WHERE id = :id
`
	result, updateErr := tx.NamedExec(updateSql, model)
	if updateErr != nil {
		return updateErr
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("更新失败")
	}
	return nil
}

func insertArticle(model AddArticleModel, tx *sqlx.Tx) (newId int, err error) {
	insertSql := `
insert into biz_article (title, user_id, cover_image, qrcode_path, is_markdown, content, content_md, top, type_id,
                         status, recommended, original, description, keywords, comment, create_time)
values (:title, :user_id, :cover_image, '', :is_markdown, :content, :content_md, 0, :type_id,
        :status, 0, :original, :description, '', :comment, now());
`

	result, insertErr := tx.NamedExec(insertSql, model)
	if insertErr != nil {
		err = insertErr
		return
	}
	id, insertIdErr := result.LastInsertId()
	if insertIdErr != nil {
		err = insertIdErr
		return
	}
	newId = int(id)
	return
}

func setupArticleTag(articleId int, tagIds []int, tx *sqlx.Tx) error {
	deleteSql := `
delete from biz_article_tags
where article_id = ?
`
	_, deleteErr := tx.Exec(deleteSql, articleId)
	if deleteErr != nil {
		return deleteErr
	}

	insertSql := `
insert into biz_article_tags (tag_id, article_id, create_time)
values (:tag_id, :article_id, now());
`
	for _, id := range tagIds {
		_, err := tx.NamedExec(insertSql, map[string]interface{}{
			"tag_id":     id,
			"article_id": articleId,
		})
		return err
	}
	return nil
}
