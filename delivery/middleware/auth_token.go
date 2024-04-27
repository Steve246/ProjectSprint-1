package middleware

import (
	"7Zero4/usecase"
	"errors"
	"fmt"
	"log"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequiredToken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	tokenUsecase usecase.TokenUsecase
}

func checkBypassAPI(c *gin.Context) bool {
	bypassAPI := []string{ // store API that dont need auth bearer
		"/api/login",
		"/api/register",
		"/api/login/otp",
		"/api/register/otp",
	}

	for _, v := range bypassAPI {
		if c.Request.URL.Path == v {
			return true
		}
	}
	return false
}

func (a *authTokenMiddleware) RequiredToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if checkBypassAPI(c) {
			c.Next()
		} else {
			h := authHeader{}
			if err := c.ShouldBindHeader(&h); err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthrorized",
				})
				c.Abort()
				return
			}
			tokenString := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
			fmt.Println("token", tokenString)
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthrorized",
				})
				c.Abort()
				return
			}

			authToken, err := a.tokenUsecase.VerifyAccessToken(tokenString)
			// ini := errors.New("tokenExpired")
			if err != nil {
				fmt.Println(err)
				if errors.Is(err, err) {
					c.JSON(http.StatusUnauthorized, gin.H{
						"message": "Token Expired",
					})
					c.Abort()
					return
				}
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthrorized",
				})
				c.Abort()
				return
			}
			userId, err := a.tokenUsecase.FetchAccessToken(authToken.AccessUuid)
			log.Println(userId, authToken.UserID, err)
			if userId != authToken.UserID || err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "Unauthrorized2",
				})
				c.Abort()
				return
			} else {
				c.Set("authToken", *authToken)
				c.Next()
			}
		}
	}
}

func NewAuthTokenMiddleware(tokenUsecase usecase.TokenUsecase) AuthTokenMiddleware {
	middleware := new(authTokenMiddleware)
	middleware.tokenUsecase = tokenUsecase
	return middleware
}
