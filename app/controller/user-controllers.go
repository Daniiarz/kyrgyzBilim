package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository/database"
	"net/http"
)

type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	user := &entity.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.SetPassword()

	database.DB.Create(&user)

	c.JSON(http.StatusCreated, user)
}

func Login(c *gin.Context) {
	loginInfo := &LoginData{}
	user := &entity.User{}

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Where("email = ?", loginInfo.Email).First(&user)

	if user.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No user with given r email"})
		return
	}
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(loginInfo.Password), bcrypt.DefaultCost)
	print(hashedPass, "\n")
	print(user.Password, "\n")
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User was found!"})
}

func TestHashing(c *gin.Context) {
	hashed1, _ := bcrypt.GenerateFromPassword([]byte("1"), bcrypt.DefaultCost)
	//hashed2, _ := bcrypt.GenerateFromPassword([]byte("1"), bcrypt.DefaultCost)

	err := bcrypt.CompareHashAndPassword(hashed1, []byte("1"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"err": "Ok everything is fine"})
}
