package main

import (
	"context"
	"github.com/olivere/elastic/v7"
)

const (
	ES_URL      = "http://10.128.0.2:9200"
	ES_USERNAME = "elastic"
	ES_PASSWORD = "RiV4QOPOj2M4PCAe7ejf"
)

func readFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth(ES_USERNAME, ES_PASSWORD))
	if err != nil {
		return nil, err
	}

	searchResult, err := client.Search().Index(index).Query(query).Pretty(true).Do(context.Background())
	if err != nil {
		return nil, err
	}
	return searchResult, nil
}

func saveToES(i interface{}, index string, id string) error {
	client, err := elastic.NewClient(
		elastic.SetURL(ES_URL),
		elastic.SetBasicAuth(ES_USERNAME, ES_PASSWORD))
	if err != nil {
		return err
	}
	_, err = client.Index().
		Index(index).
		Id(id).
		BodyJson(i).
		Do(context.Background())
	return err
}
