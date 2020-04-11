package main

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"reflect"
	"time"
)

type BookDocument struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	PublishedAt time.Time `json:"published_at"`
}

func main() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}

	//res, err := client.CreateIndex("book_shelf").Do(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	b := BookDocument{
		Title:       "マイクロサービスアーキテクチャ",
		Author:      "Sam Newman",
		PublishedAt: time.Now(),
	}
	_, err = client.Index().Index("book_shelf").Id("5").BodyJson(b).Do(context.Background())
	if err != nil {
		panic(err)
	}
	//print(res.Index)

	query := elastic.NewMultiMatchQuery("マイ", "title").Type("phrase_prefix")

	//query := elastic.NewTermQuery("title", "クロ")
	searchResult, err := client.Search().
		Index("book_shelf").
		Query(query).
		From(0).
		Size(10).
		Pretty(true).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var bookType BookDocument
	for _, item := range searchResult.Each(reflect.TypeOf(bookType)){
		if b, ok := item.(BookDocument); ok {
			fmt.Printf("Tweet by %s: %s\n", b.Title, b.Author)
		}
	}
}
