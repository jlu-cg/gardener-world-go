package service

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type articleDocument struct {
	Article                                Article                                 `json:"article"`
	ArticleFragmentRelationDocumentDetails []ArticleFragmentRelationDocumentDetail `json:"relations"`
	ArticleFragmentRelationDetails         []ArticleFragmentRelationDetail         `json:"dependences"`
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
	articleFragmentRelationDocumentDetails := queryArticleFragmentRelationDocumentDetails(articleID)
	articleDocument.ArticleFragmentRelationDocumentDetails = articleFragmentRelationDocumentDetails
	var detail ArticleFragmentRelationDetail
	detail.ArticleID = articleID
	ArticleFragmentRelationDetails := queryArticleFragmentRelationDetails(detail)
	articleDocument.ArticleFragmentRelationDetails = ArticleFragmentRelationDetails
	articleIDStr := strconv.Itoa(articleID)
	articleDocumentBody, _ := json.Marshal(articleDocument)
	code := couchdbCreateDoc(articleCouchdbName, articleIDStr, bytes.NewReader(articleDocumentBody))
	return code
}
