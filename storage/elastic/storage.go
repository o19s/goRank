package elastic

import (
	"context"
	"fmt"
	"goRank/models"
	"reflect"

	es "gopkg.in/olivere/elastic.v5"
)

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"tweet":{
			"properties":{
				"search":{
					"type":"keyword"
				},
				"item":{
					"type":"keyword"
				},
				"timestamp":{
					"type":"date"
				}
			}
		}
	}
}`

type Elastic struct {
	client *es.Client
}

func (e *Elastic) Connect() {
	client, err := es.NewClient(
		es.SetSniff(false),
		es.SetBasicAuth("elastic", "changeme"),
	)
	if err != nil {
		// Handle error
		panic(err)
	}
	e.client = client

}

func (e Elastic) Version() {
	info, code, err := e.client.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
}

func (e Elastic) FindEventsForSearch(search string) (events []models.Event, err error) {
	termQuery := es.NewTermQuery("search", search)
	es.NewMatchAllQuery()
	searchResult, err := e.client.Search().
		Index("gorank").
		Query(termQuery).
		Sort("timestamp", true).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}

	var ttyp models.Event
	for _, item := range searchResult.Each(reflect.TypeOf(ttyp)) {
		if e, ok := item.(models.Event); ok {
			events = append(events, e)
		}
	}
	return
}

func (e Elastic) InitStorage() {
	exists, err := e.client.IndexExists("gorank").Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	if !exists {
		// Create a new index.
		createIndex, err := e.client.CreateIndex("gorank").BodyString(mapping).Do(context.Background())
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

func (e Elastic) Save(event models.Event) {
	_, err := e.client.Index().
		Index("gorank").
		Type("event").
		BodyJson(event).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
}
