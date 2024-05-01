package usecase

import (
	"7Zero4/model"
	"7Zero4/model/dto"
	"7Zero4/repository"
	"7Zero4/utils"
	"time"
)

type UserRegistrationUsecase interface {
	RegisterUser(reqRegistBody dto.RequestRegistBody) (string, error)
}

type userRegistrationUsecase struct {
	userRepo     repository.UserRepository
	mailRepo     repository.MailRepository
	passWordRepo repository.PasswordRepository
	tokenRepo    repository.TokenRepository
}

func (p *userRegistrationUsecase) RegisterUser(reqUserData dto.RequestRegistBody) (string, error) {

	// validation check request body
	errValidate := p.userRepo.ValidateUser(reqUserData.Email, reqUserData.Name, reqUserData.Password, "register")
	if errValidate != nil {
		return "", errValidate
	}

	// validation check email already registered
	exist := p.userRepo.FindByEmail(reqUserData.Email)
	if exist {
		return "", utils.ErrDuplicateValueFound
	}

	// Hash the password
	hashedPasswordStr, errHashed := p.passWordRepo.HashAndSavePassword(reqUserData.Password)
	if errHashed != nil {
		return "", utils.InvalidTypeFormat()
	}

	// Get token auth
	token, tokenErr := p.tokenRepo.CreateTokenV2(reqUserData.Email, 12)
	if tokenErr != nil {
		return "", tokenErr
	}

	// insert to database
	err := p.userRepo.Register(model.User{
		Name:             reqUserData.Name,
		Email:            reqUserData.Email,
		Password:         hashedPasswordStr,
		RegistrationDate: time.Now().Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func NewUserRegistrationUsecase(userRepo repository.UserRepository, mailRepo repository.MailRepository, passWordRepo repository.PasswordRepository, tokenRepo repository.TokenRepository) UserRegistrationUsecase {
	usecase := new(userRegistrationUsecase)
	usecase.userRepo = userRepo
	usecase.mailRepo = mailRepo
	usecase.passWordRepo = passWordRepo
	usecase.tokenRepo = tokenRepo

	return usecase
}
