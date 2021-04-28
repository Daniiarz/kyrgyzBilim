package controller

import (
	"github.com/gin-gonic/gin"
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
	obj, ok := service.DataBind(c, user)
	if !ok {
		c.JSON(http.StatusBadRequest, obj.(gin.H))
		return
	}
	parsedUser := obj.(*entity.User)
	fileUuid := service.UploadHandler(c, "profile_picture")
	user.ProfilePicture = "" + fileUuid
	err := service.RegisterUser(parsedUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}

func Login(c *gin.Context) {
	login := &LoginData{}
	obj, ok := service.DataBind(c, login)
	if !ok {
		c.JSON(http.StatusBadRequest, obj.(gin.H))
		return
	}
	parsedLogin := obj.(*LoginData)
	tokens, err := service.SingIn(parsedLogin.PhoneNumber, parsedLogin.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}

func Refresh(c *gin.Context) {
	tokens := &service.AuthTokens{}
	obj, ok := service.DataBind(c, tokens)
	if !ok {
		c.JSON(http.StatusBadRequest, obj.(gin.H))
		return
	}
	parsedTokens := obj.(*service.AuthTokens)
	err := service.RefreshAccessToken(parsedTokens)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tokens)
}
