package main

import (
	"time"
)

// QiitaArticle is the data model of an article from Qiita API
type QiitaArticle struct {
	RenderedBody   string      `json:"rendered_body"`
	Body           string      `json:"body"`
	Coediting      bool        `json:"coediting"`
	CommentsCount  int         `json:"comments_count"`
	CreatedAt      time.Time   `json:"created_at"`
	Group          interface{} `json:"group"`
	ID             string      `json:"id"`
	LikesCount     int         `json:"likes_count"`
	Private        bool        `json:"private"`
	ReactionsCount int         `json:"reactions_count"`
	Tags           []struct {
		Name     string        `json:"name"`
		Versions []interface{} `json:"versions"`
	} `json:"tags"`
	Title     string    `json:"title"`
	UpdatedAt time.Time `json:"updated_at"`
	URL       string    `json:"url"`
	User      struct {
		Description       string `json:"description"`
		FacebookID        string `json:"facebook_id"`
		FolloweesCount    int    `json:"followees_count"`
		FollowersCount    int    `json:"followers_count"`
		GithubLoginName   string `json:"github_login_name"`
		ID                string `json:"id"`
		ItemsCount        int    `json:"items_count"`
		LinkedinID        string `json:"linkedin_id"`
		Location          string `json:"location"`
		Name              string `json:"name"`
		Organization      string `json:"organization"`
		PermanentID       int    `json:"permanent_id"`
		ProfileImageURL   string `json:"profile_image_url"`
		TeamOnly          bool   `json:"team_only"`
		TwitterScreenName string `json:"twitter_screen_name"`
		WebsiteURL        string `json:"website_url"`
	} `json:"user"`
	PageViewsCount interface{} `json:"page_views_count"`
}

// Article is a struct of qiic article domain model
type Article struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Tags       []Tag  `json:"tags"`
	LikesCount int    `json:"likes_count"`
	URL        string `json:"url"`
}

// NewArticle creates a new Article data type
func NewArticle(id string, title string, tags []Tag, likesCount int, url string) (art Article) {
	art.ID = id
	art.Title = title
	art.Tags = tags
	art.LikesCount = likesCount
	art.URL = url
	return
}

// Tag is a struct of qiic data model.
type Tag struct {
	Name string `json:"name"`
}

// NewTag creates a new Tag data type
func NewTag(name string) (tag Tag) {
	tag.Name = name
	return
}
