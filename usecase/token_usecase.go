package usecase

import (
	"7Zero4/model/dto"
	"7Zero4/repository"
	"log"

	"strconv"

	"github.com/golang-jwt/jwt"
)

type TokenUsecase interface {
	VerifyAccessToken(tokenString string) (bool, error)
	FetchAccessToken(accessUuid string) (uint, error)
}

type tokenUsecase struct {
	tokenRepo repository.TokenRepository
}

type MyClaims struct {
	jwt.StandardClaims
	AuthToken dto.AuthToken
}

func (t *tokenUsecase) VerifyAccessToken(tokenString string) (bool, error) {
	return t.tokenRepo.VerifyTokenV2(tokenString)
}

func (t *tokenUsecase) FetchAccessToken(accessUuid string) (uint, error) {
	stringId, err := t.tokenRepo.FetchToken(accessUuid)
	if err != nil {
		return 0, err
	}
	log.Println(stringId)
	intId, err := strconv.Atoi(stringId)
	if err != nil {
		return 0, err
	}
	return uint(intId), nil
}

func NewTokenUsecase(tokenRepo repository.TokenRepository) TokenUsecase {
	usecase := new(tokenUsecase)
	usecase.tokenRepo = tokenRepo
	return usecase
}
