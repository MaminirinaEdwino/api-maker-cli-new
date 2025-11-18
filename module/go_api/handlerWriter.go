package goapi

import (
	"fmt"
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
	"github/mameinirinaedwino/api-maker-cli/module/utils"
	"strings"
)

func PutHandler(ep basetype.EndPoint, sgbd string)string{
	return fmt.Sprintf(`func %sHandlerPut(w http.ResponseWriter, r *http.Request){
	var body %sbodyType
	var tmp %sResponseType
	id := r.PathValue("id")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	%s
	rows, err := db.Query("%s", id)
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	rows.Scan(%s)
	%s
	}

	`, ep.Name, ep.Name, ep.Name, DBCallerHandler(sgbd), PutQueryWriter(ep.Name, ep.Attribut, sgbd), ScanParamsWriter(ep),strings.Replace(WriteResponseWriter(), "res", "tmp", 1))
}

func SelectByIdHandler(ep basetype.EndPoint, sgbd string)string {
	return fmt.Sprintf(`func %sHandlerGetById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	var tmp %sresponseType
	%s
	rows,err := db.Query("%s", id)
	rows.Next()
	rows.Scan(%s)
	%s
				}
				`, ep.Name, ep.Name, DBCallerHandler(sgbd), GetByIDQueryWriter(ep.Name, sgbd), ScanParamsWriter(ep), strings.Replace(WriteResponseWriter(), "res", "tmp", 1))
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
