package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
r:=gin.Default()
r.GET("/test", func(c *gin.Context) {
	c.String(http.StatusOK,"this is a get http")
	c.JSON(http.StatusOK,gin.H{
		"key":"value",
	})
})
r.Run(":8080")
}
