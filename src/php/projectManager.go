package php

import (
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
)

func getModels() {

}

func getHtmlTag() basetype.HtmlTag {
	var tag basetype.HtmlTag
	tag.Type = ""
	for tag.Type == "" {
		tag.Type = utils.Scanner("HTML tag type : ")
	}
	return tag
}

func createPage() basetype.Page {
	var page basetype.Page

	page.Name = utils.Scanner("Entrer le nom de la page : ")
	contentType := ""
	for contentType == "" {
		contentType = utils.Scanner("Content type : ")
		switch contentType {
		case "htmltag":
			tag := getHtmlTag()
			page.Content = append(page.Content, tag.GetHtmlTag())
			contentType = ""
		}
	}

	return page
}

func createComponents() basetype.Component {
	var component basetype.Component
	component.Name = utils.Scanner("Nom du composant")
	contentType := "content"
	for contentType != "" {
		contentType = utils.Scanner("Content Type : ")
		switch contentType {
		case "htmltag":
			tag := getHtmlTag()
			component.Content = append(component.Content, tag.GetHtmlTag())
		}
	}

	return component
}

func createModel() basetype.Model {
	var model basetype.Model
	model.Name = utils.Scanner("Nom du model : ")
	attributName := "attr"
	for attributName != "" {
		attributName = utils.Scanner("Nom attribut : ")
		if attributName != "" {
			attributType := utils.Scanner("Type attribut : ")
			attr := basetype.Attribut{
				Nom:  attributName,
				Type: attributType,
			}
			model.Attribut = append(model.Attribut, attr)
		}
	}
	return model
}

func deleteModel(modelList []basetype.Model, modelName string) []basetype.Model {
	var tmp []basetype.Model
	for m := range modelList {
		if modelList[m].Name != modelName {
			tmp = append(tmp, modelList[m])
		}
	}
	return tmp
}

func setDbname() string {
	dbName := ""
	for dbName == "" {
		fmt.Print("Entrer le nom de la base de donnée : ")
		fmt.Scanln(&dbName)
	}
	return dbName
}

func setProjectName() string {
	projectName := ""
	for projectName == "" {
		fmt.Printf("Entrer le nom du projet : ")
		fmt.Scanln(&projectName)
	}
	return projectName
}

func getAction() string {
	cmd := ""
	for cmd == "" {
		fmt.Print("Entrer une commande : ")
		fmt.Scanln(&cmd)
	}
	return cmd
}

func CreateWebAppProject() {
	var projectName string
	var Dbname string
	var model []basetype.Model
	var pages []basetype.Page
	var components []basetype.Component

	fmt.Println("Create a PHP WEB App")
	action := ""
	for action == "" {
		action = getAction()
		switch action {
		case "project_name":
			projectName = setProjectName()
		case "db_name":
			Dbname = setDbname()
		case "new_component":
			components = append(components, createComponents())
		case "delete_component":
		case "new_page":
			pages = append(pages, createPage())
		case "new_model":
			model = append(model, createModel())
		case "delete_model":
		}
	}
	fmt.Println(projectName)
	fmt.Println(Dbname)
	fmt.Println(model)
	fmt.Println(components)
}
