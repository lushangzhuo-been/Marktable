package middleware

import (
	"github.com/gin-gonic/gin"
	"mark3/internal/i18n"
	"net/http"
)

const (
	langQueryParam  = "lang"
	langCookieName  = "lang"
	defaultLanguage = "en"
)

func I18nMiddleware(i *i18n.I18n) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 获取语言设置
		lang := detectLanguage(c.Request)

		// 2. 设置到i18n实例
		i.SetLanguage(lang)

		// 3. 存储到Gin上下文
		c.Set("i18n", i)

		// 4. 设置语言cookie
		setLanguageCookie(c.Writer, lang)

		c.Next()
	}
}

func DetectLanguage(r *http.Request) string {
	// 优先级: 1.URL参数 > 2.Cookie > 3.请求头
	if lang := r.URL.Query().Get(langQueryParam); lang != "" {
		return lang
	}

	if cookie, err := r.Cookie(langCookieName); err == nil {
		return cookie.Value
	}

	if acceptLang := r.Header.Get("Accept-Language"); acceptLang != "" {
		return parseAcceptLanguage(acceptLang)
	}

	return defaultLanguage
}

func parseAcceptLanguage(header string) string {
	// 简单实现：取第一个语言标签的前两位
	if len(header) >= 2 {
		return header[:2]
	}
	return defaultLanguage
}

func SetLanguageCookie(w http.ResponseWriter, lang string) {
	http.SetCookie(w, &http.Cookie{
		Name:     langCookieName,
		Value:    lang,
		Path:     "/",
		MaxAge:   365 * 24 * 60 * 60, // 1年
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
}
