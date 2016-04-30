# goServe
An api server starter kit using gin, mongodb, jwt authentication

goServe is a seed project for you to kickstart golang restful api project.It tries to mimic Django by providing essential package such as basic authentication, social authentication, database connection and more to come. However, unlike Django, goServe only tries to serve the needs of Single Page Application and mobile application. 

## To use
1. put this project into go src directory
2. modify config.json to configure your app (mongodb url, jwt key, email)
3. go run main.go

### gin
goServe uses [Gin](https://github.com/gin-gonic/gin "gin-gonic") as the framework. Gin provides excellent middlewares and uses httprouter for its performance.

### mongodb
goServe uses [Mgo](https://github.com/go-mgo/mgo "mgo") as the MongoDB driver. Enjoy.

### jwt auth flow
goServe implements the authentication flow with [json web token](https://github.com/dgrijalva/jwt-go "jwt").
Token-based authentication provides better decoupling, scalability and real stateless api. For more information, go to [ins and outs of token-based auth](https://scotch.io/tutorials/the-ins-and-outs-of-token-based-authentication)

Auth supports:

1. email password
2. facebook auth 

