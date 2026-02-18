package utils

import (
	"bytes"
	"fmt"
	templatesutils "github/MaminirinaEdwino/api-maker-cli/src/utils/templates"
)

func ShowCliDocumentation(){
	var tmpBuffer bytes.Buffer
	// tmp, err := template.ParseFiles ("module/utils/templates/cliDoc.gotmp")
	// ErrorChecker(err)
	err := templatesutils.CliDocTemplate().Execute(&tmpBuffer, nil)
	ErrorChecker(err)
	fmt.Println(tmpBuffer.String())
	
}