package apimaker

import "github/MaminirinaEdwino/api-maker-cli/src/utils"

func GetAttrType() string {
	var attr_type string = utils.Scanner("Choose the field type\n1. int\n2. string\n3. bool\n4. float\n=> ")
	switch attr_type {
	case "string":
	case "int":
	case "bool":
	case "float":
		return "float32"
	case "relation":
	default:
		GetAttrType()
	}
	return attr_type
}
