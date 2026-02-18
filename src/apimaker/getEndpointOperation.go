package apimaker

import "github/MaminirinaEdwino/api-maker-cli/src/utils"

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
