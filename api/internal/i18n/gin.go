package i18n

import (
	"github.com/gin-gonic/gin"
	"mark3/internal/middleware"
)

// GinContextKey 定义Gin上下文的key
const GinContextKey = "i18n"

// GinMiddleware 创建Gin中间件的便捷方法
func (i *I18n) GinMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := middleware.DetectLanguage(c.Request)
		i.SetLanguage(lang)
		c.Set(GinContextKey, i)
		middleware.SetLanguageCookie(c.Writer, lang)
		c.Next()
	}
}

// FromGinContext 从Gin上下文中获取i18n实例
func FromGinContext(c *gin.Context) *I18n {
	if val, exists := c.Get(GinContextKey); exists {
		if i, ok := val.(*I18n); ok {
			return i
		}
	}
	return nil
}
