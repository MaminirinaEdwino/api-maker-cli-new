package main

import (
	"flag"
	"fmt"
	"github/MaminirinaEdwino/api-maker-cli/src/apimaker"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"

	colortext "github.com/MaminirinaEdwino/colorText"
)

func main() {
	fmt.Println(colortext.GreenString("API maker"))

	createProject := flag.Bool("create_project", false, "The command that create a new project Config")
	generateProject := flag.Bool("generate_project", false, "Generate the project from the conf file")
	flag.Parse()

	switch {
	case *createProject:
		fmt.Println("Create a project")
		apimaker.CreateProject()
	case *generateProject:
		fmt.Println("Generate Project")
	default:
		utils.ShowCliDocumentation()
	}

}
