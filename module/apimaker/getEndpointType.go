package apimaker

import "github/mameinirinaedwino/api-maker-cli/module/utils"

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