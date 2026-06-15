package validation

import (
	"errors"
	schema "imgo/app/schema/identity/v1"
	"imgo/app/utils"
)

// PostValidation interface of post validation object
type PostValidation interface {
	Add(postAdd *schema.PostAddRequest) error
	Update(postUpdate *schema.PostUpdateRequest) error
}

type postValidation struct {
	validation utils.IMValidation
}

// NewPostValidation func new post validation object
func NewPostValidation() PostValidation {
	return &postValidation{validation: utils.NewIMValidation()}
}

// Add func validate data add
func (p *postValidation) Add(postAddRequest *schema.PostAddRequest) error {

	if postAddRequest.UserID == 0 {
		return errors.New("post_invalid_user_id")
	}

	if len(postAddRequest.Title) < 3 || len(postAddRequest.Title) > 200 {
		return errors.New("post_invalid_title")
	}

	if len(postAddRequest.Content) == 0 {
		return errors.New("post_invalid_content")
	}

	return nil
}

// Update func validate data update
func (p *postValidation) Update(postUpdateRequest *schema.PostUpdateRequest) error {

	if len(postUpdateRequest.Title) > 0 && (len(postUpdateRequest.Title) < 3 || len(postUpdateRequest.Title) > 200) {
		return errors.New("post_invalid_title")
	}

	return nil
}