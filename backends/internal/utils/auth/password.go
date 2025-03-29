package auth

import "golang.org/x/crypto/bcrypt"

type PasswordHasher struct{}

func (p *PasswordHasher) Hash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hash), err
}

func (p *PasswordHasher) Validate(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}
