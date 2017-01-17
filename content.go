package main

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
	StockCount       int64  `json:"stock_count"`
	StockUsers       []User `json:"stock_users"`
	CommentCount     int64  `json:"comment_count"`
	URL              string `json:"url"`
	GistURL          string `json:"gist_url"`
	Tweet            bool   `json:"tweet"`
	Private          bool   `json:"private"`
	Stocked          bool   `json:"stocked"`
}

func NewTestArticles(ID int64, User *User, Title string, Tags []Tag) *Article {
	UUID := "abcdefghijklmnopqrstuvwxyz"
	a := Article{ID: ID, UUID: UUID,
		User: *User, Title: Title, Tags: Tags}
	return &a
}

func NewSampleArticles() []*Article {
	as := make([]*Article, 5)
	as[0] = NewTestArticles(1, NewTestUser(), "title01", NewTestTags("python", "golang"))
	as[1] = NewTestArticles(2, NewTestUser(), "title02", NewTestTags("golang", "oop"))
	as[2] = NewTestArticles(3, NewTestUser(), "title03", NewTestTags("golang", "ddd"))
	as[3] = NewTestArticles(4, NewTestUser(), "title04", NewTestTags("python", "cli"))
	as[4] = NewTestArticles(5, NewTestUser(), "title05", NewTestTags("python", "golang"))
	return as
}

type User struct {
	Name            string `json:"name"`
	URLName         string `json:"url_name"`
	ProfileImageURL string `json:"profile_image_url"`
}

func NewTestUser() *User {
	Name := "john"
	URLName := "john_url"
	ProfileImageURL := "https://qiita.com"
	u := User{Name: Name, URLName: URLName, ProfileImageURL: ProfileImageURL}
	return &u
}

type Tag struct {
	Name     string   `json:"name"`
	URLName  string   `json:"url_name"`
	IconURL  string   `json:"icon_url"`
	Versions []string `json:"versions"`
}

func NewTestTag(Name string) Tag {
	URLName := "https://qiita.com/" + Name
	IconURL := "https://qiita.com/icon/" + Name
	Versions := []string{"1.1.0", "1.2.0"}
	t := Tag{Name: Name, URLName: URLName,
		IconURL: IconURL, Versions: Versions}
	return t
}

func NewTestTags(tag1 string, tag2 string) []Tag {
	ts := make([]Tag, 2)
	ts[0] = NewTestTag(tag1)
	ts[1] = NewTestTag(tag2)
	return ts
}
