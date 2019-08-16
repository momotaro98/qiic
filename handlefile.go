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
	ID             string `json:"id"`
	Title          string `json:"title"`
	Tags           []Tag  `json:"tags"`
	ReactionsCount int    `json:"reactions_count"`
	URL            string `json:"url"`
}

// Save saves the json file into local.
func Save(arts []Article) error {
	savedArticles := make([]SavedArticle, len(arts))
	for i, art := range arts {
		savedArticles[i] = SavedArticle{
			ID:             art.ID,
			Title:          art.Title,
			Tags:           art.Tags,
			ReactionsCount: art.LikesCount,
			URL:            art.URL,
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
