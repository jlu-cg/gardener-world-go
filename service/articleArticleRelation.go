package service

//ArticleArticleRelationDetail 文章关联详情
type ArticleArticleRelationDetail struct {
	ID              int    `json:"id"`
	ArticleID       int    `json:"articleId"`
	RelateArticleID int    `json:"relateArticleId"`
	Position        int16  `json:"position"`
	Title           string `json:"title"`
}

const (
	queryArticleArticleRelationDetailsSQL = "select a.id, a.article_id, a.position, b.id, b.title from article_article_relation a inner join article b on a.relate_article_id=b.id "
	addArticleArticleRelationSQL          = "insert into article_article_relation(article_id, relate_article_id, position)values($1, $2, $3)"
	deleteArticleArticleRelationSQL       = "delete from article_article_relation"
	updateArticleArticleRelationPosSQL    = "update article_article_relation set position=$1 where id=$2"
	queryArticleArticleRelationByIDSQL    = "select id, article_id, relate_article_id, position from article_article_relation where id=$1"
)

func queryArticleArticleRelationDetails(detail ArticleArticleRelationDetail) []ArticleArticleRelationDetail {

	var articleArticleRelationDetails []ArticleArticleRelationDetail
	whereSQL := " where 1=1 "
	if detail.ArticleID > 0 {
		whereSQL += " and a.article_id=" + intToSafeString(detail.ArticleID)
	} else {
		return articleArticleRelationDetails
	}

	whereSQL += " order by a.position asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleArticleRelationDetailsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleArticleRelationDetails
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleArticleRelationDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.Position, &temp.RelateArticleID, &temp.Title)
		articleArticleRelationDetails = append(articleArticleRelationDetails, temp)
	}

	return articleArticleRelationDetails
}

func addArticleArticleRelation(detail ArticleArticleRelationDetail) int {
	if detail.ID > 0 {
		queryDetail := queryArticleArticleRelationDetailByID(detail.ID)
		if queryDetail.ID <= 0 {
			return -1
		}
	}
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleArticleRelationSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(detail.ArticleID, detail.RelateArticleID, detail.Position)
	if err != nil {
		return -1

	}
	return 0
}

func updateArticleArticleRelations(details []ArticleArticleRelationDetail) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleArticleRelationPosSQL)
	if err != nil {
		return -1
	}

	for _, detail := range details {
		_, err = stmt.Exec(detail.Position, detail.ID)
		if err != nil {
			return -1
		}
	}
	return 0
}

func deleteArticleArticleRelations(relation ArticleArticleRelationDetail) int {
	hasCondition := false
	whereSQL := " where 1=1 "

	if relation.ID > 0 {
		whereSQL += " and id=" + intToSafeString(relation.ID)
		hasCondition = true
	}

	if !hasCondition {
		return -1
	}

	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleArticleRelationSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec()
	if err != nil {
		return -1
	}
	return 0
}

func queryArticleArticleRelationDetailByID(ID int) ArticleArticleRelationDetail {
	var articleArticleRelationDetail ArticleArticleRelationDetail
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleArticleRelationByIDSQL, ID)
	defer rows.Close()
	if rows == nil {
		return articleArticleRelationDetail
	}
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&articleArticleRelationDetail.ID, &articleArticleRelationDetail.ArticleID, &articleArticleRelationDetail.RelateArticleID, &articleArticleRelationDetail.Position)
	}

	return articleArticleRelationDetail
}
