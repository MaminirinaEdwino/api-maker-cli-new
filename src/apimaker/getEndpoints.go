package apimaker

import (
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
)

func GetEndPoints() ([]basetype.EndPoint, []basetype.EndPoint) {
	var endPointDb []basetype.EndPoint
	var endPointNoDb []basetype.EndPoint
	var typeEndPoint string
	var nameEndPoint string
	var operation string
	var attribute []basetype.Attribut

	nameEndPoint = utils.Scanner("Enter endPoint name : ")
	fmt.Println(nameEndPoint)

	for nameEndPoint != "" {
		typeEndPoint = GetEndPointType()
		fmt.Println(typeEndPoint)
		operation = GetEndPointOperation()
		fmt.Println(operation)
		attribute = GetEndPointAttributes()
		fmt.Println(attribute)
		endPoint := basetype.EndPoint{
			Name:      nameEndPoint,
			Operation: operation,
			Attribut:  attribute,
		}
		if typeEndPoint == "db" {
			endPointDb = append(endPointDb, endPoint)
		} else {
			endPointNoDb = append(endPointNoDb, endPoint)
		}
		fmt.Println(endPoint)
		nameEndPoint = utils.Scanner("Enter endPoint name : ")
	}
	//get endPoint
	return endPointDb, endPointNoDb
}
