package entity

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pedro-costa22/first-crud-go/src/common/utils"
	"github.com/pedro-costa22/first-crud-go/src/config/rest_err"
	"golang.org/x/crypto/bcrypt"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	JWT_EXPIRES_IN = "JWT_EXPIRES_IN"
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


func (u *UserEntity) GenerateJWT() (string, error) {
	secret := os.Getenv(JWT_SECRET_KEY)
	expiresStr  := os.Getenv(JWT_EXPIRES_IN)

	expires, err := strconv.Atoi(expiresStr)
	if err != nil {
		return "", rest_err.NewInternalServerError("invalid JWT expiration time")
	}

	cliams := jwt.MapClaims{
		"id": u.ID, 
		"name": u.Name,
		"email": u.Email,
		"age": u.Age,
		"exp": time.Now().Add(time.Hour * time.Duration(expires)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	tokenStr, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", rest_err.NewInternalServerError("error signing token")
	}

	return tokenStr, nil
}
