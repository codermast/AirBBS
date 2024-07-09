package middlewares

import (
	"codermast.com/airblog/config"
	"codermast.com/airblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 放行的路由列表
var unauthenticatedRoutes = []string{"login"}

// UserLoginAuthMiddleware 用户登录校验
func UserLoginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestPath := c.Request.URL.Path

		urlParts := strings.Split(requestPath, "/")

		requestUrl := urlParts[len(urlParts)-1]

		// 0. 放行指定路由
		for _, route := range unauthenticatedRoutes {
			if requestUrl == route {
				c.Next()
				return
			}
		}

		// 1. 获取 JWT 信息

		header := c.GetHeader(config.GetJWTConfig().Header)

		if header == "" {
			c.JSON(http.StatusUnauthorized, utils.Error("Authorization header is required"))
			c.Abort()
			return
		}

		parts := strings.Split(header, " ")
		jwtToken := parts[1]

		// 检查 Authorization 头部格式，应为 "Bearer {token}"
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, utils.Error("Authorization header format must be Bearer {token}"))
			c.Abort()
			return
		}

		err := utils.VerifyJWTToken(jwtToken)

		// 校验失败
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Error("JWT token is invalid"))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
