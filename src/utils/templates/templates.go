package templatesutils

import "text/template"

func CliDocTemplate () *template.Template{
	content := `
Use these command : 
	- create_project : To create a new Project config
	- generate_project : To generate the code of the project 
	`
	temp := template.New("template")
	temp.Parse(content)

	return temp
}