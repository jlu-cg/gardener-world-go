package service

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type articleDocument struct {
	Article                          Article                           `json:"article"`
	ArticleRelationDocumentDetails   []ArticleRelationDocumentDetail   `json:"relations"`
	ArticleDependenceDocumentDetails []ArticleDependenceDocumentDetail `json:"dependences"`
}

const (
	articleCouchdbName = "article"
)

//CouchdbArticleGenerateDocument 生成文档数据
func CouchdbArticleGenerateDocument(articleID int) int {

	article := GetArticleByID(articleID, 0)
	couchdbGetAuth()
	if article.ID == 0 {
		return 1
	}
	var articleDocument articleDocument
	articleDocument.Article = article
	articleRelationDocumentDetails := queryArticleRelationDocumentDetails(articleID)
	articleDocument.ArticleRelationDocumentDetails = articleRelationDocumentDetails
	articleDependenceDocumentDetails := queryArticleDependenceDocumentDetails(articleID)
	articleDocument.ArticleDependenceDocumentDetails = articleDependenceDocumentDetails
	articleIDStr := strconv.Itoa(articleID)
	articleDocumentBody, _ := json.Marshal(articleDocument)
	code := couchdbCreateDoc(articleCouchdbName, articleIDStr, bytes.NewReader(articleDocumentBody))
	return code
}
