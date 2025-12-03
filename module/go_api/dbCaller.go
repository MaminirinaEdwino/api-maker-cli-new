package goapi

import (
	"bytes"
	"github/mameinirinaedwino/api-maker-cli/module/utils"
	"text/template"
)

func DbCallerPG() string {
	var tmpBuffer bytes.Buffer
	tmp, err := template.ParseFiles("module/go_api/templates/dbCaller.gotmp")
	utils.ErrorChecker(err)
	err = tmp.Execute(&tmpBuffer, nil)
	utils.ErrorChecker(err)
// 	return `
// db, err := sql.Open("postgres", database_url)
// if err != nil {
// 	log.Fatal(err)
// }
// defer db.Close()`
	return tmpBuffer.String()
}

func DBCallerHandler(sgbd string) string {
	if sgbd =="postgresql" {
		return DbCallerPG()
	}
	return ""
}