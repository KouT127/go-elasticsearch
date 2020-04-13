package elasticsearch

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

func (d BookDocument) create() error {
	b := BookDocument{
		Title:       "マイクロサービスアーキテクチャ",
		Author:      "Sam Newman",
		PublishedAt: time.Now(),
	}
	ctx := context.Background()
	_, err = client.Index().Index(bookDocumentIndexName).BodyJson(b).Do(ctx)
	if err != nil {
		return fmt.Errorf("create documents failed: %s", err)
	}
	return nil
}

func SearchBook(title string) ([]BookDocument, error) {
	var bookType BookDocument
	docs := make([]BookDocument, 10)
	ctx := context.Background()
	query := elastic.
		NewMultiMatchQuery(title, "title").
		Type("phrase_prefix")

	searchResult, err := client.Search().
		Index(bookDocumentIndexName).
		Query(query).
		From(0).
		Size(10).
		Pretty(true).
		Do(ctx)

	if err != nil {
		return nil, fmt.Errorf("search failed: %s", err)
	}

	for idx, item := range searchResult.Each(reflect.TypeOf(bookType)) {
		if book, ok := item.(BookDocument); ok {
			docs[idx] = book
		}
	}
	return docs, nil
}
