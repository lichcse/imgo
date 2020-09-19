package utils

import "golang.org/x/crypto/bcrypt"

// IMCrypt interface
type IMCrypt interface {
	Hash(str string) (string, error)
	CheckHash(str, hash string) bool
}

type imCrypt struct{}

// NewIMCrypt func
func NewIMCrypt() IMCrypt {
	return &imCrypt{}
}

// Hash func
func (c *imCrypt) Hash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckHash func
func (c *imCrypt) CheckHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
