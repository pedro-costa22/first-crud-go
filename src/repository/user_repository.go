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

func (r *UserRepository) FindByID(id string) (entity.UserEntity, error) {
	var user entity.UserEntity
	if err := r.DB.Where("id = ?", id).First(&user).Error; err != nil {
        return user, err
    }
    return user, nil
}

func (r *UserRepository) FindByEmail(email string) (entity.UserEntity, error) {
	var user entity.UserEntity
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) Update(id string, updates map[string]interface{}) error {
	if err := r.DB.Model(&entity.UserEntity{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(id string) error {
	if err := r.DB.Where("id = ?", id).Delete(&entity.UserEntity{}).Error; err != nil {
		return err
	}

	return nil
}
