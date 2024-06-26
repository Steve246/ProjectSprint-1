package repository

import (
	"7Zero4/model"
	"7Zero4/utils"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindPasswordByEmail(email string) (model.User, error)
	Register(data model.User) error
	FindByEmail(email string) bool
	FindBy(selected interface{}, by interface{}) error
	ValidateUser(email string, name string, password string, user string) error
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) ValidateUser(email string, name string, password string, user string) error {

	if user == "register" {
		if email == "" {
			return utils.ReqBodyNotValidError()
		}

		fmt.Println("ini email --> ", email)

		// emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`

		emailRegex := `(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
		match, _ := regexp.MatchString(emailRegex, email)
		if !match {
			return utils.ReqBodyNotValidError()
		}

		// Check if name is not null and length is between 5 and 50
		if name == "" {
			return utils.ReqBodyNotValidError()
		}
		nameLength := len(strings.TrimSpace(name))

		fmt.Println("ini nameLength --> ", nameLength)

		if nameLength < 5 || nameLength > 50 {
			return utils.ReqBodyNotValidError()
		}

		// Check if password is not null and length is between 5 and 15
		if password == "" {
			return utils.ReqBodyNotValidError()
		}
		passwordLength := len(password)
		if passwordLength < 5 || passwordLength > 15 {
			return utils.ReqBodyNotValidError()
		}

		return nil
	}

	if user == "login" {
		if email == "" {
			return utils.ReqBodyNotValidError()
		}

		emailRegex := `(?i)^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`
		match, _ := regexp.MatchString(emailRegex, email)
		if !match {
			return utils.ReqBodyNotValidError()
		}

		// Check if name is not null and length is between 5 and 50
		// if name == "" {
		// 	return utils.ErrNameNull
		// }
		// nameLength := len(strings.TrimSpace(name))
		// if nameLength < 5 || nameLength > 50 {
		// 	return utils.ErrInvalidName
		// }

		// Check if password is not null and length is between 5 and 15
		if password == "" {
			return utils.ReqBodyNotValidError()
		}
		passwordLength := len(password)
		if passwordLength < 5 || passwordLength > 15 {
			return utils.ReqBodyNotValidError()
		}

		return nil
	}

	return nil

}

func (u *userRepository) FindPasswordByEmail(email string) (model.User, error) {
	var user model.User
	u.db.Raw("SELECT * FROM users WHERE email = ?", email).Scan(&user)
	if (user == model.User{}) {
		return model.User{}, errors.New("User not found")
	}
	return user, nil

}

func (u *userRepository) Register(data model.User) error {
	result := u.db.Exec("INSERT INTO users (created_at,updated_at,name,email,password,registration_date) VALUES (?,?,?,?,?,?)", data.RegistrationDate, data.RegistrationDate, data.Name, data.Email, data.Password, data.RegistrationDate)
	return result.Error
}

func (u *userRepository) FindByEmail(email string) bool {
	result := u.db.Exec("SELECT * FROM users WHERE email = ?", email)
	return result.RowsAffected != 0
}

func (u *userRepository) FindBy(selected interface{}, by interface{}) error {
	result := u.db.Exec("SELECT * FROM users WHERE ?", by).Scan(selected)
	return result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	repo := new(userRepository)
	repo.db = db
	return repo
}
