package main

import (

	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"gorgeousServer/controller"
	"gorgeousServer/auth"
	"gorgeousServer/mongodb"
)

func init() {
	mongodb.Connect()
}

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
	r.Use(mongodb.DBConnectMW)
	r.Use(mongodb.DBErrorMW)

	/* 
		api 
	*/
	r.GET("/", controller.TestController)
	r.GET("/action", controller.TestJsonController)
	// auth api
	r.POST("/login", auth.LoginUserWithEmail)
	r.POST("/signup/email", auth.SignUpWithEmail)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
