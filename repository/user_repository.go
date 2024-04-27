package repository

import (
	"7Zero4/model"
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindPasswordByEmail(email string) (model.User, error)
	Create(newData interface{}) error
	FindByEmail(email string) bool
	FindBy(selected interface{}, by interface{}) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) FindPasswordByEmail(email string) (model.User, error) {
	var user model.User
	result := u.db.Where("email=?", email).Find(&user)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, nil
		} else {
			return user, err
		}
	}

	return user, nil

}

func (u *userRepository) Create(newData interface{}) error {
	result := u.db.Create(newData)
	return result.Error
}

func (u *userRepository) FindByEmail(email string) bool {
	var user model.User
	result := u.db.First(&user, "email = ?", email).Error
	if errors.Is(result, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (u *userRepository) FindBy(selected interface{}, by interface{}) error {
	result := u.db.Where(by).First(selected)
	return result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	repo := new(userRepository)
	repo.db = db
	return repo
}
