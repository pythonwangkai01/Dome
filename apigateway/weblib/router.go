package weblib

import (
	"apigateway/weblib/handlers"
	"apigateway/weblib/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(200, "hello success")
		})

		//用户服务
		v1.POST("/user/register", handlers.UserRegister)
		v1.POST("/user/login", handlers.UserLogin)

		//token登录
		authod := v1.Group("/")
		authod.Use(middleware.JWT())
		v1.DELETE("/user/delte", handlers.UserDelte)
		v1.POST("/users", handlers.GetUsersList)
		v1.POST("/user", handlers.GetUser)
	}

	return ginRouter
}
