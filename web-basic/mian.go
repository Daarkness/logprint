package main

import (
	"github.com/Daarkness/logprint/web-basic/web"
	"github.com/gin-gonic/gin"
)

func main() {
	web.Userinfo()
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.Run(":8080")

}
