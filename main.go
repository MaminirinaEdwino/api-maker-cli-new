package main

import (
	"flag"
	"fmt"
	"github/mameinirinaedwino/api-maker-cli/module/apimaker"
	colortext "github.com/MaminirinaEdwino/colorText"
	"github.com/rivo/tview"
)

func main() {
	fmt.Println(colortext.GreenString("API maker"))

	createProject := flag.Bool("create_project", false, "The command that create a new project Config")
	generateProject := flag.Bool("generate_project", false, "Generate the project from the conf file")
	showBox := flag.Bool("showBox", false, "Show a teste")
	flag.Parse()

	switch {
	case *createProject:
		fmt.Println("Create a project")
		apimaker.CreateProject()
	case *generateProject:
		fmt.Println("Generate Project")
	case *showBox:
		box := tview.NewBox().SetBorder(true).SetTitle("Hello world")
		if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
			panic(err)
		}
	default:
		fmt.Println(`
Use these command : 
	- create_project : To create a new Project config
	- generate_project : To generate the code of the project 
		`)
	}

}
