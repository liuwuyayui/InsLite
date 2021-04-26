package main

import (
	"context"
	"fmt"
	"github.com/magiconair/properties"
	"github.com/olivere/elastic/v7"
)

const (
	POST_INDEX = "post"
	USER_INDEX = "user"
)

func main() {
	p := properties.MustLoadFile("credentials.properties", properties.UTF8)
	ES_URL, _ := p.Get("ES_URL")
	ES_USERNAME, _ := p.Get("ES_USERNAME")
	ES_PASSWORD, _ := p.Get("ES_PASSWORD")
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth(ES_USERNAME, ES_PASSWORD))
	if err != nil {
		panic(err)
	}

	exists, err := client.IndexExists(POST_INDEX).Do(context.Background())
	if err != nil {
		panic(err)
	}

	if !exists {
		mapping := `{
            "mappings": {
                "properties": {
                    "id":       { "type": "keyword" },
                    "user":     { "type": "keyword" },
                    "message":  { "type": "text" },
                    "url":      { "type": "keyword", "index": false },
                    "type":     { "type": "keyword", "index": false }
                }
            }
        }`
		_, err := client.CreateIndex(POST_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}

	if !exists {
		mapping := `{
			"mappings": {
				"properties": {
					"username": {"type": "keyword"},
					"password": {"type": "keyword"},
					"age":      {"type": "long", "index": false},
					"gender":   {"type": "keyword", "index": false}
				}
			}
		}`
		_, err = client.CreateIndex(USER_INDEX).Body(mapping).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Indexes are created.")
}
