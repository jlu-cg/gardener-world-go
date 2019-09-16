package service

//ArticleRelationDetail 文章关联详情
type ArticleRelationDetail struct {
	ID         int    `json:"id"`
	ArticleID  int    `json:"articleId"`
	FragmentID int    `json:"fragmentId"`
	Pos        int16  `json:"pos"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
}

//ArticleRelationDocumentDetail 生成文档关联
type ArticleRelationDocumentDetail struct {
	ID         int    `json:"id"`
	FragmentID int    `json:"fragmentId"`
	Pos        int16  `json:"pos"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}

const (
	queryArticleRelationDetailsSQL         = "select a.id, a.article_id, a.pos, b.id, b.title, b.summary from article_relation a inner join fragment b on a.fragment_id=b.id "
	queryArticleRelationDocumentDetailsSQL = "select a.id, a.pos, b.id, b.title, b.content from article_relation a inner join fragment b on a.fragment_id=b.id where a.article_id=$1 order by a.pos asc"
	addArticleRelationSQL                  = "insert into article_relation(article_id, fragment_id, pos)values($1, $2, $3)"
	deleteArticleRelationSQL               = "delete from article_relation where id=$1"
	updateArticleRelationPosSQL            = "update article_relation set pos=$1 where id=$2"
	queryArticleRelationDetailByIDSQL      = "select id, article_id, fragment_id, pos from article_relation where id=$1"
)

func queryArticleRelationDetails(detail ArticleRelationDetail) []ArticleRelationDetail {

	var articleRelationDetails []ArticleRelationDetail
	whereSQL := " where 1=1 "
	if detail.ArticleID > 0 {
		whereSQL += " and a.article_id=" + intToSafeString(detail.ArticleID)
	} else {
		return articleRelationDetails
	}

	whereSQL += " order by a.pos asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleRelationDetailsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleRelationDetails
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleRelationDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.Pos, &temp.FragmentID, &temp.Title, &temp.Summary)
		articleRelationDetails = append(articleRelationDetails, temp)
	}

	return articleRelationDetails
}

func addArticleRelation(detail ArticleRelationDetail) int {
	queryDetail := queryArticleRelationDetailByID(detail.ID)
	if queryDetail.ID <= 0 {
		return -1
	}
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(detail.ArticleID, detail.FragmentID, detail.Pos)
	if err != nil {
		return -1

	}
	return 0
}

func updateArticleRelations(details []ArticleRelationDetail) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleRelationPosSQL)
	if err != nil {
		return -1
	}

	for _, detail := range details {
		_, err = stmt.Exec(detail.Pos, detail.ID)
		if err != nil {
			return -1
		}
	}
	return 0
}

func deleteArticleRelationByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleRelationSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}

func queryArticleRelationDetailByID(ID int) ArticleRelationDetail {
	var articleRelationDetail ArticleRelationDetail
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleRelationDetailByIDSQL, ID)
	defer rows.Close()
	if rows == nil {
		return articleRelationDetail
	}
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&articleRelationDetail.ID, &articleRelationDetail.ArticleID, &articleRelationDetail.FragmentID, &articleRelationDetail.Pos)
	}

	return articleRelationDetail
}

func queryArticleRelationDocumentDetails(articleID int) []ArticleRelationDocumentDetail {

	var articleRelationDocumentDetails []ArticleRelationDocumentDetail
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleRelationDocumentDetailsSQL, articleID)
	defer rows.Close()
	if rows == nil {
		return articleRelationDocumentDetails
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleRelationDocumentDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Pos, &temp.FragmentID, &temp.Title, &temp.Content)
		articleRelationDocumentDetails = append(articleRelationDocumentDetails, temp)
	}

	return articleRelationDocumentDetails
}
