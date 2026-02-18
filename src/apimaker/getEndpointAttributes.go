package apimaker

import (
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
)

func GetEndPointAttributes() []basetype.Attribut {
	var attr_list []basetype.Attribut
	var temp basetype.Attribut

	temp.Nom = utils.Scanner("Enter a attribute Name : ")
	for temp.Nom != "" {
		temp.Type = GetAttrType()
		attr_list = append(attr_list, temp)
		temp.Nom = utils.Scanner("Enter a attribute Name : ")
	}

	return attr_list
}
