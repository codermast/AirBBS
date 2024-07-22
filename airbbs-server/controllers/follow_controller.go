package controllers

import (
	"codermast.com/airbbs/constant"
	"codermast.com/airbbs/models/vo"
	"codermast.com/airbbs/services"
	"codermast.com/airbbs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 关注控制器

type FollowController struct{}

func NewFollowController() *FollowController {
	return &FollowController{}
}

// FollowUser 关注指定用户	POST /follow/:id
func (fc *FollowController) FollowUser(c *gin.Context) {
	// 目标用户ID
	targetID := c.Param("id")

	// 当前登录用户ID
	userID := c.GetString(constant.USERID)

	if targetID == "" || userID == "" {
		c.JSON(http.StatusBadRequest, utils.Error("targetID or userID is required"))
		return
	}

	err := services.FollowUser(userID, targetID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("关注成功！"))
}

// UnfollowUser 取消关注指定用户	DELETE /follow/:id
func (fc *FollowController) UnfollowUser(c *gin.Context) {
	// 目标用户ID
	targetID := c.Param("id")

	// 当前登录用户ID
	userID := c.GetString(constant.USERID)

	if targetID == "" || userID == "" {
		c.JSON(http.StatusBadRequest, utils.Error("targetID or userID is required"))
		return
	}

	err := services.UnfollowUser(userID, targetID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessMsg("取关成功！"))
}

// GetUserFans 查看指定用户的粉丝列表	GET /follow/:id
func (fc *FollowController) GetUserFans(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		c.JSON(http.StatusBadRequest, utils.Error("userId is required"))
		return
	}

	// 粉丝列表
	var fans []vo.UserVO

	fans, err := services.GetUserFans(userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Error("failed to get user fans"))
		return
	}

	c.JSON(http.StatusOK, utils.Success("查询成功！", fans))
}

//// IsMutualFollow 检查是否互关	GET /follow/:id1/:id2
//func (fc *FollowController) IsMutualFollow(c *gin.Context) {
//	userId1 := c.Param("id1")
//	userId2 := c.Param("id2")
//
//	// 是否互关判断
//	isMutualFollow := services.IsMutualFollow(userId1, userId2)
//
//	if isMutualFollow {
//		c.JSON(http.StatusOK, utils.SuccessMsg("已互关！"))
//	} else {
//		c.JSON(http.StatusOK, utils.SuccessMsg("未互关！"))
//	}
//}
