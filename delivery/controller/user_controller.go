package controller

import (
	"7Zero4/delivery/api"
	"7Zero4/model/dto"
	"7Zero4/usecase"
	"7Zero4/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router    *gin.RouterGroup
	routerDev *gin.RouterGroup
	ucLogin   usecase.UserLoginUseCase
	ucRegist  usecase.UserRegistrationUsecase
	ucToken   usecase.TokenUsecase
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
		return

	}
	token, err := u.ucRegist.RegisterUser(bodyRequest)
	if err != nil {
		// Email conflict exist
		if utils.IsErrDuplicateValueFound(err) {
			c.JSON(http.StatusConflict, gin.H{
				"errorCode": "409",
				"message":   "Email already registered",
			})
			return
		}

		// Validation Error
		if utils.IsValidationError(err) {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorCode": "400",
				"message":   "Register user error: " + err.Error(),
			})
			return
		}

		// Handle StatusInternalServerError (500)
		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode": "500",
			"message":   "Internal Server Error: " + err.Error(),
		})

	}

	successData := dto.SuccessRegistBody{
		Email:       bodyRequest.Email,
		Name:        bodyRequest.Name,
		AccessToken: token,
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "User registered successfully",
		"data":    successData,
	})
}

func (u *UserController) requestLogin(c *gin.Context) {
	// var bodyRequest dto.RequestLoginBody
}

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
