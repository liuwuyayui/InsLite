package main

import (
	"context"
	"github.com/magiconair/properties"
	"github.com/olivere/elastic/v7"
)

func readFromES(query elastic.Query, index string) (*elastic.SearchResult, error) {
	p := properties.MustLoadFile("credentials.properties", properties.UTF8)
	ES_USERNAME, _ := p.Get("ES_USERNAME")
	ES_PASSWORD, _ := p.Get("ES_PASSWORD")
	ES_URL, _ := p.Get("ES_URL")
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
	p := properties.MustLoadFile("credentials.properties", properties.UTF8)
	EsUsername, _ := p.Get("ES_USERNAME")
	EsPassword, _ := p.Get("ES_PASSWORD")
	EsUrl, _ := p.Get("ES_URL")
	client, err := elastic.NewClient(
		elastic.SetURL(EsUrl),
		elastic.SetBasicAuth(EsUsername, EsPassword))
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
