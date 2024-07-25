package controllers

import (
	"codermast.com/airbbs/constant"
	"codermast.com/airbbs/models/po"
	"codermast.com/airbbs/models/ro"
	"codermast.com/airbbs/services"
	"codermast.com/airbbs/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"net/http"
)

type BlinkController struct{}

func NewBlinkController() *BlinkController {
	return &BlinkController{}
}

// CreateBlink 发布动态
func (bl *BlinkController) CreateBlink(c *gin.Context) {
	var blinkCreateRo ro.BlinkCreateRequest

	if err := c.ShouldBindJSON(&blinkCreateRo); err != nil {
		c.JSON(http.StatusBadRequest, utils.Error("动态格式解析错误！"))
		return
	}

	var blink po.Blink
	curUserId := c.GetString(constant.USERID)
	blink.Author = curUserId

	_ = copier.Copy(&blink, &blinkCreateRo)

	err := services.CreateBlink(&blink)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("发布成功！", blink))
}

// GetBlinkList 查看动态列表 GET /blink/list
func (bl *BlinkController) GetBlinkList(c *gin.Context) {
	var blinkList []po.Blink

	err := services.GetBlinkList(&blinkList)

	if err != nil {
		c.JSON(http.StatusBadRequest, utils.Error(fmt.Sprintf("%v", err)))
		return
	}

	c.JSON(http.StatusOK, utils.Success("查看成功！", blinkList))
}
