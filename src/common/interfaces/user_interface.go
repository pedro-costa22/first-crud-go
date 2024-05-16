package interfaces

import (
	"github.com/pedro-costa22/first-crud-go/src/common/request"
	"github.com/pedro-costa22/first-crud-go/src/config/database/entity"
)

type IUserService interface {
	Create(req request.UserCreateRequest) (*entity.UserEntity, error)
	FindByID() error
	FindByEmail() error
	Update() error
	Delete() error
}

type IUserRepository interface {
	Save(user *entity.UserEntity) error
	FindByID() 
	FindByEmail() 
	Update() 
	Delete() 
}
