package daos

import (
	"codermast.com/airbbs/models/po"
	"errors"
)

func CreateBlink(blink *po.Blink) error {
	result := DB.Table("blinks").Create(blink)

	if result.Error != nil || result.RowsAffected == 0 {
		return errors.New("创建动态失败！")
	}

	return nil
}

func GetBlinkList(blinkList *[]po.Blink) error {

	result := DB.Table("blinks").Order("id desc").Find(blinkList)

	return result.Error
}
