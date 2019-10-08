package service

import "github.com/gardener/gardener-world-go/config"

//ArticleFragmentRelationDetail 文章关联详情
type ArticleFragmentRelationDetail struct {
	ID         int    `json:"id"`
	ArticleID  int    `json:"articleId"`
	FragmentID int    `json:"fragmentId"`
	Position   int16  `json:"position"`
	Title      string `json:"title"`
}

//ArticleFragmentRelationDocumentDetail 生成文档关联
type ArticleFragmentRelationDocumentDetail struct {
	ID         int    `json:"id"`
	FragmentID int    `json:"fragmentId"`
	Position   int16  `json:"position"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

const (
	queryArticleFragmentRelationDetailsSQL         = "select a.id, a.article_id, a.position, b.id, b.title from article_fragment_relation a inner join fragment b on a.fragment_id=b.id "
	queryArticleFragmentRelationDocumentDetailsSQL = "select a.id, a.position, b.id, b.title, b.content from article_fragment_relation a inner join fragment b on a.fragment_id=b.id where a.article_id=$1 order by a.position asc"
	queryArticleFragmentRelationDetailByIDSQL      = "select id, article_id, fragment_id, position from article_fragment_relation where id=$1"
	addArticleFragmentRelationSQL                  = "insert into article_fragment_relation(article_id, fragment_id, position)values($1, $2, $3)"
	updateArticleFragmentRelationPosSQL            = "update article_fragment_relation set position=$1 where id=$2"
	deleteArticleFragmentRelationSQL               = "delete from article_fragment_relation"
)

func queryArticleFragmentRelationDetails(detail ArticleFragmentRelationDetail) []ArticleFragmentRelationDetail {

	var articleRelationDetails []ArticleFragmentRelationDetail
	whereSQL := " where 1=1 "
	if detail.ArticleID > 0 {
		whereSQL += " and a.article_id=" + intToSafeString(detail.ArticleID)
	} else {
		return articleRelationDetails
	}

	whereSQL += " order by a.position asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleFragmentRelationDetailsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleRelationDetails
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleFragmentRelationDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.Position, &temp.FragmentID, &temp.Title)
		articleRelationDetails = append(articleRelationDetails, temp)
	}

	return articleRelationDetails
}

func addArticleFragmentRelation(detail ArticleFragmentRelationDetail) int {
	queryDetail := queryArticleFragmentRelationDetailByID(detail.ID)
	if queryDetail.ID > 0 {
		return config.DBErrorNoData
	}
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleFragmentRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	_, err = stmt.Exec(detail.ArticleID, detail.FragmentID, detail.Position)
	if err != nil {
		return config.DBErrorExecution

	}
	return config.DBSuccess
}

func updateArticleFragmentRelations(details []ArticleFragmentRelationDetail) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleFragmentRelationPosSQL)
	if err != nil {
		return config.DBErrorConnection
	}

	for _, detail := range details {
		_, err = stmt.Exec(detail.Position, detail.ID)
		if err != nil {
			return config.DBErrorExecution
		}
	}
	return config.DBSuccess
}

func deleteArticleFragmentRelations(relation ArticleFragmentRelationDetail) int {
	hasCondition := false
	whereSQL := " where 1=1 "

	if relation.ID > 0 {
		whereSQL += " and id=" + intToSafeString(relation.ID)
		hasCondition = true
	}

	if !hasCondition {
		return config.DBErrorSQLNoCondition
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleFragmentRelationSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec()
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func queryArticleFragmentRelationDetailByID(ID int) ArticleFragmentRelationDetail {
	var articleFragmentRelationDetail ArticleFragmentRelationDetail
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleFragmentRelationDetailByIDSQL, ID)
	defer rows.Close()
	if rows == nil {
		return articleFragmentRelationDetail
	}
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&articleFragmentRelationDetail.ID, &articleFragmentRelationDetail.ArticleID, &articleFragmentRelationDetail.FragmentID, &articleFragmentRelationDetail.Position)
	}

	return articleFragmentRelationDetail
}

func queryArticleFragmentRelationDocumentDetails(articleID int) []ArticleFragmentRelationDocumentDetail {

	var articleFragmentRelationDocumentDetails []ArticleFragmentRelationDocumentDetail
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleFragmentRelationDocumentDetailsSQL, articleID)
	defer rows.Close()
	if rows == nil {
		return articleFragmentRelationDocumentDetails
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleFragmentRelationDocumentDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Position, &temp.FragmentID, &temp.Title, &temp.Content)
		articleFragmentRelationDocumentDetails = append(articleFragmentRelationDocumentDetails, temp)
	}

	return articleFragmentRelationDocumentDetails
}
