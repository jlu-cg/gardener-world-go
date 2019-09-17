package service

//ArticleFragmentAttributeRelation 文章碎片对应的可对比属性
type ArticleFragmentAttributeRelation struct {
	ID             int `json:"id"`
	FragmentID     int `json:"fragmentId"`
	ArticleID      int `json:"articleId"`
	TagAttributeID int `json:"tagAttributeId"`
}

const (
	queryArticleFragmentAttributeRelationsSQL = "select id, fragment_id, article_id, tag_attribute_id from article_fragment_attribute_relation "
	addArticleFragmentAttributeRelationSQL    = "insert into article_fragment_attribute_relation(fragment_id, article_id, tag_attribute_id) values($1, $2, $3)"
	updateArticleFragmentAttributeRelationSQL = "update article_fragment_attribute_relation fragment_id=$1, article_id=$2, tag_attribute_id=$3 where id=$4"
	deleteArticleFragmentAttributeRelationSQL = "delete from article_fragment_attribute_relation where id=$1"
)

func addArticleFragmentAttributeRelation(relation ArticleFragmentAttributeRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleFragmentAttributeRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.FragmentID, relation.ArticleID, relation.TagAttributeID)
	if err != nil {
		return -1

	}
	return 0
}

func updateArticleFragmentAttributeRelation(relation ArticleFragmentAttributeRelation) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleFragmentAttributeRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(relation.FragmentID, relation.ArticleID, relation.TagAttributeID, relation.ID)
	if err != nil {
		return -1
	}
	return 0
}

func deleteArticleFragmentAttributeRelationByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleFragmentAttributeRelationSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
