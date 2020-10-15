package entity

const (
	// UserStatusDefault new user
	UserStatusDefault int = 0
	// UserStatusActive user active
	UserStatusActive int = 1
	// UserStatusInActive user in-active
	UserStatusInActive int = 2
	// UserConfirmTypeEmail confirm email
	UserConfirmTypeEmail int = 1
	// UserConfirmTypeChangePassword confirm change password
	UserConfirmTypeChangePassword int = 2
)

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

// TableName func get table name
func (u *User) TableName() string {
	return "im_user"
}
