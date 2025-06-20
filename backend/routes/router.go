package routes

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var jwtSecret = []byte("thisisasupersecretkey")
var connstring string = "postgres://iithpool:password123@localhost:5432/canteenpool?sslmode=disable"

func RegisterRoutes(r *gin.Engine) {
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		panic(err)
	}

	//defer db.Close()
	fmt.Println("Registering routes...")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", func(context *gin.Context) {
		var data map[string]string
		err := context.BindJSON(&data)

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid json",
			})
			return
		}
		password := data["password"]
		username := data["username"]
		var passhash string
		err = db.QueryRow("select passhash from users where username= $1", username).Scan(&passhash)
		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusBadRequest, gin.H{

				"message": "user does not exist",
			})
			return
		}

		fmt.Println(passhash)
		err = bcrypt.CompareHashAndPassword([]byte(passhash), []byte(password))
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"message": "wrong password",
			})
			return
		} else {
			fmt.Println("success")
			token, tokerr := generateToken(username)
			if tokerr != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"message": "error generating token",
				})
				return
			} else {
				context.JSON(http.StatusOK, gin.H{
					"token": token,
				})
			}
		}
		fmt.Println(data)
	})
}

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
		"iat":      time.Now().Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return jwtToken.SignedString(jwtSecret)
}
