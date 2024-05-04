package usecase

import (
	"7Zero4/model/dto"
	"7Zero4/repository"
	"7Zero4/utils"
)

type UserLoginUseCase interface {
	LoginUser(reqLoginBody dto.RequestLoginBody) (dto.SuccessLoginBody, error)
}

type userLoginUsecase struct {
	tokenRepo    repository.TokenRepository
	userRepo     repository.UserRepository
	passWordRepo repository.PasswordRepository
}

func (u *userLoginUsecase) LoginUser(reqLoginBody dto.RequestLoginBody) (dto.SuccessLoginBody, error) {

	var successData dto.SuccessLoginBody

	errValidate := u.userRepo.ValidateUser(reqLoginBody.Email, "", reqLoginBody.Password, "login")
	if errValidate != nil {
		return successData, errValidate
	}

	dbPass, errdbPass := u.userRepo.FindPasswordByEmail(reqLoginBody.Email)
	if errdbPass != nil {
		return successData, utils.UserNotFoundError()
	}

	errPassword := u.passWordRepo.VerifyPassword([]byte(dbPass.Password), []byte(reqLoginBody.Password))
	if errPassword != nil {
		return successData, utils.PasswordWrongError()
	}

	// Get token auth
	token, tokenErr := u.tokenRepo.CreateTokenV2(reqLoginBody.Email, 12)
	if tokenErr != nil {
		return successData, tokenErr
	}

	// Populate the success data struct
	successData = dto.SuccessLoginBody{
		Email:       dbPass.Email,
		Password:    dbPass.Password,
		Name:        dbPass.Name, // You can replace this with the actual name you want to return
		AccessToken: token,
	}

	return successData, nil
}

func NewUserLoginUsecase(
	tokenRepo repository.TokenRepository,
	userRepo repository.UserRepository,
	passWordRepo repository.PasswordRepository) UserLoginUseCase {
	usecase := new(userLoginUsecase)

	usecase.tokenRepo = tokenRepo
	usecase.userRepo = userRepo
	usecase.passWordRepo = passWordRepo
	return usecase
}
