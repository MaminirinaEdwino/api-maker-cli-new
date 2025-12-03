package utils

import (
	"bytes"
	"fmt"
	"text/template"
)

func ShowCliDocumentation(){
	var tmpBuffer bytes.Buffer
	tmp, err := template.ParseFiles("module/utils/templates/cliDoc.gotmp")
	ErrorChecker(err)
	err = tmp.Execute(&tmpBuffer, nil)
	ErrorChecker(err)
	fmt.Println(tmpBuffer.String())
}