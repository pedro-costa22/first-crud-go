package entity

import (
	"github.com/pedro-costa22/first-crud-go/src/common/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserEntity struct {
	ID       utils.ID `json:"id"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Name     string   `json:"name"`
	Age      int8     `json:"age"`
}

func NewUser(name, email, password string, age int8) (*UserEntity, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &UserEntity{
		ID:       utils.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
		Age: age,
	}, nil
}

func (u *UserEntity) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}