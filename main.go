package main

import (

	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"gorgeous/controller"
	"gorgeous/auth"
)

func main() {
	r := gin.New()

	/* 
		middleware 
	*/
	r.Use(gin.Logger())
	r.Use(cors.Middleware(cors.Config{
	    Origins:        "*",
	    Methods:        "GET, PUT, POST, DELETE",
	    RequestHeaders: "Origin, Authorization, Content-Type",
	    ExposedHeaders: "",
	    MaxAge: 50 * time.Second,
	    Credentials: true,
	    ValidateHeaders: false,
	}))

	/* 
		api 
	*/
	r.GET("/", controller.TestController)
	r.GET("/action", controller.TestJsonController)
	// auth api
	r.POST("/login", auth.LoginUserWithPassword)


	// js app
	r.Static("asset", "./asset")
	r.Static("app", "./app")

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
