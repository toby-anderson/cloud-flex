package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/toby-anderson/cloud-flex/utils/token"
)

func JwtAuthHandler() gin.HandlerFunc {
	return func(ginc *gin.Context) {
		err := token.TokenValid(ginc)
		if err != nil {
			ginc.String(http.StatusUnauthorized, "Unauthorized")
			ginc.Abort()
			return
		}
		ginc.Next()
	}
}
