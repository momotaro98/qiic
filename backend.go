package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tomnomnom/linkheader"
)

const (
	qiitaUserStockURI                 = "https://qiita.com/api/v2/users/%s/stocks?page=%d&per_page=20"
	qiitaGetAuthenticatedUserItemsURI = "https://qiita.com/api/v2/authenticated_user/items?page=%d&per_page=20"
)

func setAuthHeader(req *http.Request, token string) {
	val := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", val)
}

type GetRequest struct {
	Token string
	Page  int
}

type ArticlesGetRequester interface {
	AssembleURL() (url string)
	SetAuthHeader(req *http.Request)
}

type Links struct {
	linkheader.Links
}

func GetArticles(r ArticlesGetRequester) (articles []*Article, links *Links, err error) {
	url := r.AssembleURL()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	//setAuthHeader(req, reqObj.Token)
	r.SetAuthHeader(req)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read response body (%s): %v", url, err)
	}

	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		return nil, nil, fmt.Errorf("response is not 200. res: %+v, body: %+v\n", res, body)
	}

	if h := res.Header["Link"]; len(h) == 1 {
		header := h[0]
		links = &Links{
			linkheader.Parse(header),
		}
		for _, link := range links.Links {
			fmt.Printf("URL: %s; Rel: %s\n", link.URL, link.Rel)
		}
	}

	var articlesFromApi []QiitaArticle
	if err := json.Unmarshal(body, &articlesFromApi); err != nil {
		return nil, nil, fmt.Errorf("unable to unmarshal response (%s): %v", url, err)

	}
	// Convert API data model to qiic domain model
	for _, a := range articlesFromApi {
		var tags []Tag
		for _, t := range a.Tags {
			tag := NewTag(t.Name)
			tags = append(tags, tag)
		}
		article := NewArticle(a.ID, a.Title, tags, a.LikesCount, a.URL)
		articles = append(articles, &article)
	}

	return articles, links, nil
}

type ReqGetAuthenticatedUserItems struct {
	GetRequest
}

func (r *ReqGetAuthenticatedUserItems) AssembleURL() string {
	return fmt.Sprintf(qiitaGetAuthenticatedUserItemsURI, r.Page)
}

func (r *ReqGetAuthenticatedUserItems) SetAuthHeader(req *http.Request) {
	setAuthHeader(req, r.Token)
}

// func CollectAuthenticatedUserItems(reqObj *ReqGetAuthenticatedUserItems) ([]*Article, error) {
// }

// UserStockRequest is Qiita v2 API resource "GET /api/v2/users/:url_name/stocks"
type UserStockRequest struct {
	UserName string
	GetRequest
}

func (r *UserStockRequest) AssembleURL() string {
	return fmt.Sprintf(qiitaUserStockURI, r.UserName, r.Page)
}

func (r *UserStockRequest) SetAuthHeader(req *http.Request) {
	// Do nothing
}
