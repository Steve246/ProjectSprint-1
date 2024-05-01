package usecase

import (
	"7Zero4/model/dto"
	"7Zero4/repository"
	"encoding/json"
	"fmt"
	"log"

	"strconv"

	"github.com/golang-jwt/jwt"
)

type TokenUsecase interface {
	VerifyAccessToken(tokenString string) (*dto.AuthToken, error)
	FetchAccessToken(accessUuid string) (uint, error)
}

type tokenUsecase struct {
	tokenRepo repository.TokenRepository
}

type MyClaims struct {
	jwt.StandardClaims
	AuthToken dto.AuthToken
}

func (t *tokenUsecase) VerifyAccessToken(tokenString string) (*dto.AuthToken, error) {
	claims, err := t.tokenRepo.VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}
	authTokenMap := claims["AuthToken"].(map[string]interface{})
	jsonString, _ := json.Marshal(authTokenMap)
	var authToken dto.AuthToken
	json.Unmarshal(jsonString, &authToken)
	fmt.Println("ini isian token decode --> ", authToken)
	return &authToken, err
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
