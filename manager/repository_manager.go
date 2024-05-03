package manager

import "7Zero4/repository"

type RepositoryManager interface {
	PasswordRepo() repository.PasswordRepository
	UserRepo() repository.UserRepository
	// OtpRepo() repository.OtpRepository
	MailRepo() repository.MailRepository
	TokenRepo() repository.TokenRepository
	CatRepo() repository.CatRepository
}

type repositoryManager struct {
	infra Infra
}

func (r *repositoryManager) PasswordRepo() repository.PasswordRepository {
	return repository.NewPasswordRepository(r.infra.SqlDb())
}

func (r *repositoryManager) UserRepo() repository.UserRepository {
	return repository.NewUserRepository(r.infra.SqlDb())
}

func (r *repositoryManager) CatRepo() repository.CatRepository {
	return repository.NewCatRepository(r.infra.SqlDb())
}

// func (r *repositoryManager) OtpRepo() repository.OtpRepository {
// 	return repository.NewOtpRepository(r.infra.RedisClient())
// }

func (r *repositoryManager) MailRepo() repository.MailRepository {
	return repository.NewMailRepository(r.infra.MailConfig())
}

// func (r *repositoryManager) TokenRepo() repository.TokenRepository {
// 	return repository.NewTokenRepository(r.infra.RedisClient(), r.infra.TokenConfig(), r.infra.SqlDb())
// }

func (r *repositoryManager) TokenRepo() repository.TokenRepository {
	return repository.NewTokenRepository(r.infra.TokenConfig(), r.infra.SqlDb())
}

func NewRepositoryManager(infra Infra) RepositoryManager {
	return &repositoryManager{
		infra: infra,
	}
}
