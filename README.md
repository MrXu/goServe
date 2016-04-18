# goServe
An api server starter kit using gin, mongodb, jwt authentication

goServe is a seed project for you to kickstart your golang restful api project.

## To use
1. get this project and put into your go src directory
2. modify config.json to configure your app
3. go run main.go

### gin
goServe uses [Gin](https://github.com/gin-gonic/gin "gin-gonic") as the framework. Gin provides excellent middleware support and uses one of the best router.

### mongodb
goServe uses [Mgo](https://github.com/go-mgo/mgo "mgo") as the MongoDB driver. Enjoy.

### jwt auth flow
goServe implements the authentication flow with [json web token](https://github.com/dgrijalva/jwt-go "jwt").
Token-based authentication provides better decoupling, scalability and real stateless api. For more information, go to [ins and outs of token-based auth](https://scotch.io/tutorials/the-ins-and-outs-of-token-based-authentication)

