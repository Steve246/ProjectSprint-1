package controller

import (
	"7Zero4/delivery/api"
	"7Zero4/model/dto"
	"7Zero4/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router    *gin.RouterGroup
	routerDev *gin.RouterGroup
	ucLogin   usecase.UserLoginUseCase
	ucRegist  usecase.UserRegistrationUsecase
	// ucRegist  usecase.PartnerRegistrationUsecase
	// ucGetUser usecase.GetDataUserUsecase
	api.BaseApi
}

// FIXME: Regis sama Login gak pake OTP

// TODO: tambain error code

func (u *UserController) userRegister(c *gin.Context) {
	var bodyRequest dto.RequestRegistBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		// fmt.Printf("ini isi bodyRequest --> %s", bodyRequest.Email)
		err := u.ucRegist.RegistUser(bodyRequest)
		fmt.Print(err)
		if err != nil {
			u.Failed(c, err)
			return
		}
		u.Success(c, nil)
	}
}

// func (u *UserController) requestRegist(c *gin.Context) {
// 	var bodyRequest dto.RequestRegistBody

// 	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
// 		fmt.Print("Masuk sini")
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 		})
// 	} else {
// 		// fmt.Printf("ini isi bodyRequest --> %s", bodyRequest.Email)
// 		err := u.ucRegist.RequestRegist(bodyRequest)
// 		if err != nil {
// 			u.Failed(c, err)
// 			return
// 		}
// 		u.Success(c, nil)
// 	}
// }

// func (u *UserController) verifyRegist(c *gin.Context) {
// 	var verifOtpBody dto.VerifyRegistBody

// 	if err := u.ParseRequestBody(c, &verifOtpBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 		})
// 	} else {
// 		err := u.ucRegist.VerifyRegist(verifOtpBody)
// 		if err != nil {
// 			u.Failed(c, err)
// 			return
// 		}
// 		u.Success(c, nil)
// 	}
// }

func (u *UserController) requestLogin(c *gin.Context) {
	var bodyRequest dto.RequestLoginBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := u.ucLogin.RequestLogin(bodyRequest)
		if err != nil {
			u.Failed(c, err)
			return
		}
		u.Success(c, nil)
	}
}

func (u *UserController) verifyLogin(c *gin.Context) {
	var verifOtpBody dto.VerifyLoginBody

	if err := u.ParseRequestBody(c, &verifOtpBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		bodyResponse, err := u.ucLogin.VerifyLoginOtp(verifOtpBody)
		if err != nil {
			u.Failed(c, err)
			return
		}
		u.Success(c, bodyResponse)
	}
}

func NewUserController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucLogin usecase.UserLoginUseCase, ucRegist usecase.UserRegistrationUsecase) *UserController {
	controller := UserController{
		router:    router,
		routerDev: routerDev,
		ucLogin:   ucLogin,

		ucRegist: ucRegist,

		BaseApi: api.BaseApi{},
	}

	// USER REGISTER AND LOGIN
	// router.POST("/register/otp", controller.verifyRegist)
	// router.POST("/register", controller.requestRegist)

	router.POST("/v1/user/register", controller.userRegister)

	router.POST("/login", controller.requestLogin)
	router.POST("/login/otp", controller.verifyLogin)

	return &controller
}
