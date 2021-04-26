package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
)

const (
	USER_INDEX = "user"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Age      int64  `json:"age"`
	Gender   string `json:"gender"`
}

func checkUser(username, password string) (bool, error) {
	query := elastic.NewBoolQuery()
	query.Must(elastic.NewTermQuery("username", username))
	query.Must(elastic.NewTermQuery("password", password))
	searchResult, err := readFromES(query, USER_INDEX)
	if err != nil {
		return false, err
	}

	return searchResult.TotalHits() > 0, nil
}

func addUser(user *User) (bool, error) {
	query := elastic.NewTermQuery("username", user.Username)
	searchResult, err := readFromES(query, USER_INDEX)
	if err != nil {
		return false, err
	}
	if searchResult.TotalHits() > 0 {
		return false, err
	}

	err = saveToES(user, USER_INDEX, user.Username)
	if err != nil {
		return false, err
	}
	fmt.Printf("User is added: %s\n", user.Username)
	return true, nil
}
