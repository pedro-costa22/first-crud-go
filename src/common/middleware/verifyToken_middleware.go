package middleware

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pedro-costa22/first-crud-go/src/config/rest_err"
)

var (
	JWT_SECRET_KEY = "JWT_SECRET_KEY"
	JWT_EXPIRES_IN = "JWT_EXPIRES_IN"
)

func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv(JWT_SECRET_KEY)
	tokenValue := RemoveBearerPrefix(c.Request.Header.Get("Authorization"))

	token, err := jwt.Parse(RemoveBearerPrefix(tokenValue), func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secret), nil
		}

		return nil, rest_err.NewBadRequestError("Invalid token")
	})

	if err != nil {
		errRest := rest_err.NewBadRequestError("Unauthorized")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}

	_, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewBadRequestError("Unauthorized")
		c.JSON(errRest.Code, errRest)
		c.Abort()
		return
	}
}

func RemoveBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	return token
}