package apiserver

import "github.com/gin-gonic/gin"

type server struct {
	router *gin.Engine
}

func newServer() *server {
	s := &server{
		router: gin.Default(),
	}
	return s
}
