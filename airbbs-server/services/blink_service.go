package services

import (
	"codermast.com/airbbs/daos"
	"codermast.com/airbbs/models/po"
)

// 创建动态
func CreateBlink(blink *po.Blink) error {
	return daos.CreateBlink(blink)
}

func GetBlinkList(blinkList *[]po.Blink) error {
	return daos.GetBlinkList(blinkList)
}
