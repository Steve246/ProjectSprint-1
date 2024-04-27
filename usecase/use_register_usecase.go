package usecase

import (
	"7Zero4/model"
	"7Zero4/model/dto"
	"7Zero4/repository"
	"7Zero4/utils"
	"fmt"
	"log"

	"time"
)

type UserRegistrationUsecase interface {
	RequestRegist(reqRegistBody dto.RequestRegistBody) error
	VerifyRegist(verifyRegistBody dto.VerifyRegistBody) error
}

type userRegistrationUsecase struct {
	otpRepo      repository.OtpRepository
	userRepo     repository.UserRepository
	mailRepo     repository.MailRepository
	passWordRepo repository.PasswordRepository
}

func (p *userRegistrationUsecase) RequestRegist(reqRegistBody dto.RequestRegistBody) error {
	// Cek apakah email sudah terdaftar
	found := p.userRepo.FindByEmail(reqRegistBody.Email)

	if found {
		return utils.DataDuplicateError()
	}

	// Membuat OTP dan menyimpannya ke redis menggunakan email sebagai key
	newOtp := p.otpRepo.CreateOtp()
	log.Println(newOtp)
	registReqString, err := utils.ToJsonString(reqRegistBody)
	if err != nil {
		return err
	}
	// INI NGATUR EXPIRED TIME, BISA DI SETTING MAU HABIS BERAPA DETIK
	// OTP INI SIMPEN DI LOCALSTORAGE di FE
	if err := p.otpRepo.StoreOtp(newOtp, registReqString, 180*time.Second); err != nil {
		return err
	}
	go p.mailRepo.SendMail(reqRegistBody.Email, fmt.Sprintf("Hello This Is Spendy Apps, This is Ur Registration OTP <b>%s</b>", newOtp))
	if err != nil {
		return err
	}
	return nil
}

func (p *userRegistrationUsecase) VerifyRegist(verifyRegistBody dto.VerifyRegistBody) error {
	savedRegistBody, err := p.otpRepo.FetchOtp(verifyRegistBody.Otp)
	fmt.Println(savedRegistBody)
	if err != nil {
		return utils.WrongOtpError()
	}
	var registBody dto.RequestRegistBody
	err = utils.FromJsonString(savedRegistBody, &registBody)
	if err != nil {
		return err
	}
	fmt.Println(registBody)

	// hashed password

	// Hash the password

	hashedPasswordStr, errHashed := p.passWordRepo.HashAndSavePassword(registBody.Password)

	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registBody.Password), bcrypt.DefaultCost)
	if errHashed != nil {
		return utils.InvalidTypeFormat()
	}

	// Convert hashedPassword from []byte to string
	// hashedPasswordStr := string(hashedPassword)

	// bikin time baru

	currentTime := time.Now()
	timeString := currentTime.Format("2006-01-02 15:04:05")

	// disini kita akan bikin user baru
	newUser := model.User{
		Name:             registBody.Name,
		Email:            registBody.Email,
		Password:         hashedPasswordStr,
		RegistrationDate: timeString,
		// kita gak bikin balance, karena itu harus di create sendiri sama user
	}

	if err := p.userRepo.Create(&newUser); err != nil {
		return err
	}

	return nil
}

func NewUserRegistrationUsecase(otpRepo repository.OtpRepository, userRepo repository.UserRepository, mailRepo repository.MailRepository, passWordRepo repository.PasswordRepository) UserRegistrationUsecase {
	usecase := new(userRegistrationUsecase)
	usecase.otpRepo = otpRepo
	usecase.userRepo = userRepo
	usecase.mailRepo = mailRepo
	usecase.passWordRepo = passWordRepo

	return usecase
}
