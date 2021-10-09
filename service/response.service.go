package service

import "savi8sant8s/api/data"

func MakeRes(httpStatus int, apiStatus string, message string) (int, data.Message){
	return httpStatus, data.Message{
		ApiStatus: apiStatus, 
		Message: message,
	}
}