package service

//Article 文章
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	TagID   int    `json:"tagId"`
}

const (
	queryArticlesSQL     = "select id, title, summary, tag_id from article "
	addArticleSQL        = "insert into article(title, summary, tag_id)values($1, $2, $3)"
	updateArticleSQL     = "update article title=$1, summary=$2, tag_id=$3 where id=$4"
	deleteArticleByIDSQL = "delete from article where id=$1"
)

func queryArticles(article Article, lastID int) []Article {
	connection := connect()
	defer release(connection)

	whereSQL := " where "

	if article.ID > 0 {
		whereSQL += " id=" + intToSafeString(article.ID) + " and "
	}

	if article.TagID > 0 {
		whereSQL += " tag_id=" + intToSafeString(article.TagID) + " and "
	}
	whereSQL += " id>" + intToSafeString(lastID) + " limit 20 "
	rows, err := connection.Query(queryArticlesSQL + whereSQL)
	defer rows.Close()
	var articles []Article
	if rows == nil {
		return articles
	}
	if err != nil {
		panic(err)
	}

	var temp Article
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Title, &temp.Summary, &temp.TagID)
		articles = append(articles, temp)
	}

	return articles
}

func addArticle(article Article) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addArticleSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(article.Title, article.Summary, article.TagID)
	if err != nil {
		return -1
	}
	return 0
}

func updateArticle(article Article) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateArticleSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(article.Title, article.Summary, article.TagID, article.ID)
	if err != nil {
		return -1
	}
	return 0
}

func queryArticleByID(articleID int) Article {
	var article Article
	article.ID = articleID
	articles := queryArticles(article, 0)
	return articles[0]
}

func deleteArticleByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteArticleByIDSQL)
	if err != nil {
		return -1
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return -1
	}
	return 0
}
