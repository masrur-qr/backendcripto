package main

import (
	"mongo/repo/controllers"
	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()

	r.POST("/signup", controllers.Signup)
	r.POST("/login",controllers.Login)

	r.Run(":2020")
}

