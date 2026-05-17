package goapi

import (
	"bytes"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"text/template"
)

func DbCallerPG() string {
	var tmpBuffer bytes.Buffer
	tmp, err := template.ParseFiles("module/go_api/templates/dbCaller.gotmp")
	utils.ErrorChecker(err)
	err = tmp.Execute(&tmpBuffer, nil)
	utils.ErrorChecker(err)
	return tmpBuffer.String()
}

func DBCallerHandler(sgbd string) string {
	if sgbd == "postgresql" {
		return DbCallerPG()
	}
	return ""
}
