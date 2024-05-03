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

	return "", nil

}

// func (u *userLoginUsecase) RequestLogin(rqLoginBody dto.RequestLoginBody) error {

// 	errEmail := u.userRepo.FindByEmail(rqLoginBody.Email)
// 	// equivalent to errEmail != true
// 	if !errEmail {
// 		return utils.DataDuplicateError()
// 	}

// 	dbPass, errdbPass := u.userRepo.FindPasswordByEmail(rqLoginBody.Email)

// 	if errdbPass != nil {
// 		return utils.DataNotFoundError()
// 	}

// 	errPassword := u.passWordRepo.VerifyPassword([]byte(dbPass.Password), []byte(rqLoginBody.Password))

// 	if errPassword != nil {
// 		return utils.PasswordNotMatch()
// 	}

// 	// Membuat OTP dan menyimpannya ke redis menggunakan email sebagai key
// 	newOtp := u.otpRepo.CreateOtp()
// 	if err := u.otpRepo.StoreOtp(newOtp, rqLoginBody.Email, 300*time.Second); err != nil {
// 		return err
// 	}

// 	log.Println(newOtp)
// 	go u.mailRepo.SendMail(rqLoginBody.Email, fmt.Sprintf("Hello This Is Spendy Apps, This is Ur Login OTP <b>%s</b>", newOtp))

// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return nil
// }

// func (u *userLoginUsecase) VerifyLoginOtp(reLoginBody dto.VerifyLoginBody) (dto.VerifyLoginBodyResponse, error) {

// 	// fetch otp
// 	savedEmail, err := u.otpRepo.FetchOtp(reLoginBody.Otp)

// 	if err != nil {
// 		log.Println(err)
// 		return dto.VerifyLoginBodyResponse{}, utils.WrongOtpError()
// 	}

// 	// cek apakah user udh register, atau blm, kalau blm error
// 	var selected model.User
// 	err = u.userRepo.FindBy(&selected, model.User{Email: savedEmail})
// 	if err != nil {
// 		return dto.VerifyLoginBodyResponse{}, utils.DataNotFoundError()
// 	}

// 	// some initialization for generate token
// 	accessUuid := uuid.New().String()
// 	duration := 30 * 24 * time.Hour
// 	now := time.Now().UTC()
// 	end := now.Add(duration)

// 	newToken, err := u.tokenRepo.CreateToken(func(appName string) jwt.Claims {
// 		return MyClaims{
// 			StandardClaims: jwt.StandardClaims{
// 				Issuer:    appName,
// 				IssuedAt:  now.Unix(),
// 				ExpiresAt: end.Unix(),
// 			},
// 			AuthToken: dto.AuthToken{
// 				AccessUuid:   accessUuid,
// 				UserID:       selected.ID,
// 				UserEmail:    selected.Email,
// 				UserPassword: selected.Password,
// 				UserName:     selected.Name,
// 			},
// 		}
// 	})
// 	log.Println(newToken)
// 	if err != nil {
// 		return dto.VerifyLoginBodyResponse{}, err
// 	}
// 	if err := u.tokenRepo.StoreToken(accessUuid, selected.ID, duration); err != nil {
// 		return dto.VerifyLoginBodyResponse{}, err
// 	}

// 	return dto.VerifyLoginBodyResponse{Token: newToken}, nil

// 	// return dto.VerifyLoginBodyResponse{}, nil
// }

// func NewUserLoginUsecase(otpRepo repository.OtpRepository,
// 	tokenRepo repository.TokenRepository,
// 	userRepo repository.UserRepository,
// 	mailRepo repository.MailRepository, passWordRepo repository.PasswordRepository) UserLoginUseCase {
// 	usecase := new(userLoginUsecase)
// 	usecase.otpRepo = otpRepo
// 	usecase.tokenRepo = tokenRepo
// 	usecase.userRepo = userRepo
// 	usecase.mailRepo = mailRepo
// 	usecase.passWordRepo = passWordRepo
// 	return usecase
// }

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
