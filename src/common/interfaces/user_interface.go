package interfaces

import (
	"github.com/pedro-costa22/first-crud-go/src/common/request"
	"github.com/pedro-costa22/first-crud-go/src/config/database/entity"
)

type IUserService interface {
	Create(req request.UserCreateRequest) (*entity.UserEntity, error)
	FindByID(id string) (entity.UserEntity, error) 
	FindByEmail(email string) (entity.UserEntity, error) 
	Update(id string, updates map[string]interface{}) (entity.UserEntity, error) 
	Delete(id string) error
}

type IUserRepository interface {
	Save(user *entity.UserEntity) error
	FindByID(id string) (entity.UserEntity, error) 
	FindByEmail(email string) (entity.UserEntity, error) 
	Update(id string, updates map[string]interface{}) error
	Delete(id string) error
}
