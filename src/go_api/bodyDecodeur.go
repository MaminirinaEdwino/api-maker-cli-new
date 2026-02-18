package goapi

import (
	"bytes"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"text/template"
)

func WriteBodyDecodeur(endPoint string) string {
	var tmpBuffer bytes.Buffer

	tmp, err := template.ParseFiles("module/go_api/templates/bodyDecodeur.gotmp")
	utils.ErrorChecker(err)

	data := struct {
		EndPointName string
		ErrorChecker string
	}{
		EndPointName: endPoint,
		ErrorChecker: utils.WriteErrorCheker("Parsing Error"),
	}

	err = tmp.Execute(&tmpBuffer, data)
	utils.ErrorChecker(err)
	// return fmt.Sprintf("var body %sbodyType \ndecoder := json.NewDecoder(r.Body) \nerr := decoder.Decode(&body)\n%s", endPoint, utils.WriteErrorCheker("Parsing Error"))
	return tmpBuffer.String()
}
