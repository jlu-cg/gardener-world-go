package service

//ArticleArticleRelationDetail 文章关联详情
type ArticleArticleRelationDetail struct {
	ID              int    `json:"id"`
	ArticleID       int    `json:"articleId"`
	RelateArticleID int    `json:"relateArticleID"`
	Pos             int16  `json:"pos"`
	Title           string `json:"title"`
}

const (
	queryArticleDependenceDetailsSQL = "select a.id, a.article_id, a.pos, b.id, b.title from article_article_relation a inner join article b on a.relate_article_id=b.id "
	addArticleDependenceSQL          = "insert into article_article_relation(article_id, relate_article_id, pos)values($1, $2, $3)"
	deleteArticleDependenceSQL       = "delete from article_article_relation where id=$1"
	updateArticleDependencePosSQL    = "update article_article_relation set pos=$1 where id=$2"
	queryArticleDependenceByIDSQL    = "select id, article_id, relate_article_id, pos from article_article_relation where id=$1"
)

func queryArticleDependenceDetails(detail ArticleArticleRelationDetail) []ArticleArticleRelationDetail {

	var articleDependenceDetails []ArticleArticleRelationDetail
	whereSQL := " where 1=1 "
	if detail.ArticleID > 0 {
		whereSQL += " and a.article_id=" + intToSafeString(detail.ArticleID)
	} else {
		return articleDependenceDetails
	}

	whereSQL += " order by a.pos asc "
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleDependenceDetailsSQL + whereSQL)
	defer rows.Close()
	if rows == nil {
		return articleDependenceDetails
	}
	if err != nil {
		panic(err)
	}

	var temp ArticleArticleRelationDetail
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.ArticleID, &temp.Pos, &temp.RelateArticleID, &temp.Title)
		articleDependenceDetails = append(articleDependenceDetails, temp)
	}

	return articleDependenceDetails
}

func addArticleDependence(detail ArticleArticleRelationDetail) int {
	if detail.ID > 0 {
		queryDetail := queryArticleDependenceDetailByID(detail.ID)
		if queryDetail.ID <= 0 {
			return -1
		}
	}
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleDependenceSQL)
	if err != nil {
		return -1
	}

	_, err = stmt.Exec(detail.ArticleID, detail.RelateArticleID, detail.Pos)
	if err != nil {
		return -1

	}
	return 0
}

func updateArticleDependences(details []ArticleArticleRelationDetail) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleDependencePosSQL)
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

func deleteArticleDependenceByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleDependenceSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}

func queryArticleDependenceDetailByID(ID int) ArticleArticleRelationDetail {
	var articleDependenceDetail ArticleArticleRelationDetail
	connection := connect()
	defer release(connection)

	rows, err := connection.Query(queryArticleDependenceByIDSQL, ID)
	defer rows.Close()
	if rows == nil {
		return articleDependenceDetail
	}
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		rows.Scan(&articleDependenceDetail.ID, &articleDependenceDetail.ArticleID, &articleDependenceDetail.RelateArticleID, &articleDependenceDetail.Pos)
	}

	return articleDependenceDetail
}
