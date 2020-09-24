package http

import "github.com/gin-gonic/gin"

func NewURLHandler(router *gin.Engine) {
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, "hello")
	})
}
