package view

import (
	"github.com/pedro-costa22/first-crud-go/src/common/response"
	"github.com/pedro-costa22/first-crud-go/src/config/database/entity"
)

func CreateUserView(user entity.UserEntity) response.UserCreateResponse {
	return response.UserCreateResponse{
		ID: user.ID.String(),
		Email: user.Email,
		Name: user.Name,
		Age: user.Age,
	}
}