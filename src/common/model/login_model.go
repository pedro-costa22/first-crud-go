package model

import "golang.org/x/crypto/bcrypt"

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserLogin(email, password string) *UserLogin {
	return &UserLogin{
		Email:    email,
		Password: password,
	}
}

func (u *UserLogin) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(u.Password))
	return err == nil
}