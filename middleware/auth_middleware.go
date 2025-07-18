package middleware

import (
	"mark3/global"
	"mark3/pkg/ctl"
	"mark3/pkg/jwt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctl.UnLogin(nil, ctx)
			ctx.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			global.GVA_LOG.Error(err.Error())
			ctl.UnLogin(nil, ctx)
			ctx.Abort()
			return
		}
		ctx.Set("userid", claims.Userid)
		ctx.Next()
	}
}
