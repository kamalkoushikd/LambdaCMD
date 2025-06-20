package middleware

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"net/http"
	//"os"
	"strings"
)

var connstring string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"

var jwts = []byte("thisisasupersecretkey")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		fmt.Println(token)
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
		}
		tokenString := strings.TrimPrefix(token, "Bearer")
		tokencontent, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwts, nil

		})
		if tokencontent == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			c.Abort()
			return
		} else {

		}
		c.Next()
	}
}
