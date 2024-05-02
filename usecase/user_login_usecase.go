package usecase

import (
	"7Zero4/model/dto"
	"7Zero4/repository"
	"7Zero4/utils"
)

type UserLoginUseCase interface {
	LoginUser(reqLoginBody dto.RequestLoginBody) (string, error)
}

type userLoginUsecase struct {
	tokenRepo    repository.TokenRepository
	userRepo     repository.UserRepository
	passWordRepo repository.PasswordRepository
}

func (u *userLoginUsecase) LoginUser(reqLoginBody dto.RequestLoginBody) (string, error) {

	errValidate := u.userRepo.ValidateUser(reqLoginBody.Email, "", reqLoginBody.Password, "login")
	if errValidate != nil {
		return "", errValidate
	}

	dbPass, errdbPass := u.userRepo.FindPasswordByEmail(reqLoginBody.Email)
	if errdbPass != nil {
		return "", utils.UserNotFoundError()
	}

	errPassword := u.passWordRepo.VerifyPassword([]byte(dbPass.Password), []byte(reqLoginBody.Password))
	if errPassword != nil {
		return "", utils.PasswordWrongError()
	}

	// Get token auth
	token, tokenErr := u.tokenRepo.CreateTokenV2(reqLoginBody.Email, 12)
	if tokenErr != nil {
		return "", tokenErr
	}

	return token, nil
}

func NewUserLoginUsecase(tokenRepo repository.TokenRepository,
	userRepo repository.UserRepository,
	passWordRepo repository.PasswordRepository) UserLoginUseCase {
	usecase := new(userLoginUsecase)
	usecase.tokenRepo = tokenRepo
	usecase.userRepo = userRepo
	usecase.passWordRepo = passWordRepo
	return usecase
}
