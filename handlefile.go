package qiic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

const fileName = ".qiic-json"

var filePath = path.Join(os.Getenv("HOME"), fileName)

// Save saves the json file into local.
func Save(arts []*Article) error {
	bytes, err := json.Marshal(arts) // encoding
	if err != nil {
		return fmt.Errorf("unable to marshal savedArticles")
	}
	err = ioutil.WriteFile(filePath, bytes, 0644)
	if err != nil {
		return fmt.Errorf("unable to write to %s", filePath)
	}
	return nil
}

// Load loads the json file from local.
func Load() ([]*Article, error) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to load %s\nGet your articles with executing the following command\n$ qiic u", filePath)
	}
	var articles []*Article
	if err := json.Unmarshal(body, &articles); err != nil {
		return nil, fmt.Errorf("unable to unmashal body")
	}
	return articles, nil
}
