package repository

import (
	"7Zero4/model"
	"7Zero4/utils"
	"errors"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindPasswordByEmail(email string) (model.User, error)
	Create(newData interface{}) error
	FindByEmail(email string) bool
	FindBy(selected interface{}, by interface{}) error
	ValidateUser(email string, name string, password string) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) ValidateUser(email string, name string, password string) error {
	if email == "" {
		return utils.ErrEmailNull
	}

	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return utils.ErrInvalidEmail
	}

	// Check if name is not null and length is between 5 and 50
	if name == "" {
		return utils.ErrNameNull
	}
	nameLength := len(strings.TrimSpace(name))
	if nameLength < 5 || nameLength > 50 {
		return utils.ErrInvalidName
	}

	// Check if password is not null and length is between 5 and 15
	if password == "" {
		return utils.ErrPasswordNull
	}
	passwordLength := len(password)
	if passwordLength < 5 || passwordLength > 15 {
		return utils.ErrInvalidPassword
	}

	return nil
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
	if result.RowsAffected == 0 {
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
