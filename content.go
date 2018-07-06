package main

// Article is a struct of Qiita data model.
type Article struct {
	ID               int64  `json:"id"`
	UUID             string `json:"uuid"`
	User             User   `json:"user"`
	Title            string `json:"title"`
	Body             string `json:"body"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	CreatedAtInWords string `json:"created_at_in_words"`
	UpdatedAtInWords string `json:"updated_at_in_words"`
	Tags             []Tag  `json:"tags"`
	StockCount       int    `json:"stock_count"`
	// StockUsers       []User `json:"stock_users"` // Causion! StockUsers will be string type because json's StockUsers can be empty.
	CommentCount int    `json:"comment_count"`
	URL          string `json:"url"`
	GistURL      string `json:"gist_url"`
	Tweet        bool   `json:"tweet"`
	Private      bool   `json:"private"`
	Stocked      bool   `json:"stocked"`
}

// NewTestArticle is a constructor of Article.
func NewTestArticle(ID int64, User *User, Title string, Tags []Tag) Article {
	UUID := "abcdefghijklmnopqrstuvwxyz"
	a := Article{ID: ID, UUID: UUID,
		User: *User, Title: Title, Tags: Tags}
	return a
}

// NewTestArticles is a constructor for tests.
func NewTestArticles() []Article {
	as := make([]Article, 5)
	as[0] = NewTestArticle(1, NewTestUser(), "title01", NewTestTags("python", "golang"))
	as[1] = NewTestArticle(2, NewTestUser(), "title02", NewTestTags("golang", "oop"))
	as[2] = NewTestArticle(3, NewTestUser(), "title03", NewTestTags("golang", "ddd"))
	as[3] = NewTestArticle(4, NewTestUser(), "title04", NewTestTags("python", "cli"))
	as[4] = NewTestArticle(5, NewTestUser(), "title05", NewTestTags("python", "golang"))
	return as
}

// User is a struct of a user on Qiita.
type User struct {
	ID              int    `json:"id"`
	Following       bool   `json:"following"`
	URLName         string `json:"url_name"`
	ProfileImageURL string `json:"profile_image_url"`
}

// NewTestUser a constructor for tests.
func NewTestUser() *User {
	ID := 100
	Following := false
	URLName := "john_url"
	ProfileImageURL := "https://qiita.com"
	u := User{ID: ID, Following: Following, URLName: URLName, ProfileImageURL: ProfileImageURL}
	return &u
}

// Tag is a struct of Qiita data model.
type Tag struct {
	Name     string   `json:"name"`
	URLName  string   `json:"url_name"`
	IconURL  string   `json:"icon_url"`
	Versions []string `json:"versions"`
}

// NewTestTag is a constructor for tests.
func NewTestTag(Name string) Tag {
	URLName := "https://qiita.com/" + Name
	IconURL := "https://qiita.com/icon/" + Name
	Versions := []string{"1.1.0", "1.2.0"}
	t := Tag{Name: Name, URLName: URLName,
		IconURL: IconURL, Versions: Versions}
	return t
}

// NewTestTags is a constructor for tests.
func NewTestTags(tag1 string, tag2 string) []Tag {
	ts := make([]Tag, 2)
	ts[0] = NewTestTag(tag1)
	ts[1] = NewTestTag(tag2)
	return ts
}
