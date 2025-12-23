package goapi

import (
	"bytes"
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"strings"
	"text/template"
)

func PutHandler(ep basetype.EndPoint, sgbd string) string {
	tmp, err := template.ParseFiles("module/go_api/templates/putHandler.gotmp")
	utils.ErrorChecker(err)
	var tmpBuffer bytes.Buffer
	data := struct {
		EndPointName    string
		DbCallerHandler string
		PutQuery        string
		ScanParams      string
		ResponseWriter  string
	}{
		EndPointName:    ep.Name,
		DbCallerHandler: DBCallerHandler(sgbd),
		PutQuery:        PutQueryWriter(ep.Name, ep.Attribut, sgbd),
		ScanParams:      ScanParamsWriter(ep),
		ResponseWriter:  strings.Replace(WriteResponseWriter(), "res", "tmp", 1),
	}
	err = tmp.Execute(&tmpBuffer, data)
	utils.ErrorChecker(err)
	return tmpBuffer.String()
}

func SelectByIdHandler(ep basetype.EndPoint, sgbd string) string {
	var tmpBuffer bytes.Buffer
	tmp, err := template.ParseFiles("module/go_api/templates/selectByIdHandler.gotmp")
	utils.ErrorChecker(err)
	data := struct {
		EndPointName    string
		DbCallerHandler string
		SelectByIdQuery string
		ScanParams      string
		ResponseWriter  string
	}{
		EndPointName:    ep.Name,
		DbCallerHandler: DBCallerHandler(sgbd),
		SelectByIdQuery: GetByIDQueryWriter(ep.Name, sgbd),
		ScanParams:      ScanParamsWriter(ep),
		ResponseWriter:  strings.Replace(WriteResponseWriter(), "res", "tmp", 1),
	}
	err = tmp.Execute(&tmpBuffer, data)
	utils.ErrorChecker(err)
	return tmpBuffer.String()
}

func SelectHandler(ep basetype.EndPoint, sgbd string) string {
	return fmt.Sprintf(`
	func %sHandleGetAll(w http.ResponseWriter, r *http.Request){
	%s
	var res []%sresponseType

	rows, err := db.Query("%s")

	for rows.Next(){
		var tmp %sresponseType
		rows.Scan(%s)
		res = append(res, tmp)
	}
	%s
	}
	`,
		ep.Name,
		DBCallerHandler(sgbd),
		ep.Name,
		GetQueryWriter(ep.Name, sgbd),
		ep.Name,
		ScanParamsWriter(ep),
		WriteResponseWriter())
}

func InsertHandler(ep basetype.EndPoint, sgbd string) string {
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

func DeleteHandler(ep basetype.EndPoint, sgbd string) string {
	return fmt.Sprintf(`func %sHandlerDelete(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	type response struct{
		Message string
	}
	%s
	rows,err := db.Query("%s", id)
	%s
	rows.Next()
	tmp := response{
		Message: "users deleted",
	}
	%s
	}
	`, ep.Name, DBCallerHandler(sgbd), DeleteQueryWriter(ep.Name, sgbd), utils.WriteErrorCheker("erreur lors du suppression"), WriteResponseWriter())
}
