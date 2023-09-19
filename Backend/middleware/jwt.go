package middleware

import (
	"net/http"
	"resume_backend/pkg/ctl"
	"resume_backend/pkg/e"
	"resume_backend/pkg/utils"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		code := e.Success
		token := ctx.GetHeader("Authorization")
		if token == "" {
			code = http.StatusNotFound
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "缺少token",
			})
			ctx.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			code = e.Error
			ctx.JSON(e.InvalidParams, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   "解析token失败",
			})
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(ctl.NewContext(ctx.Request.Context(), &ctl.UserInfo{
			Id:       claims.Id,
			UserName: claims.UserName,
		}))
		ctx.Next()
	}
}
