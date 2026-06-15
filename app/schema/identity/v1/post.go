package schema

// PostAddRequest struct for adding a new post
type PostAddRequest struct {
	UserID  uint64 `json:"user_id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

// PostUpdateRequest struct for updating a post
type PostUpdateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  int    `json:"status"`
}

// PostDetailResponse struct for post detail response
type PostDetailResponse struct {
	ID         uint64 `json:"id"`
	UserID     uint64 `json:"user_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
	Status     int    `json:"status"`
}

// PostListResponse struct for post list response
type PostListResponse struct {
	Posts     []PostDetailResponse `json:"posts"`
	Total     int64                `json:"total"`
	Page      int                  `json:"page"`
	PageSize  int                  `json:"page_size"`
}