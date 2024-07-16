package middlewares

import (
	"codermast.com/airbbs/constant"
	"codermast.com/airbbs/daos"
	"codermast.com/airbbs/models"
	"codermast.com/airbbs/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 拦截的路由列表
var adminAuthenticatedRoutes = []models.Router{
	{"/articles/all", "GET"},
	{"/articles/", "DELETE"},
}

// UserAdminAuthMiddleware 是否为管理员校验
func UserAdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestMethod := c.Request.Method
		requestUrl := c.Request.URL.Path

		// 0. 放行除拦截之外的其他路由
		for _, route := range adminAuthenticatedRoutes {
			if !(requestUrl == route.Url || matchesWildcardRoute(route.Url, requestUrl)) && requestMethod == route.Method {
				c.Next()
				return
			}
		}

		// 1. 获取 User 信息
		userId := c.GetString(constant.USERID)

		// 2. 判断用户是否为管理员

		err := daos.IsAdminByUserId(userId)
		if err != nil {
			// 不是则拦截
			c.JSON(http.StatusUnauthorized, utils.Error("权限不足！"))
			c.Abort()
		} else {
			// 是管理员则放行
			c.Next()
		}
	}
}