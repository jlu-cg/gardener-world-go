package service

import (
	"bytes"
	"encoding/json"
	"strconv"
)

//ArticleDocument 文章文档
type ArticleDocument struct {
	ArticleDocumentAdd
	Rev string `json:"_rev"`
}

type ArticleDocumentAdd struct {
	Article                                Article                                 `json:"article"`
	ArticleFragmentRelationDocumentDetails []ArticleFragmentRelationDocumentDetail `json:"relations"`
	ArticleArticleRelationDetails          []ArticleArticleRelationDetail          `json:"dependences"`
}

const (
	articleCouchdbName = "article"
)

//CouchdbArticleGenerateDocument 生成文档数据
func CouchdbArticleGenerateDocument(articleID int) int {

	article := QueryArticleByID(articleID)
	if article.ID == 0 {
		return 1
	}

	article.Status = 1
	article.ID = articleID
	SaveArticle(article)

	articleIDStr := strconv.Itoa(articleID)
	oldArticleDocument := CouchdbGetArticleDocumentByArticleID(articleID)

	articleFragmentRelationDocumentDetails := queryArticleFragmentRelationDocumentDetails(articleID)
	var detail ArticleArticleRelationDetail
	detail.ArticleID = articleID
	articleArticleRelationDetails := queryArticleArticleRelationDetails(detail)

	var code int
	if oldArticleDocument.Rev != "" {
		var articleDocument ArticleDocument
		articleDocument.Article = article
		articleDocument.ArticleArticleRelationDetails = articleArticleRelationDetails
		articleDocument.ArticleFragmentRelationDocumentDetails = articleFragmentRelationDocumentDetails
		articleDocument.Rev = oldArticleDocument.Rev
		articleDocumentBody, _ := json.Marshal(articleDocument)
		code = couchdbUpdateDoc(articleCouchdbName, articleIDStr, bytes.NewReader(articleDocumentBody))
	} else {
		var articleDocumentAdd ArticleDocumentAdd
		articleDocumentAdd.Article = article
		articleDocumentAdd.ArticleArticleRelationDetails = articleArticleRelationDetails
		articleDocumentAdd.ArticleFragmentRelationDocumentDetails = articleFragmentRelationDocumentDetails
		articleDocumentBody, _ := json.Marshal(articleDocumentAdd)
		code = couchdbCreateDoc(articleCouchdbName, articleIDStr, bytes.NewReader(articleDocumentBody))
	}
	return code
}

//CouchdbArticleCancelDocument 取消文章发布
func CouchdbArticleCancelDocument(articleID int) int {

	if articleID == 0 {
		return 1
	}

	var article Article
	article.Status = 2
	article.ID = articleID
	SaveArticle(article)

	articleIDStr := strconv.Itoa(articleID)
	oldArticleDocument := CouchdbGetArticleDocumentByArticleID(articleID)
	var code int
	if oldArticleDocument.Rev != "" {
		code = couchdbDeleteDoc(articleCouchdbName, articleIDStr, oldArticleDocument.Rev)
		return code
	} else {
		return 1
	}
}

//CouchdbGetArticleDocumentByArticleID 通过文章ID查询文档
func CouchdbGetArticleDocumentByArticleID(articleID int) ArticleDocument {
	articleIDStr := strconv.Itoa(articleID)
	articleDocumentBytes := couchdbQueryDocByID(articleCouchdbName, articleIDStr)
	var articleDocument ArticleDocument
	json.Unmarshal(articleDocumentBytes, &articleDocument)
	return articleDocument
}
