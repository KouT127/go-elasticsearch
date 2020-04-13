package elasticsearch

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var (
	client *elastic.Client
	err    error
)

const (
	bookDocumentIndexName = "books"
)

type logger struct{}

func (l logger) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v)
}

func InitClient() {
	//logger := logger{}

	client, err = elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false),
		//elastic.SetInfoLog(logger),
		//elastic.SetTraceLog(logger),
		//elastic.SetErrorLog(logger),
	)
	if err != nil {
		panic(err)
	}

	indices := []string{
		bookDocumentIndexName,
	}

	for _, name := range indices {
		ctx := context.Background()
		service := elastic.NewIndicesExistsService(client)
		service.Index([]string{name})
		exists, err := service.Do(ctx)
		if err != nil {
			panic(err)
		}
		if !exists {
			_, err := client.CreateIndex(name).Do(ctx)
			if err != nil {
				panic(err)
			}
		}
	}
}
