package service

//ArticleTagRelation 文章标签关系
type ArticleTagRelation struct {
	ID        int `json:"id"`
	ArticleID int `json:"articleId"`
	TagID     int `json:"tagId"`
}

const (
	queryArticleTagRelationsSQL = "select id, article_id, tag_id from article_tag_relation "
	addArticleTagRelationSQL    = "insert into article_tag_relation(article_id, tag_id) values($1, $2)"
	updateArticleTagRelationSQL = "update article_tag_relation article_id=$1, tag_id=$2 where id=$3"
	deleteArticleTagRelationSQL = "delete from article_tag_relation where id=$1"
)

func addArticleTagRelation(relation ArticleTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleTagRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.ArticleID, relation.TagID)
	if err != nil {
		return -1

	}
	return 0
}

func updateArticleTagRelation(relation ArticleTagRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleTagRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.ArticleID, relation.TagID, relation.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteArticleTagRelationByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleTagRelationSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
