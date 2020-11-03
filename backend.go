package qiic

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"golang.org/x/sync/errgroup"

	"github.com/tomnomnom/linkheader"
)

type contextKey string

const (
	reqPerPage = 30

	userNameContextKey contextKey = "userNameKey"
	tokenContextKey    contextKey = "tokenKey"

	qiitaGetAuthenticatedUserItemsURI = "https://qiita.com/api/v2/authenticated_user/items?page=%d&per_page=%d"
	qiitaGetUserItemsURI              = "https://qiita.com/api/v2/users/%s/items?page=%d&per_page=%d"
	qiitaUserStockURI                 = "https://qiita.com/api/v2/users/%s/stocks?page=%d&per_page=%d"
)

func SetUserName(parents context.Context, val string) context.Context {
	return context.WithValue(parents, userNameContextKey, val)
}

func GetUserName(ctx context.Context) (string, error) {
	v := ctx.Value(userNameContextKey)

	user, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("user name not found")
	}

	return user, nil
}

func SetToken(parents context.Context, val string) context.Context {
	return context.WithValue(parents, tokenContextKey, val)
}

func GetToken(ctx context.Context) (string, error) {
	v := ctx.Value(tokenContextKey)

	token, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("token not found")
	}

	return token, nil
}

func setAuthHeader(ctx context.Context, req *http.Request) {
	token, err := GetToken(ctx)
	if err != nil {
		return
	}
	val := fmt.Sprintf("Bearer %s", token)
	req.Header.Set("Authorization", val)
}

type GetRequest struct {
	Page int
}

type ArticlesGetRequester interface {
	AssembleGetRequest(ctx context.Context) (req *http.Request)
	Replicate(ctx context.Context, request GetRequest) ArticlesGetRequester
}

type Links struct {
	linkheader.Links
}

func GetArticles(ctx context.Context, r ArticlesGetRequester) (articles []*Article, links *Links, err error) {
	req := r.AssembleGetRequest(ctx)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to read response body (%+v): %v", req, err)
	}

	if res.StatusCode != 200 {
		return nil, nil, fmt.Errorf("Qiita API response is not 200. \nreq: %+v, \nres: %+v, \nbody: %s\n", req, res, string(body))
	}

	if h := res.Header["Link"]; len(h) == 1 {
		header := h[0]
		links = &Links{
			linkheader.Parse(header),
		}
	}

	var articlesFromApi []QiitaArticle
	if err := json.Unmarshal(body, &articlesFromApi); err != nil {
		return nil, nil, fmt.Errorf("unable to unmarshal response (%+v): %v", req, err)
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

func (r *ReqGetAuthenticatedUserItems) AssembleGetRequest(ctx context.Context) *http.Request {
	aURL := fmt.Sprintf(qiitaGetAuthenticatedUserItemsURI, r.Page, reqPerPage)
	req, _ := http.NewRequest(http.MethodGet, aURL, nil)
	setAuthHeader(ctx, req)
	return req
}

func (r *ReqGetAuthenticatedUserItems) Replicate(ctx context.Context, request GetRequest) ArticlesGetRequester {
	return &ReqGetAuthenticatedUserItems{
		GetRequest: request,
	}
}

type ReqGetUserItems struct {
	GetRequest
}

func (r *ReqGetUserItems) AssembleGetRequest(ctx context.Context) *http.Request {
	userName, err := GetUserName(ctx)
	if err != nil {
		panic(err)
	}
	aURL := fmt.Sprintf(qiitaGetUserItemsURI, userName, r.Page, reqPerPage)
	req, _ := http.NewRequest(http.MethodGet, aURL, nil)
	return req
}

func (r *ReqGetUserItems) Replicate(ctx context.Context, request GetRequest) ArticlesGetRequester {
	return &ReqGetUserItems{
		GetRequest: request,
	}
}

func parseURLPage(tURL string) (page int, err error) {
	u, err := url.Parse(tURL)
	if err != nil {
		return
	}
	for key, values := range u.Query() {
		if key == "page" {
			page, err = strconv.Atoi(values[0])
			if err != nil {
				return
			}
		}
	}
	return
}

func CollectUserItems(ctx context.Context, reqObj ArticlesGetRequester) ([]*Article, error) {
	// Get 1st page articles
	retArticles, links, err := GetArticles(ctx, reqObj)
	if err != nil {
		return nil, err
	}

	gen := func(links *Links) []ArticlesGetRequester {
		var lastPage int
		for _, link := range links.Links {
			if link.Rel == "last" {
				page, err := parseURLPage(link.URL)
				if err != nil {
					panic(err)
				}
				lastPage = page
			}
		}
		reqs := make([]ArticlesGetRequester, lastPage-1)
		for i := 2; i <= lastPage; i++ {
			reqs[i-2] = reqObj.Replicate(ctx, GetRequest{Page: i})
		}
		return reqs
	}

	reqObjects := gen(links)
	var mu sync.Mutex
	g, ctx := errgroup.WithContext(ctx)
	for _, r := range reqObjects {
		r := r
		g.Go(func() error {
			arts, _, err := GetArticles(ctx, r)
			if err != nil {
				return err
			}
			mu.Lock()
			defer mu.Unlock()
			retArticles = append(retArticles, arts...)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return retArticles, nil
}

// UserStockRequest is Qiita v2 API resource "GET /api/v2/users/:url_name/stocks"
type UserStockRequest struct {
	UserName string
	GetRequest
}

func (r *UserStockRequest) AssembleGetRequest(ctx context.Context) *http.Request {
	aURL := fmt.Sprintf(qiitaUserStockURI, r.UserName, r.Page, reqPerPage)
	req, _ := http.NewRequest(http.MethodGet, aURL, nil)
	return req
}

func (r *UserStockRequest) Replicate(ctx context.Context, request GetRequest) ArticlesGetRequester {
	userName, _ := GetUserName(ctx)
	return &UserStockRequest{
		UserName:   userName,
		GetRequest: request,
	}
}
