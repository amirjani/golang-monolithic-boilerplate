package Controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mahdidl/golang_boilerplate/Common/Helper"
	"github.com/mahdidl/golang_boilerplate/Common/Response"
	"github.com/mahdidl/golang_boilerplate/Common/Validator"
	User "github.com/mahdidl/golang_boilerplate/Components/Auth/Request"
	Response2 "github.com/mahdidl/golang_boilerplate/Components/Auth/Response"

	"log"

	"net/http"
)

type AuthController struct {
	authService *AuthService
}

func NewAuthController(authService *AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (authController *AuthController) AccessToken(context *gin.Context) {
	var accessTokenReq User.AccessTokenRequest
	Helper.Decode(context.Request, &accessTokenReq)

	validationError := Validator.ValidationCheck(accessTokenReq)
	log.Println(validationError)
	if validationError != nil {
		response := Response.GeneralResponse{Error: true, Message: validationError.Error()}
		context.JSON(http.StatusBadRequest, gin.H{"response": response})
	}

	token, err := authController.authService.CreateAccessToken(accessTokenReq)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"response": err})
	}

	// all ok
	// create general response
	response1 := Response.GeneralResponse{Error: false, Message: "successful", Data: Response2.AccessTokenResponse{AccessToken: token}}
	context.JSON(http.StatusOK, gin.H{"response": response1})
}
