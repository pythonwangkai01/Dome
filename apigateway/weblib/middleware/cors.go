package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

//Cors 跨域配置
func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method

		//请求头部
		origin := ctx.Request.Header.Get("Origin")
		//设置header
		var headerKeys []string
		for k := range ctx.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			//允许访问所有域 就个文件的关键
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//  header的类型
			ctx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置,跨域关键设置 让浏览器可以解析                                                                                                 可以返回其他子段
			ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			// 缓存请求信息 单位为秒
			ctx.Header("Access-Control-Max-Age", "172800")
			//跨域请求是否需要带cookie信息 默认设置为true                                                                                                                                                           // 缓存请求信息 单位为秒
			ctx.Header("Access-Control-Allow-Credentials", "false")
			//设置返回格式是json                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			ctx.Set("content-type", "application/json")
		}
		// OPTIONS指所有的请求类型
		if method == "OPTIONS" {
			ctx.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		ctx.Next()
	}
}
