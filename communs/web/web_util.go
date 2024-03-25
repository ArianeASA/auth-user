package web

import "encoding/json"

type ErrorResponse struct {
	Cause string
}

func NewError(msg string) ErrorResponse {
	return ErrorResponse{Cause: msg}
}

func GetJson(obj interface{}) string {
	objJson, _ := json.Marshal(obj)
	return string(objJson)
}

func Headers() map[string]string {
	return map[string]string{
		"Content-Type": "application/json",
	}
}
