package apimaker

import (
	"fmt"
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
	"github/mameinirinaedwino/api-maker-cli/module/utils"
)

func CreateProject() {
	var projectname string
	var sgbd string
	var db_name string
	var EndPointDB []basetype.EndPoint
	var EndPointNotDB []basetype.EndPoint

	projectname = utils.Scanner("Enter the name of your api : ")
	fmt.Println(projectname)
	sgbd = GetSGBD()
	fmt.Println("SGBD : " + sgbd)
	db_name = utils.Scanner("Enter the name of the database :")
	fmt.Println(db_name)

	fmt.Println(len(EndPointDB))
	fmt.Println(len(EndPointNotDB))

	EndPointDB, EndPointNotDB = GetEndPoints()
	fmt.Println(EndPointDB)
	fmt.Println(EndPointNotDB)

	WriteCode(projectname, sgbd, db_name, EndPointDB, EndPointNotDB)
}