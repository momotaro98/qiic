package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	qiitaUserStockURI                 = "https://qiita.com/api/v2/users/%s/stocks?page=%d&per_page=20"
	qiitaGetAuthenticatedUserItemsURI = "https://qiita.com/api/v2/authenticated_user/items?page=%d&per_page=20"
)

func setAuthHeader(req *http.Request, token string) {
	val := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", val)
}

type ReqGetAuthenticatedUserItems struct {
	Token string
	Page  int
}

func GetAuthenticatedUserItems(reqObj *ReqGetAuthenticatedUserItems) ([]*Article, error) {
	url := fmt.Sprintf(qiitaGetAuthenticatedUserItemsURI, reqObj.Page)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	setAuthHeader(req, reqObj.Token)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, fmt.Errorf("response is not 200. res: %+v, body: %+v\n", res, body)
	}

	if err != nil {
		return nil, fmt.Errorf("Unable to read response body (%s): %v", url, err)
	}

	var articlesFromApi []QiitaArticle
	if err := json.Unmarshal(body, &articlesFromApi); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal response (%s): %v", url, err)

	}
	// Convert API data model to qiic domain model
	var ret []*Article
	for _, a := range articlesFromApi {
		var tags []Tag
		for _, t := range a.Tags {
			tag := NewTag(t.Name)
			tags = append(tags, tag)
		}
		article := NewArticle(a.ID, a.Title, tags, a.LikesCount, a.URL)
		ret = append(ret, &article)
	}

	return ret, nil
}

// UserStockAPI is Qiita v2 API resource "GET /api/v2/users/:url_name/stocks"
type UserStockAPI struct {
	UserName string
	Page     int
}

// NewUserStockAPI is a func.
func NewUserStockAPI(UserName string, Page int) *UserStockAPI {
	us := UserStockAPI{UserName: UserName, Page: Page}
	return &us
}

func (us *UserStockAPI) fetch(url string) ([]*Article, error) {
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

	var articlesFromApi []QiitaArticle
	if err := json.Unmarshal(body, &articlesFromApi); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal response (%s): %v", url, err)

	}
	// Convert API data model to qiic domain model
	var ret []*Article
	for _, a := range articlesFromApi {
		var tags []Tag
		for _, t := range a.Tags {
			tag := NewTag(t.Name)
			tags = append(tags, tag)
		}
		article := NewArticle(a.ID, a.Title, tags, a.LikesCount, a.URL)
		ret = append(ret, &article)
	}

	return ret, nil
}

// Fetch (HTTP Access)
func (us *UserStockAPI) Fetch() ([]*Article, error) {
	articles, err := us.fetch(fmt.Sprintf(qiitaUserStockURI, us.UserName, us.Page))
	if err != nil {
		return nil, err
	}
	return articles, nil
}
