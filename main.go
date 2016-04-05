package main

import (
	"github.com/gin-gonic/gin"
	"gorgeous/controller"
)

func main() {
	r := gin.New()

	// middleware
	r.Use(gin.Logger())

	// api
	r.GET("/", controller.TestController)
	r.GET("/action", controller.TestJsonController)

	// js app
	r.Static("asset", "./asset")
	r.Static("app", "./app")

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
