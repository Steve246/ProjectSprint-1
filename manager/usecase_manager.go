package manager

import "7Zero4/usecase"

type UsecaseManager interface {
	RegistUsecase() usecase.UserRegistrationUsecase
	TokenUsecase() usecase.TokenUsecase
	LoginUsecase() usecase.UserLoginUseCase
	CatUsecase() usecase.CatUseCase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

func (u *usecaseManager) RegistUsecase() usecase.UserRegistrationUsecase {
	return usecase.NewUserRegistrationUsecase(u.repoManager.UserRepo(), u.repoManager.MailRepo(), u.repoManager.PasswordRepo(), u.repoManager.TokenRepo())

}

func (u *usecaseManager) TokenUsecase() usecase.TokenUsecase {
	return usecase.NewTokenUsecase(u.repoManager.TokenRepo())
}

// func (u *usecaseManager) LoginUsecase() usecase.UserLoginUseCase {
// 	return usecase.NewUserLoginUsecase(u.repoManager.OtpRepo(), u.repoManager.TokenRepo(), u.repoManager.UserRepo(), u.repoManager.MailRepo(), u.repoManager.PasswordRepo())
// }

func (u *usecaseManager) LoginUsecase() usecase.UserLoginUseCase {
	return usecase.NewUserLoginUsecase(u.repoManager.TokenRepo(), u.repoManager.UserRepo(), u.repoManager.MailRepo(), u.repoManager.PasswordRepo())
}

func (u *usecaseManager) CatUsecase() usecase.CatUseCase {
	return usecase.NewCatUsecase(u.repoManager.CatRepo())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
