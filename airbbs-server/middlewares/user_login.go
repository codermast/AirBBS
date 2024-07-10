package middlewares

import (
	"codermast.com/airbbs/config"
	"codermast.com/airbbs/models"
	"codermast.com/airbbs/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 放行的路由列表
var unauthenticatedRoutes = []models.Router{
	{"/users/login", "POST"},
	{"/articles/", "GET"},
	{"/articles/**", "GET"},
}

// UserLoginAuthMiddleware 用户登录校验
func UserLoginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestMethod := c.Request.Method
		requestUrl := c.Request.URL.Path

		// 0. 放行指定路由
		for _, route := range unauthenticatedRoutes {
			if (requestUrl == route.Url || matchesWildcardRoute(route.Url, requestUrl)) && requestMethod == route.Method {
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

		claims, err := utils.VerifyJWTToken(jwtToken)

		// 校验失败
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Error("JWT token is invalid"))
			c.Abort()
		} else {
			c.Set("userID", claims.UserID)
			c.Next()
		}
	}
}

// matchesWildcardRoute 检查是否匹配通配符路由
func matchesWildcardRoute(wildcardRoute string, actualRoute string) bool {
	// 这里简单示例，你可以根据实际情况进行更复杂的通配符匹配
	return len(actualRoute) >= len(wildcardRoute) && actualRoute[:len(wildcardRoute)-1] == wildcardRoute[:len(wildcardRoute)-1]
}
