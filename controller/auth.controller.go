package controller

import (
	"net/http"
	"xpm-auth/data/request"
	"xpm-auth/data/response"
	service "xpm-auth/service/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService service.AuthService
}

func NewAuthController(service service.AuthService) *AuthController {
	return &AuthController{
		authService: service,
	}
}

func (controller *AuthController) Register(ctx *gin.Context) {
	createUserRequest := request.RegisterRequest{}
	err := ctx.ShouldBindJSON(&createUserRequest)
	ctx.Header("Content-Type", "application/json")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}

	result, customError := controller.authService.Register(createUserRequest)
	if customError != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Fail",
			Data:   customError.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, response.Response{
			Code:   http.StatusOK,
			Status: "Ok",
			Data:   result,
		})
	}

}

func (controller *AuthController) Login(ctx *gin.Context) {
	loginRequest := request.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginRequest)
	ctx.Header("Content-Type", "application/json")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad Request",
			Data:   err.Error(),
		})
		return
	}
	token, jwtErr := controller.authService.Login(loginRequest)
	ctx.SetCookie("access_token", token, 3600, "/", "localhost", true, true)
	if jwtErr != nil {
		ctx.JSON(http.StatusUnauthorized, response.Response{
			Code:   http.StatusUnauthorized,
			Status: "Fail",
			Data:   jwtErr.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   "Login Successful.",
	})
}
