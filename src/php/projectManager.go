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
		}
	}

	return page
}

func createComponents() basetype.Component {
	var component basetype.Component

	component.Name = utils.Scanner("Nom du composant")
	contentType := ""
	for contentType == "" {
		contentType = utils.Scanner("Content Type : ")
		switch contentType {
		case "htmltag":
			tag := getHtmlTag()
			component.Content = append(component.Content, tag.GetHtmlTag())
		}
	}

	return component
}

func createModel() {

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
	var model []string
	var pages []basetype.Page
	var components []string

	fmt.Println("Create a PHP WEB App")
	action := ""
	for action != getAction() {
		switch action {
		case "project_name":
			projectName = setProjectName()
		case "db_name":
			Dbname = setDbname()
		case "new_component":
		case "delete_component":
		case "new_page":
			pages = append(pages, createPage())
		case "new_model":
		case "delete_model":
		case "quit":
			break
		}
	}
	fmt.Println(projectName)
	fmt.Println(Dbname)
	fmt.Println(model)
	fmt.Println(components)
}
