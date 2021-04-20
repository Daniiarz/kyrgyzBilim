package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"kyrgyz-bilim/entity"
	"kyrgyz-bilim/repository"
	"os"
	"strconv"
	"time"
)

type AuthTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

var secretKey = os.Getenv("SECRET_KEY")
var connection repository.UserRepository

func SingIn(phone, password string) (*AuthTokens, error) {
	connection = repository.NewUserRepository()
	user := connection.GetUserByPhone(phone)
	if user == nil {
		return &AuthTokens{}, errors.New("no user found")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+secretKey))
	if err != nil {
		return &AuthTokens{}, err
	}
	tokens := ObtainTokenPair(user)
	return tokens, nil
}

func RegisterUser(user *entity.User) error {
	connection = repository.NewUserRepository()
	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password+secretKey), bcrypt.DefaultCost)
	user.Password = string(password)
	connection.Create(user)
	if user.Id == 0 {
		return errors.New("such user already exists")
	}
	return nil
}

func ObtainTokenPair(user *entity.User) *AuthTokens {
	return &AuthTokens{
		AccessToken:  newAccessToken(user),
		RefreshToken: newRefreshToken(user),
	}
}

func newAccessToken(user *entity.User) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
		IssuedAt:  time.Now().Unix(),
	})
	accessToken, _ := claims.SignedString([]byte(secretKey))

	return accessToken
}

func newRefreshToken(user *entity.User) string {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24 * 3).Unix(), //1 day
		IssuedAt:  time.Now().Unix(),
		Issuer:    strconv.Itoa(int(user.Id)),
	})
	refreshToken, _ := claims.SignedString([]byte(secretKey))

	return refreshToken
}

func RefreshAccessToken(tokens *AuthTokens) error {
	user, err := ValidateToken(tokens.RefreshToken)
	if err != nil {
		return err
	}
	tokens.AccessToken = newAccessToken(user)
	return nil
}

func ValidateToken(token string) (*entity.User, error) {
	authToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := authToken.Claims.(jwt.MapClaims); ok && authToken.Valid {
		connection = repository.NewUserRepository()
		id, err := strconv.Atoi(claims["iss"].(string))
		if err != nil {
			return nil, errors.New("not a valid valid token")
		}
		user := connection.GetUserById(id)
		if user.Id == 0 {
			return nil, errors.New("not a valid valid token")
		}
		return user, nil
	} else {
		return nil, errors.New("not a valid token")
	}
}
