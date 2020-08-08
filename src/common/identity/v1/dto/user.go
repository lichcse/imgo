package dto

// UserAdd user add input
type UserAdd struct {
	FullName string `json:"full_name" swaggertype:"string" maxLength:"3" maxLength:"50" example:"Lich Truong"`
	Username string `json:"username" swaggertype:"string" format:"^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$" example:"lichtv"`
	Email    string `json:"email" swaggertype:"string" example:"example@imgo.com"`
	Password string `json:"password" swaggertype:"string" example:"W3^&(80)&&^x"`
}

// UserUpdate user update input
type UserUpdate struct {
	FullName string `json:"full_name" swaggertype:"string" maxLength:"3" maxLength:"50" example:"Lich Truong"`
	Username string `json:"username" swaggertype:"string" format:"^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$" example:"lichtv"`
	Email    string `json:"email" swaggertype:"string" example:"example@imgo.com"`
	Password string `json:"password" swaggertype:"string" example:"W3^&(80)&&^x"`
}

// UserInputConfirm user input confirm
type UserInputConfirm struct {
	Code string `json:"code"`
	Type string `json:"type"`
}

// UserInputChangePassword user input change password
type UserInputChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

// UserResponse user response
type UserResponse struct {
	ID         uint64 `json:"id" swaggertype:"integer" example:"1"`
	FullName   string `json:"full_name" swaggertype:"string" example:"Lich Truong"`
	Username   string `json:"username" swaggertype:"string" example:"lichtv"`
	Email      string `json:"email" swaggertype:"string" example:"example@imgo.com"`
	CreatedAt  string `json:"created_at" swaggertype:"string" example:"1991-02-13 10:10:10"`
	ModifiedAt string `json:"modified_at" swaggertype:"string" example:"2020-07-15 10:10:10"`
	Status     int    `json:"status" swaggertype:"integer" example:"1"`
}
