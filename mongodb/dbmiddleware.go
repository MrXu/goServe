package mongodb

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	DBMiddlewareName = "db"
)

func DBConnectMW(c *gin.Context) {
	s := Session.Clone()
	defer s.Close()
	c.Set(DBMiddlewareName, s.DB(Mongo.Database))
	c.Next()
}

func DBErrorMW(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors":c.Errors})
	}
}
