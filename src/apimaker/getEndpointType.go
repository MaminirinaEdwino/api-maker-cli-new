package apimaker

import "github/MaminirinaEdwino/api-maker-cli/src/utils"

func GetEndPointType() string {
	typeEndPoint := utils.Scanner("Choose the type of the endPoint \n" + "Endpoint that Interact with db or not (db / notdb) : ")
	switch typeEndPoint {
	case "db":
		return "db"
	case "notdb":
		return "notdb"
	}
	return GetEndPointType()
}
