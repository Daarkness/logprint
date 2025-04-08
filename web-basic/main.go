package main

import (
	"github.com/Daarkness/logprint/web-basic/webook/internal/web"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	server := gin.Default()
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // 允许访问的地址 * 代表所有地址 不推荐使用 *  因为不安全 可以使用白名单 ，
		// 且前端如果使用很严格的规则，这里是不允许使用的
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的方法
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"}, // 允许的头信息

		ExposeHeaders:    []string{"Content-Length"}, // 暴露的头信息
		AllowCredentials: true,                       // 是否允许携带 cookie

		MaxAge: 12 * 60 * 60, // cookie 有效期 12 小时
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "https://") || strings.HasPrefix(origin, "https://") {
				return true
			}
			return false
		},
	}))
	u := web.UserHandler{}
	u.RegisetrRoutes(server)
	err := server.Run(":8080")
	if err != nil {
		return
	}

}
