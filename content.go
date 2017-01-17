package main

type Article struct {
	ID               *int64  `json:"id"`
	UUID             *string `json:"uuid"`
	User             User    `json:"user"`
	Title            *string `json:"title"`
	Body             *string `json:"body"`
	CreatedAt        *string `json:"created_at"`
	UpdatedAt        *string `json:"updated_at"`
	CreatedAtInWords *string `json:"created_at_in_words"`
	UpdatedAtInWords *string `json:"updated_at_in_words"`
	Tags             []Tag   `json:"tags"`
	StockCount       *int64  `json:"stock_count"`
	StockUsers       []User  `json:"stock_users"`
	CommentCount     *int64  `json:"comment_count"`
	URL              *string `json:"url"`
	GistURL          *string `json:"gist_url"`
	Tweet            *bool   `json:"tweet"`
	Private          *bool   `json:"private"`
	Stocked          *bool   `json:"stocked"`
}

func NewSampleArticle(ID int64, User User, Title string, Tags []Tag) *Article {
	UUID := "abcdefg"
	a := Article{ID: &ID, UUID: &UUID,
		User: User, Title: &Title, Tags: Tags}
	return &a
}

type User struct {
	Name            *string `json:"name"`
	URLName         *string `json:"url_name"`
	ProfileImageURL *string `json:"profile_image_url"`
}

type Tag struct {
	Name     *string   `json:"name"`
	URLName  *string   `json:"url_name"`
	IconURL  *string   `json:"icon_url"`
	Versions []*string `json:"versions"`
}
