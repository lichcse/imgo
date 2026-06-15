package entity

const (
	// PostStatusDefault new post
	PostStatusDefault int = 0
	// PostStatusPublished post published
	PostStatusPublished int = 1
	// PostStatusDraft post is draft
	PostStatusDraft int = 2
)

// Post struct post entity
type Post struct {
	ID         uint64 `json:"id"`
	UserID     uint64 `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
	Status     int    `json:"status"`
}

// TableName func get table name
func (p *Post) TableName() string {
	return "im_post"
}