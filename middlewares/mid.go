package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWTMiddleware() echo.MiddlewareFunc {

	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: middleware.AlgorithmHS256,
		SigningKey:    []byte(os.Getenv("SECRET_JWT")),
	})

}

func CreateToken(id int, email, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["role"] = role
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Duration(1) * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}

func ExtractTokenEmail(e echo.Context) string {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		peran := claims["email"].(string)
		return peran
	}
	return ""
}
func ExtractTokenRole(e echo.Context) string {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		peran := claims["role"].(string)
		return peran
	}
	return ""
}
func ExtractTokenId(e echo.Context) int {
	team := e.Get("user").(*jwt.Token)
	if team.Valid {
		claims := team.Claims.(jwt.MapClaims)
		userId := claims["id"].(float64)
		return int(userId)
	}
	return 0
}
