package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"google.golang.org/appengine"
)

func main() {
	r := gin.Default()
	ApiRouter(r.Group("/api"))
	appengine.Main()
}

func ApiRouter(api *gin.RouterGroup)  {
	api.GET("/",GetID)
}

func GetID(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"text":"Hello World"})
}