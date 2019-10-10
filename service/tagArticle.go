package service

import "github.com/gardener/gardener-world-go/config"

//Tag 标签
type TagArticle struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

const (
	queryTagArticlesSQL     = "select id, name, type from tag_article "
	addTagArticleSQL        = "insert into tag_article(name, type)values($1, $2)"
	updateTagArticleSQL     = "update tag_article set name=$1, type=$2 where id=$3"
	deleteTagArticleByIDSQL = "delete from tag_article where id=$1"
)

func queryTagArticles(tagArticle TagArticle, lastID int) []TagArticle {

	connection := connect()
	defer release(connection)

	whereSQL := " where 1=1 "

	if tagArticle.Name != "" {
		whereSQL += " and name like '" + strToSafeString(tagArticle.Name) + "%' "
	}
	if tagArticle.ID > 0 {
		whereSQL += " and id=" + intToSafeString(tagArticle.ID)
	}
	if tagArticle.Type > 0 {
		whereSQL += " and type=" + intToSafeString(tagArticle.Type)
	}
	if lastID >= 0 {
		whereSQL += " and id>" + intToSafeString(lastID) + " limit 20 "
	}
	rows, err := connection.Query(queryTagArticlesSQL + whereSQL)
	if err != nil {
		panic(err)
	}
	var tagArticles []TagArticle
	if rows == nil {
		return tagArticles
	}
	defer rows.Close()
	if err != nil {
		panic(err)
	}

	var temp TagArticle
	for rows.Next() {
		rows.Scan(&temp.ID, &temp.Name, &temp.Type)
		tagArticles = append(tagArticles, temp)
	}

	return tagArticles
}

//保存标签
func addTagArticle(tagArticle TagArticle) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(addTagArticleSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(tagArticle.Name, tagArticle.Type)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

//更新标签
func updateTagArticle(tagArticle TagArticle) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(updateTagArticleSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(tagArticle.Name, tagArticle.Type)
	if err != nil {
		return config.DBErrorExecution
	}
	return config.DBSuccess
}

func deleteTagArticleByID(id int) int {
	connection := connect()
	defer release(connection)

	stmt, err := connection.Prepare(deleteTagArticleByIDSQL)
	if err != nil {
		return config.DBErrorConnection
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return config.DBErrorExecution
	}

	return config.DBSuccess
}
