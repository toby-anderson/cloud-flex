package token

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/satori/go.uuid"
)

func GenerateToken(user_id uuid.UUID) (string, error) {

	token_lifespan, err := strconv.Atoi(os.Getenv("JWT_HOUR_LIFESPAN"))

	if err != nil {
		fmt.Println("Using default JWT lifespan of 1 hour")
		token_lifespan = 1
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(token_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(os.Getenv("API_SECRET")))

}

func TokenValid(ginc *gin.Context) error {
	tokenString := ExtractToken(ginc)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func ExtractToken(ginc *gin.Context) string {
	token := ginc.Query("token")
	if token != "" {
		return token
	}
	bearerToken := ginc.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	} else if len(strings.Split(bearerToken, " ")) == 1 {
		return bearerToken
	}
	return ""
}

func ExtractTokenID(ginc *gin.Context) (uuid.UUID, error) {

	tokenString := ExtractToken(ginc)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return uuid.NewV4(), err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := uuid.FromString(claims["user_id"].(string))
		if err != nil {
			return uuid.NewV4(), err
		}
		return uid, nil
	}
	return uuid.NewV4(), nil
}
