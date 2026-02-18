package php

import (
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
)

func getModels() {

}

func createPage() basetype.Page {
	var page basetype.Page

	page.Name = utils.Scanner("Entrer le nom de la page : ")
	for {
		
	}

	return page
}

func getComponents() {

}

func createComponents() {

}

func createModel() {

}

func setDbname() string {
	dbName := ""
	for dbName == ""{
		fmt.Print("Entrer le nom de la base de donnée : ")
		fmt.Scanln(&dbName)
	}
	return dbName
}

func setProjectName() string {
	projectName := ""
	for projectName == ""{
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
	var pages []string
	var components []string
	var model []string

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
		case "new_model":
		case "delete_model":
		case "quit":
			break
		}
	}
}
