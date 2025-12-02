package apimaker

import (
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
	"strings"
)

func WriteResponseType(endPoint basetype.EndPoint) string {
	responseBody := strings.Replace(WriteBodyType(endPoint), "//ID", "ID int `json:\"id\"`", 1)
	responseBody = strings.Replace(responseBody, "bodyType", "responseType", 1)
	return responseBody
}