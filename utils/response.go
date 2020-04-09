package utils

import (
	"strings"
)

func DecodeResponse(responseString string) map[string]string {
	responseMap := make(map[string]string)
	responseArr := strings.Split(responseString, "&")
	for _, eachResponse := range responseArr {
		temp := strings.Split(eachResponse, "=")
		responseMap[temp[0]] = temp[1]
	}
	return responseMap
}
