package controller

import (
	"7Zero4/delivery/api"
	"7Zero4/model/dto"
	"7Zero4/usecase"
	"7Zero4/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router    *gin.RouterGroup
	routerDev *gin.RouterGroup
	ucLogin   usecase.UserLoginUseCase
	ucRegist  usecase.UserRegistrationUsecase
	ucToken   usecase.TokenUsecase
	ucCat     usecase.CatUseCase
	api.BaseApi
}

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

			AccessToken: "qwertyuiopasdfghjklzxcvbnm", // This should be the actual access token
		}

		detailMsg := "User logged successfully "

		u.Success(c, successData, detailMsg)
	}
}

func (u *UserController) userRegister(c *gin.Context) {
	var bodyRequest dto.RequestRegistBody

	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {

		u.Failed(c, utils.ServerError())
		return

	} else {
		token, err := u.ucRegist.RegisterUser(bodyRequest)

		if err != nil {
			u.Failed(c, err)
			return
		}

		successData := dto.SuccessRegistBody{
			Email:       bodyRequest.Email,
			Name:        bodyRequest.Name,
			AccessToken: token,
		}

		detailMsg := "User register successfully "

		u.Success(c, successData, detailMsg)
	}

}

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

func (u *UserController) createCat(c *gin.Context) {
	var bodyRequest dto.RequestCreateCat
	if err := u.ParseRequestBody(c, &bodyRequest); err != nil {
		u.Failed(c, utils.ReqBodyNotValidError())
		return
	}

	createCatErr := u.ucCat.CreateCat(bodyRequest)
	if createCatErr != nil {
		if utils.IsValidationError(createCatErr) {
			c.JSON(http.StatusBadRequest, gin.H{
				"errorCode": "400",
				"message":   "create cat error: " + createCatErr.Error(),
			})
			return
		}

		// add 401 for token

		c.JSON(http.StatusInternalServerError, gin.H{
			"errorCode": "500",
			"message":   "Internal Server Error: " + createCatErr.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "success",
		"data": dto.SuccessCreateCat{
			ID:        strconv.Itoa(rand.Intn(26)),
			CreatedAt: time.Now().Format(time.RFC3339),
		},
	})
}

func NewUserController(router *gin.RouterGroup, routerDev *gin.RouterGroup, ucLogin usecase.UserLoginUseCase, ucRegist usecase.UserRegistrationUsecase, ucCat usecase.CatUseCase) *UserController {
	controller := UserController{
		router:    router,
		routerDev: routerDev,
		ucLogin:   ucLogin,
		ucRegist:  ucRegist,
		ucCat:     ucCat,

		BaseApi: api.BaseApi{},
	}

	// USER REGISTER AND LOGIN
	// router.POST("/register/otp", controller.verifyRegist)
	// router.POST("/register", controller.requestRegist)

	router.POST("/v1/user/register", controller.userRegister)
	router.POST("/v1/user/login", controller.userLogin)

	// manage cat
	router.POST("/v1/cat", controller.createCat)
	// router.POST("/login/otp", controller.verifyLogin)

	return &controller
}
