package middlewares

import (
	"codermast.com/airbbs/config"
	"codermast.com/airbbs/constant"
	"codermast.com/airbbs/models"
	"codermast.com/airbbs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strings"
)

// 放行的路由列表
var unauthenticatedRoutes = []models.Router{
	{"/users/login", "POST"},
	{"/users/register", "POST"},
	{"/articles/page", "GET"},
	{`/articles/\d+`, "GET"},
}

// UserLoginAuthMiddleware 用户登录校验
func UserLoginAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		requestMethod := c.Request.Method
		requestUrl := c.Request.URL.Path
		uid := c.Param("uid")

		fmt.Println(uid)

		// 0. 放行指定路由
		for _, route := range unauthenticatedRoutes {
			if (requestUrl == route.Url || matchRoute(requestUrl, route.Url)) && requestMethod == route.Method {
				c.Next()
				return
			}
		}

		// 1. 获取 JWT 信息

		header := c.GetHeader(config.GetJWTConfig().Header)

		if header == "" || !strings.HasPrefix(header, "Bearer") {
			c.JSON(http.StatusUnauthorized, utils.Error("Authorization header is required"))
			c.Abort()
			return
		}

		parts := strings.Split(header, " ")

		if len(parts) != 2 {
			c.JSON(http.StatusUnauthorized, utils.Error("Authorization header is required"))
			c.Abort()
			return
		}

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
			c.Set(constant.USERID, claims.UserID)
			userid, _ := c.Get(constant.USERID)
			userId := c.GetString(constant.USERID)

			fmt.Printf("Get : %s \n", userid)
			fmt.Printf("GetString : %s \n", userId)
			c.Next()
		}
	}
}

// 路由匹配
func matchRoute(sourceRoute string, ruleRoute string) bool {

	// 如果完全匹配则放行
	if sourceRoute == ruleRoute {
		return true
	}

	// 定义正则表达式
	re := regexp.MustCompile(ruleRoute)

	// 返回匹配结果
	return re.MatchString(sourceRoute)
}
