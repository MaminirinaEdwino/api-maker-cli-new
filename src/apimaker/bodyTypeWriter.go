package apimaker

import (
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"strings"
)

func WriteBodyType(endPoint basetype.EndPoint) string {
	var AttrList string
	for _, attr := range endPoint.Attribut {
		AttrList += fmt.Sprintf("%s %s `json:\"%s\"`\n", strings.ToUpper(attr.Nom), attr.Type, attr.Nom)
	}
	return fmt.Sprintf(`
	type %sbodyType struct{
		//ID
		%s
	}
	`, endPoint.Name, AttrList)
}
