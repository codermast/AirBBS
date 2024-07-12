package utils

import "codermast.com/airbbs/models"

func SuccessData(data interface{}) models.Result {
	return Success("", data)
}

func SuccessMsg(msg string) models.Result {
	return Success(msg, nil)
}

func Success(msg string, data interface{}) models.Result {
	result := models.Result{
		Code:    200,
		Message: msg,
		Data:    data,
	}
	return result
}

func Error(msg string) models.Result {
	result := models.Result{
		Code:    400,
		Message: msg,
		Data:    nil,
	}
	return result
}
