package entity

// User struct user entity
type User struct {
	ID         uint64 `json:"id"`
	FullName   string `json:"full_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	CreatedAt  string `json:"created_at"`
	ModifiedAt string `json:"modified_at"`
	Status     int    `json:"status"`
}

// TableName func
func (u *User) TableName() string {
	return "im_user"
}
