package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Qiita v1 API "GET /api/v1/users/:url_name/stocks"
type UserStockAPI struct {
	UserName string
}

const (
	qiitaUserStockURI = "https://qiita.com/api/v1/users/%s/stocks"
)

func NewUserStockAPI(UserName string) *UserStockAPI {
	us := UserStockAPI{UserName: UserName}
	return &us
}

func (us *UserStockAPI) fetch(url string) ([]Article, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Unable to get (%s): %v", url, err)
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("Unable to get (%s): http status %d", url, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Unable to read response body (%s): %v", url, err)
	}

	var articles []Article
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal response (%s): %v\nThe json body is: %s", url, err, string(body))
	}
	return articles, nil
}

func (us *UserStockAPI) Fetch() []Article {
	// Fetch (HTTP Access)
	articles, err := us.fetch(fmt.Sprintf(qiitaUserStockURI, us.UserName))
	if err != nil {
		log.Fatalf("Failed to fetch the user stock data: %v\n", err)
	}
	return articles
}
