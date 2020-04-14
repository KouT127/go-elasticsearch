package main

import (
	"fmt"
	"go-elasticsearch/modules/elasticsearch"
	"time"
)

func main() {
	elasticsearch.InitClientWithLogger("http://localhost:9200")
	var err error
	docs, err := elasticsearch.SearchBooks("マイ")
	if err != nil {
		panic(err)
	}
	for _, doc := range docs {
		fmt.Printf("book by title: %s author: %s \n", doc.Title, doc.Author)
	}

	docs = make([]elasticsearch.BookDocument, 100)
	for i := 0; i < 100; i++ {
		docs[i] = elasticsearch.BookDocument{
			Title:       "マイクロサービスアーキテクチャ",
			Author:      "Sam Newman",
			PublishedAt: time.Now(),
		}
	}

	elasticsearch.BulkCreateBooks(docs)
}
