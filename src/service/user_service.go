package service

import (
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

	// TO DO: VALIDAR SE EMAIL JÁ EXISTE, CASO SIM RETORNAR ERROR, CASO NÃO REALIZAR A CRIAÇÃO
	err = u.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserService) FindByID() error {
	return nil
}

func (u *UserService) FindByEmail() error {
	return nil
}

func (u *UserService) Update() error {
	return nil
}

func (u *UserService) Delete() error {
	return nil
}