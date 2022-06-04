package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//接受服务实例,并存到gin.key中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//将实例存在gin.key中
		ctx.Keys = make(map[string]interface{})
		ctx.Keys["UserService"] = service[0]
		ctx.Next()
	}
}

//错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			/*
				Recover 是一个Go语言的内建函数，
				可以让进入宕机流程中的 goroutine 恢复过来，
				recover 仅在延迟函数 defer 中有效
				recover 的宕机恢复机制就对应其他语言中的 try/catch 机制

			*/
			r := recover()
			if r != nil {
				ctx.JSON(200, gin.H{
					"code": http.StatusNotFound,
					"msg":  fmt.Sprintf("%s", r),
				})
				//下面不执行了
				ctx.Abort()
			}
		}()
		ctx.Next()
	}
}
