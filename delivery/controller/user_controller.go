package controller

import (
	"7Zero4/delivery/api"
	"7Zero4/model/dto"
	"7Zero4/usecase"
	"7Zero4/utils"
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

// TODO: tambain cara login

func (u *UserController) userLogin(c *gin.Context) {
	var bodyRequest dto.RequestLoginBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ServerError())

	} else {
		err := u.ucLogin.LoginUser(bodyRequest)

		if err != nil {
			u.Failed(c, err)
			return
		}

		successData := dto.SuccessLoginBody{
			Email:    bodyRequest.Email,
			Password: bodyRequest.Password,
			// TODO: ini bikin token logic, sementara masih dummy
			AccessToken: "qwertyuiopasdfghjklzxcvbnm", // This should be the actual access token
		}

		u.Success(c, successData)
	}
}

func (u *UserController) userRegister(c *gin.Context) {
	var bodyRequest dto.RequestRegistBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {

		u.Failed(c, utils.ServerError())

	} else {

		err := u.ucRegist.RegistUser(bodyRequest)

		if err != nil {
			// Check if the error is from RegisUser usecase
			u.Failed(c, err)
			return

		}

		successData := dto.SuccessRegistBody{
			Email: bodyRequest.Email,
			Name:  bodyRequest.Name,
			// TODO: ini bikin token logic, sementara masih dummy
			AccessToken: "qwertyuiopasdfghjklzxcvbnm", // This should be the actual access token
		}

		// c.JSON(http.StatusCreated, gin.H{
		// 	"Message": "User logged successfully",
		// 	"data":    successData,
		// })

		data := gin.H{
			"Message": "User register successfully",
			"data":    successData,
		}

		u.Success(c, data)

	}
}

func (u *UserController) requestRegist(c *gin.Context) {
	var bodyRequest dto.RequestRegistBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		fmt.Print("Masuk sini")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		// fmt.Printf("ini isi bodyRequest --> %s", bodyRequest.Email)
		err := u.ucRegist.RequestRegist(bodyRequest)
		if err != nil {
			u.Failed(c, err)
			return
		}
		u.Success(c, nil)
	}
}

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

// func (u *UserController) requestLogin(c *gin.Context) {
// 	var bodyRequest dto.RequestLoginBody

// 	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 		})
// 	} else {
// 		err := u.ucLogin.RequestLogin(bodyRequest)
// 		if err != nil {
// 			u.Failed(c, err)
// 			return
// 		}
// 		u.Success(c, nil)
// 	}
// }

// func (u *UserController) verifyLogin(c *gin.Context) {
// 	var verifOtpBody dto.VerifyLoginBody

// 	if err := u.ParseRequestBody(c, &verifOtpBody); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"message": err.Error(),
// 		})
// 	} else {
// 		bodyResponse, err := u.ucLogin.VerifyLoginOtp(verifOtpBody)
// 		if err != nil {
// 			u.Failed(c, err)
// 			return
// 		}
// 		u.Success(c, bodyResponse)
// 	}
// }

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

	router.POST("/v1/user/login", controller.userLogin)

	// router.POST("/login", controller.requestLogin)
	// router.POST("/login/otp", controller.verifyLogin)

	return &controller
}
