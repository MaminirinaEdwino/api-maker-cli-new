package apimaker

import "github/mameinirinaedwino/api-maker-cli/module/utils"

func GetEndPointOperation() string {
	operation := utils.Scanner("Choose the operation for the endpoint : ")

	switch operation {
	case "create":
		return "create"
	case "read":
		return "read"
	case "update":
		return "update"
	case "delete":
		return "delete"
	case "crud":
		return "crud"
	}

	return GetEndPointOperation()
}