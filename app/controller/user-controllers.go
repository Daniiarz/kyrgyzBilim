package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/openpgp/errors"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/service"
	"net/http"
)

type LoginData struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	user := &entity.User{}
	if err := c.ShouldBindJSON(&user); err != nil {

		var invalidArgs []errors.InvalidArgumentError

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errs{
				invalidArgs =
			}
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := service.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func Login(c *gin.Context) {
	loginInfo := &LoginData{}
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong credentials"})
		return
	}
	tokens, err := service.SingIn(loginInfo.PhoneNumber, loginInfo.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}

func Refresh(c *gin.Context) {
	tokens := &service.AuthTokens{}
	if err := c.ShouldBindJSON(&tokens); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong tokens"})
		return
	}
	err := service.RefreshAccessToken(tokens)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}
