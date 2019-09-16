package service

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gardener/gardener-world-go/config"
)

type couchdbResp struct {
	Ok   bool   `json:"ok"`
	ID   string `json:"id"`
	Rev  string `json:"rev"`
	Name string `json:"name"`
}

var (
	couchdbURL  string
	userName    string
	password    string
	authSession string
)

const (
	couchdbAuthFormat           = "%s/_session"
	couchdbDatabaseOpFormat     = "%s/%s"
	couchdbDocumentCreateFormat = "%s/%s/%s"
	couchdbDocumentUpdateFormat = "%s/%s/%s/"
	couchdbDocumentDeleteFormat = "%s/%s/%s?rev=%s"
	couchdbDocumentQueryFormat  = "%s/%s/%s"
)

//InitCouchDb 初始化数据配置
func InitCouchDb(config *config.WorldConfig) {
	couchdbURL = config.CouchdbConfig.URL
	userName = config.CouchdbConfig.UserName
	password = config.CouchdbConfig.Password
}

func couchdbGetAuth() {
	client := &http.Client{}
	authStr := fmt.Sprintf("{\"name\":\"%s\", \"password\":\"%s\"}", userName, password)
	body := strings.NewReader(authStr)
	requestURL := fmt.Sprintf(couchdbAuthFormat, couchdbURL)
	request, _ := http.NewRequest("POST", requestURL, body)
	request.Header.Add("Content-Type", "application/json")
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	authSession = response.Header.Get("Set-Cookie")
	fmt.Printf("authSession:%s", authSession)
	couchdbRes := couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	fmt.Printf("ok:%t,name:%s", couchdbRes.Ok, couchdbRes.Name)
}

func couchdbCreateDb(dbName string) int {
	client := &http.Client{}
	requestURL := fmt.Sprintf(couchdbDatabaseOpFormat, couchdbURL, dbName)
	request, _ := http.NewRequest("PUT", requestURL, nil)
	if authSession == "" {
		couchdbGetAuth()
	}
	request.Header.Add("cookie", authSession)
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	couchdbRes := couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	couchdbGetAuth()
	request.Header.Add("cookie", authSession)
	response, _ = client.Do(request)
	b, _ = ioutil.ReadAll(response.Body)
	couchdbRes = couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	return -1
}

func couchdbDeleteDb(dbName string) int {
	client := &http.Client{}
	requestURL := fmt.Sprintf(couchdbDatabaseOpFormat, couchdbURL, dbName)
	request, _ := http.NewRequest("DELETE", requestURL, nil)
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	couchdbRes := couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}

	couchdbGetAuth()
	request.Header.Add("cookie", authSession)
	response, _ = client.Do(request)
	b, _ = ioutil.ReadAll(response.Body)
	couchdbRes = couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	return 1
}

func couchdbCreateDoc(dbName string, ID string, document io.Reader) int {
	client := &http.Client{}
	requestURL := fmt.Sprintf(couchdbDocumentCreateFormat, couchdbURL, dbName, ID)
	request, _ := http.NewRequest("PUT", requestURL, document)
	if authSession == "" {
		couchdbGetAuth()
	}
	request.Header.Add("cookie", authSession)
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	couchdbRes := couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	fmt.Printf("ok:%t", couchdbRes.Ok)
	if couchdbRes.Ok {
		return 1
	}
	couchdbGetAuth()
	request.Header.Add("cookie", authSession)
	response, _ = client.Do(request)
	b, _ = ioutil.ReadAll(response.Body)
	couchdbRes = couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	return -1
}

func couchdbUpdateDoc(dbName string, ID string, document io.Reader) int {
	client := &http.Client{}
	requestURL := fmt.Sprintf(couchdbDocumentUpdateFormat, couchdbURL, dbName, ID)
	request, _ := http.NewRequest("PUT", requestURL, document)
	if authSession == "" {
		couchdbGetAuth()
	}
	request.Header.Add("cookie", authSession)
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	couchdbRes := couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	couchdbGetAuth()
	request.Header.Add("cookie", authSession)
	response, _ = client.Do(request)
	b, _ = ioutil.ReadAll(response.Body)
	couchdbRes = couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	return -1
}

func couchdbDeleteDoc(dbName string, ID string, rev string) int {
	client := &http.Client{}
	requestURL := fmt.Sprintf(couchdbDocumentDeleteFormat, couchdbURL, dbName, ID, rev)
	request, _ := http.NewRequest("DELETE", requestURL, nil)
	if authSession == "" {
		couchdbGetAuth()
	}
	request.Header.Add("cookie", authSession)
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	couchdbRes := couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	couchdbGetAuth()
	request.Header.Add("cookie", authSession)
	response, _ = client.Do(request)
	b, _ = ioutil.ReadAll(response.Body)
	couchdbRes = couchdbResp{}
	json.Unmarshal(b, &couchdbRes)
	if couchdbRes.Ok {
		return 1
	}
	return -1
}

func couchdbQueryDocByID(dbName string, ID string) []byte {
	client := &http.Client{}
	requestURL := fmt.Sprintf(couchdbDocumentQueryFormat, couchdbURL, dbName, ID)
	request, _ := http.NewRequest("GET", requestURL, nil)
	if authSession == "" {
		couchdbGetAuth()
	}
	request.Header.Add("cookie", authSession)
	response, _ := client.Do(request)
	b, _ := ioutil.ReadAll(response.Body)
	return b
}
