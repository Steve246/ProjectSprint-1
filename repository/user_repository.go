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
	result := u.db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return user, nil
		} else {
			return user, result.Error
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
	result := u.db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func (u *userRepository) FindBy(selected interface{}, by interface{}) error {
	result := u.db.Raw("SELECT * FROM users WHERE ?", by).Scan(selected)
	return result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	repo := new(userRepository)
	repo.db = db
	return repo
}
