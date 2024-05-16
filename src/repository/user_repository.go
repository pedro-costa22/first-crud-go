package repository

import (
	"github.com/pedro-costa22/first-crud-go/src/config/database/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Save(user *entity.UserEntity) error {
	result := r.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) FindByID()    {}
func (r *UserRepository) FindByEmail() {}
func (r *UserRepository) Update()      {}
func (r *UserRepository) Delete()      {}
