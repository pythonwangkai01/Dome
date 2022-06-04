package middleware

import (
	"apigateway/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code uint32
		code = 200
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusMethodNotAllowed
		} else {
			_, err := utils.ParseToken(token)
			if err != nil {
				code = http.StatusUnauthorized
			}
		}
		if code != 200 {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
