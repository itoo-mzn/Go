package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strings"

	opensearch "github.com/opensearch-project/opensearch-go"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

const IndexName = "go-test-index1"

func main() {
	// opensearchクライアント作成
	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Addresses: []string{"https://localhost:9200"},
		Username:  "admin",
		Password:  "admin",
	})
	if err != nil {
		log.Fatal("cannot initialize", err)
	}
	fmt.Println(client.Info())

	// index作成
	settings := strings.NewReader(`{
		'settings': {
			'index': {
				'number_of_shards': 1,
				'number_of_replicas': 2
			}
		}
	}`)
	res := opensearchapi.IndicesCreateRequest{
		Index: IndexName,
		Body: settings,
	}
	fmt.Println("Create Index", res)

	// indexにdocument追加
	doc := strings.NewReader(`{
		"title": "Moneyball",
		"direcror": "Bennett Miller",
		"year": "2021"
	}`)
	docID := "1"
	req := opensearchapi.IndexRequest{
		Index: IndexName,
		DocumentID: docID,
		Body: doc,
	}
	insertRes, err := req.Do(context.Background(), client)
	if err != nil {
		log.Fatal("fail insert document", err)
	}
	fmt.Println("insert: ", insertRes)

	// 一括操作（index, create, update）
	blk, err := client.Bulk(
		strings.NewReader(`
		{ "index" : { "_index" : "go-test-index1", "_id" : "2" } }
    { "title" : "Interstellar", "director" : "Christopher Nolan", "year" : "2014"}
    { "create" : { "_index" : "go-test-index1", "_id" : "3" } }
    { "title" : "Star Trek Beyond", "director" : "Justin Lin", "year" : "2015"}
    { "update" : {"_id" : "3", "_index" : "go-test-index1" } }
    { "doc" : {"year" : "2016"} }
`),
	)
	// NOTE: ↑↑ 2つ上の行のインデント気になるけど、整列させたらエラーになる
	if err != nil {
		log.Fatal("fail bulk operations", err)
	}
	fmt.Println("bulk: ", blk)

	// 検索
	// titleのスコアはdirecrorの2倍
	q := strings.NewReader(`{
		"size": 5,
		"query": {
			"multi_match": {
				"query": "miller",
				"fields": ["title^2", "director"]
			}
		}
	}`)
	search := opensearchapi.SearchRequest{
		Index: []string{IndexName},
		Body: q,
	}
	searchRes, err := search.Do(context.Background(), client)
	if err != nil {
		log.Fatal("fail bulk operations", err)
	}
	fmt.Println("search: ", searchRes)
	defer searchRes.Body.Close()

	// 削除
	delete := opensearchapi.DeleteRequest{
		Index: IndexName,
		DocumentID: docID,
	}
	deleteRes, err := delete.Do(context.Background(), client)
	if err != nil {
		log.Fatal("fail delete", err)
	}
	fmt.Println("delete: ", deleteRes)
	defer deleteRes.Body.Close()
}

