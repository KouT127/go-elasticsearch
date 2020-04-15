package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"time"
)

type Author struct {
	Name string `json:"name"`
}

type IndustryIdentifier struct {
	Type       string `json:"type"`
	Identifier int    `json:"identifier"`
}

type BookDocument struct {
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Authors             []Author             `json:"authors"`
	IndustryIdentifiers []IndustryIdentifier `json:"industry_identifiers"`
	ThumbnailURL        string               `json:"thumbnail_url"`
	PublishedAt         time.Time            `json:"published_at"`
}

func (d BookDocument) create() error {
	b := BookDocument{
		Title:       "マイクロサービスアーキテクチャ",
		Authors:     []Author{
			{Name: "Sam Newman"},
		},
		PublishedAt: time.Now(),
	}
	ctx := context.Background()
	_, err = client.Index().Index(bookDocumentIndexName).BodyJson(b).Do(ctx)
	if err != nil {
		return fmt.Errorf("create documents failed: %s", err)
	}
	return nil
}

func BulkCreateBooks(docs []BookDocument) error {
	ctx := context.Background()
	bulkRequests := client.Bulk()
	for _, doc := range docs {
		req := elastic.NewBulkIndexRequest().Index(bookDocumentIndexName).Id(uuid.NewV4().String()).Doc(doc)
		bulkRequests.Add(req)
	}
	_, err := bulkRequests.Do(ctx)
	if err != nil {
		return fmt.Errorf("bulkrequest failed")
	}
	return nil
}

func SearchBooks(title string) ([]BookDocument, error) {
	var bookType BookDocument

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

	docs := make([]BookDocument, len(searchResult.Hits.Hits))
	for idx, item := range searchResult.Each(reflect.TypeOf(bookType)) {
		if book, ok := item.(BookDocument); ok {
			docs[idx] = book
		}
	}
	return docs, nil
}

func DeleteBooksIndex() error{
	ctx := context.Background()
	_, err := client.DeleteIndex(bookDocumentIndexName).Do(ctx)
	if err != nil {
		return fmt.Errorf("Delete index failed: %s", err)
	}
	return nil
}