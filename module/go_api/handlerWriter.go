package goapi

import (
	"fmt"
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
	"github/mameinirinaedwino/api-maker-cli/module/utils"
)
func InsertHandler(ep basetype.EndPoint, sgbd string, )string{
	return fmt.Sprintf(`
	func %sHandlePost(w http.ResponseWriter, r *http.Request){
		%s
		%s
	res, err := db.Exec("%s", %s)
		%s
		%s
	}`+"\n",
				ep.Name,
				WriteBodyDecodeur(ep.Name),
				DBCallerHandler(sgbd),
				PostQueryWriter(ep.Name, ep.Attribut, sgbd),
				QueryParamWriter(ep.Attribut),
				utils.WriteErrorCheker("insert error"),
				WriteResponseWriter())
}