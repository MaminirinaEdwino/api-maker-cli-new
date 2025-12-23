package templates

import (
	"bytes"
	"text/template"
)

func BodyDecodeurTemplate(EndPointName string, ErrorChecker string) string {
	content := `
var body {{ .EndPointName }}bodyType 
decoder := json.NewDecoder(r.Body) 
err := decoder.Decode(&body)
{{ .ErrorChecker }}
`	
	var tmpBuffer bytes.Buffer
	template  := template.Template{}
	data := struct{
		EndPointName string
		ErrorCkecher string
	}{
		EndPointName: EndPointName,
		ErrorCkecher: ErrorChecker,
	}
	template.Parse(content)
	template.Execute(&tmpBuffer, data)
	return tmpBuffer.String()
}
