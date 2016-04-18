package main

import (

	"time"

	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"goServe/controller"
	"goServe/auth"
	"goServe/mongodb"
	"goServe/config"
)

func init() {
	config.GetConfig()
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
		api for authentication functionalities 
	*/
	r.POST("/login", auth.LoginUserWithEmail)
	r.POST("/signup/email", auth.SignUpWithEmail)


	/*
		api requiring authentication
	*/
	private := r.Group("/api/private")
	private.Use(auth.TokenAuthMiddleware)
	private.GET("/example", controller.TestJsonController)


	/*
		api not requiring authentication
	*/
	public := r.Group("/api/public")
	public.GET("/example", controller.TestController)



	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
