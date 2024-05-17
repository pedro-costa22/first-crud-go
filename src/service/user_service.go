package service

import (
	"errors"
	"fmt"

	"github.com/pedro-costa22/first-crud-go/src/common/interfaces"
	"github.com/pedro-costa22/first-crud-go/src/common/request"
	"github.com/pedro-costa22/first-crud-go/src/config/database/entity"
)

type UserService struct {
	userRepository interfaces.IUserRepository
}

func NewUserService(r interfaces.IUserRepository) *UserService {
	return &UserService{userRepository: r}
}

func (u *UserService) Create(req request.UserCreateRequest) (*entity.UserEntity, error) {
	user, err := entity.NewUser(
		req.Name, 
		req.Email, 
		req.Password, 
		req.Age,
	)

	if err != nil {
		return nil, err
	}

	err = u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) FindByID(id string) (entity.UserEntity, error)  {
	user, err := u.userRepository.FindByID(id)
	fmt.Println(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserService) FindByEmail(email string) (entity.UserEntity, error)  {
	user, err := u.userRepository.FindByEmail(email) 
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *UserService) Update(id string, updates map[string]interface{}) (entity.UserEntity, error)  {
	user, err := u.FindByID(id)
	if err != nil {
		return user, errors.New("User not found")
	}
	
	err = u.userRepository.Update(id, updates)
	if err != nil {
		return user, err
	}

	user, _ = u.FindByID(id)
	return user, nil
}

func (u *UserService) Delete(id string) error {
	_, err := u.FindByID(id)
	if err != nil {
		return errors.New("User not found")
	}

	err = u.userRepository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}