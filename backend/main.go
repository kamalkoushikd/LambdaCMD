package main

import (
	"fmt"
	// "log"
	"backend/routes"
	"github.com/gin-gonic/gin"
)

var jwtSeret = []byte("thisisasupersecretkey")

func main() {

	r := gin.Default()
	routes.RegisterRoutes(r)
	err := r.Run(":8080")
	fmt.Println(err)
}
