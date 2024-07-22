package utils

import (
	"codermast.com/airbbs/models/pojo"
)

func SuccessData(data interface{}) pojo.Result {
	return Success("", data)
}

func SuccessMsg(msg string) pojo.Result {
	return Success(msg, nil)
}

func Success(msg string, data interface{}) pojo.Result {
	result := pojo.Result{
		Code:    200,
		Message: msg,
		Data:    data,
	}
	return result
}

func Error(msg string) pojo.Result {
	result := pojo.Result{
		Code:    400,
		Message: msg,
		Data:    nil,
	}
	return result
}
