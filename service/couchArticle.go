package service

import (
	"bytes"
	"encoding/json"
	"strconv"
)

//ArticleDocument 文章文档
type ArticleDocument struct {
	Article                                Article                                 `json:"article"`
	ArticleFragmentRelationDocumentDetails []ArticleFragmentRelationDocumentDetail `json:"relations"`
	ArticleArticleRelationDetails          []ArticleArticleRelationDetail          `json:"dependences"`
}

const (
	articleCouchdbName = "article"
)

//CouchdbArticleGenerateDocument 生成文档数据
func CouchdbArticleGenerateDocument(articleID int) int {

	article := GetArticleByID(articleID, 0)
	if article.ID == 0 {
		return 1
	}
	var articleDocument ArticleDocument
	articleDocument.Article = article
	articleFragmentRelationDocumentDetails := queryArticleFragmentRelationDocumentDetails(articleID)
	articleDocument.ArticleFragmentRelationDocumentDetails = articleFragmentRelationDocumentDetails
	var detail ArticleArticleRelationDetail
	detail.ArticleID = articleID
	articleArticleRelationDetails := queryArticleArticleRelationDetails(detail)
	articleDocument.ArticleArticleRelationDetails = articleArticleRelationDetails
	articleIDStr := strconv.Itoa(articleID)
	articleDocumentBody, _ := json.Marshal(articleDocument)
	code := couchdbCreateDoc(articleCouchdbName, articleIDStr, bytes.NewReader(articleDocumentBody))
	return code
}

//CouchdbGetArticleDocumentByArticleID 通过文章ID查询文档
func CouchdbGetArticleDocumentByArticleID(articleID int) ArticleDocument {
	articleIDStr := strconv.Itoa(articleID)
	articleDocumentBytes := couchdbQueryDocByID(articleCouchdbName, articleIDStr)
	var articleDocument ArticleDocument
	json.Unmarshal(articleDocumentBytes, &articleDocument)
	return articleDocument
}
