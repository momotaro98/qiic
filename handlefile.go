package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const fileName = ".qiic-json"

var filePath = path.Join(os.Getenv("HOME"), fileName)

// SavedArticle is strcut for saving file into local machine.
type SavedArticle struct {
	ID               int64  `json:"id"`
	UUID             string `json:"uuid"`
	User             User   `json:"user"` // User type is at content.go
	Title            string `json:"title"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	CreatedAtInWords string `json:"created_at_in_words"`
	UpdatedAtInWords string `json:"updated_at_in_words"`
	Tags             []Tag  `json:"tags"` // Tag type is at content.go
	StockCount       int    `json:"stock_count"`
	CommentCount     int    `json:"comment_count"`
	URL              string `json:"url"`
	GistURL          string `json:"gist_url"`
	Tweet            bool   `json:"tweet"`
	Private          bool   `json:"private"`
	Stocked          bool   `json:"stocked"`
}

// Save saves the json file into local.
func Save(arts []Article) error {
	savedArticles := make([]SavedArticle, len(arts))
	for i, art := range arts {
		savedArticles[i] = SavedArticle{
			ID:               art.ID,
			UUID:             art.UUID,
			User:             art.User,
			Title:            art.Title,
			CreatedAt:        art.CreatedAt,
			UpdatedAt:        art.UpdatedAt,
			CreatedAtInWords: art.CreatedAtInWords,
			UpdatedAtInWords: art.UpdatedAtInWords,
			Tags:             art.Tags,
			StockCount:       art.StockCount,
			CommentCount:     art.CommentCount,
			URL:              art.URL,
			GistURL:          art.GistURL,
			Tweet:            art.Tweet,
			Private:          art.Private,
			Stocked:          art.Stocked,
		}
	}
	bytes, err := json.Marshal(savedArticles) // encoding
	if err != nil {
		return fmt.Errorf("Unable to marshal savedArticles")
	}
	err = ioutil.WriteFile(filePath, bytes, 0644)
	if err != nil {
		return fmt.Errorf("Unable to write to %s", filePath)
	}
	return nil
}

// Load loads the json file from local.
func Load() ([]Article, error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("Unable to load %s\nGet your articles with executing the following command\n$ qiic u", filePath)
	}
	var articles []Article
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, fmt.Errorf("Unable to unmashal body")
	}
	return articles, nil
}
