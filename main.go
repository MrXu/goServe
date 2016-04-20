package main

import (
	Config "goServe/config"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/itsjamie/gin-cors"
	"goServe/controller"
	"goServe/auth"
	Mongo "goServe/mongodb"
	Email "goServe/email"
)

func init() {
	Config.GetConfig()
	Email.ConfigEmail()
	Mongo.Connect()
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
	r.Use(Mongo.DBConnectMW)
	r.Use(Mongo.DBErrorMW)

	/* 
		api for authentication functionalities 
	*/
	r.POST("/login", auth.LoginUserWithEmail)
	r.POST("/signup/email", auth.SignUpWithEmail)
	r.POST("/signup/email/verify-email", auth.ActivateAccountAfterEmailSignup)
	facebookR := r.Group("facebook")
	auth.ConfigFacebook(facebookR, Config.GetFacebookClientId(), Config.GetFacebookClientSecret(), "http://localhost:8080/facebook/callback")

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
