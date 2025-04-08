package web

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes() *gin.Engine {
	server := gin.Default()
	reqisterUserRoutes(server)
	return server
}

func reqisterUserRoutes(r *gin.Engine) {
	u := &UserHandler{}
	r.POST("/login", u.Login)
}
