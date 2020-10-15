package utils

import "golang.org/x/crypto/bcrypt"

// IMCrypt interface of crypt object
type IMCrypt interface {
	Hash(str string) (string, error)
	CheckHash(str, hash string) bool
}

type imCrypt struct{}

// NewIMCrypt func new crypt object
func NewIMCrypt() IMCrypt {
	return &imCrypt{}
}

// Hash func generate hash string
func (c *imCrypt) Hash(str string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckHash func check hash string
func (c *imCrypt) CheckHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
