package service

//Article 文章
type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

const (
	queryArticlesSQL     = "select id, title from article "
	addArticleSQL        = "insert into article(title)values($1)"
	updateArticleSQL     = "update article set title=$1 where id=$4"
	deleteArticleByIDSQL = "delete from article where id=$1"
)

func queryArticles(article Article, lastID int) []Article {
	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if article.ID > 0 {
		whereSQL += " and id=" + intToSafeString(article.ID)
	}

	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
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
		rows.Scan(&temp.ID, &temp.Title)
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
	_, err = stmt.Exec(article.Title)
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
	_, err = stmt.Exec(article.Title, article.ID)
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
