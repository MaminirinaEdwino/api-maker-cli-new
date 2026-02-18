package apimaker

import (
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	goapi "github/MaminirinaEdwino/api-maker-cli/src/go_api"
	"github/MaminirinaEdwino/api-maker-cli/src/postgres"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"os"
)

func WriteCode(
	projectname string,
	sgbd string,
	db_name string,
	endPointDb []basetype.EndPoint,
	endPointNoDb []basetype.EndPoint,
) {

	var RouteList []basetype.Route
	project_dir := projectname
	fmt.Println("Start Writing the project code ")
	fmt.Println("Creating project folder . . .")
	os.Mkdir(project_dir, os.ModePerm)
	file, err := os.OpenFile(project_dir+"/main.go", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)

	utils.ErrorChecker(err)

	tableScript, err := os.OpenFile(project_dir+"/database.sql", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	utils.ErrorChecker(err)

	file.WriteString("package main\n\n")

	fmt.Println("Including all necessary package . . .")

	postgres.TableGeneratorPg(endPointDb, tableScript)

	file.WriteString("import (\n")
	file.WriteString("\"fmt\"\n")
	file.WriteString("\"log\"\n")
	file.WriteString("\"net/http\"\n")
	file.WriteString("\"encoding/json\"\n")
	if len(endPointDb) > 0 {
		file.WriteString("\"database/sql\"\n")

		if sgbd == "postgresql" {
			file.WriteString("_ \"github.com/lib/pq\"\n")
		}
	}

	file.WriteString(")\n")

	// writing database simple config
	if len(endPointDb) > 0 {
		if sgbd == "postgresql" {
			database_url := postgres.DataBaseNameGo(db_name)
			file.WriteString(database_url)
		}
	}

	fmt.Println("Writing database migration code . . .")
	fmt.Println("writing all controller . . .")
	// wrting controller for endpoint db

	for _, ep := range endPointDb {
		if ep.Operation == "crud" {
			file.WriteString(WriteBodyType(ep))
			file.WriteString(WriteResponseType(ep))

			insertHandler := goapi.InsertHandler(ep, sgbd)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("POST /%s", ep.Name), Handler: fmt.Sprintf("%sHandlePost", ep.Name)})

			selectHandler := goapi.SelectHandler(ep, sgbd)

			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s", ep.Name), Handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})

			selectByIdHandler := goapi.SelectByIdHandler(ep, sgbd)

			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})

			putHandler := goapi.PutHandler(ep, sgbd)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("PUT /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerPut", ep.Name)})

			deleteHandler := goapi.DeleteHandler(ep, sgbd)

			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})

			file.WriteString(insertHandler)
			file.WriteString(selectHandler)
			file.WriteString(selectByIdHandler)
			file.WriteString(putHandler)
			file.WriteString(deleteHandler)
		}
		if ep.Operation == "create" {
			insertHandler := fmt.Sprintf("func %sHandlerPost(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("POST /%s", ep.Name), Handler: fmt.Sprintf("%sHandlePost", ep.Name)})
			file.WriteString(insertHandler)
		}
		if ep.Operation == "read" {
			getAllHandler := fmt.Sprintf("func %sHandlerGetAll(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s", ep.Name), Handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})
			getByIdHandler := fmt.Sprintf("func %sHandlerGetById(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})
			file.WriteString(getAllHandler)
			file.WriteString(getByIdHandler)
		}
		if ep.Operation == "update" {
			updateHandler := fmt.Sprintf("func %sHandlerPut(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("PUT /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerPut", ep.Name)})
			file.WriteString(updateHandler)
		}
		if ep.Operation == "delete" {
			deletehandler := fmt.Sprintf("func %sHandlerDelete(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})
			file.WriteString(deletehandler)
		}
	}

	for _, ep := range endPointNoDb {
		if ep.Operation == "crud" {
			insertHandler := fmt.Sprintf("func %sHandlePost(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("POST /%s", ep.Name), Handler: fmt.Sprintf("%sHandlePost", ep.Name)})

			selectHandler := fmt.Sprintf("func %sHandleGetAll(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s", ep.Name), Handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})

			selectByIdHandler := fmt.Sprintf("func %sHandlerGetById(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})

			putHandler := fmt.Sprintf("func %sHandlerPut(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("PUT /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerPut", ep.Name)})

			deleteHandler := fmt.Sprintf("func %sHandlerDelete(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})

			file.WriteString(insertHandler)
			file.WriteString(selectHandler)
			file.WriteString(selectByIdHandler)
			file.WriteString(putHandler)
			file.WriteString(deleteHandler)
		}
		if ep.Operation == "create" {
			insertHandler := fmt.Sprintf("func %sHandlerPost(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("POST /%s", ep.Name), Handler: fmt.Sprintf("%sHandlePost", ep.Name)})
			file.WriteString(insertHandler)
		}
		if ep.Operation == "read" {
			getAllHandler := fmt.Sprintf("func %sHandlerGetAll(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s", ep.Name), Handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})
			getByIdHandler := fmt.Sprintf("func %sHandlerGetById(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})
			file.WriteString(getAllHandler)
			file.WriteString(getByIdHandler)
		}
		if ep.Operation == "update" {
			updateHandler := fmt.Sprintf("func %sHandlerPut(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("PUT /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerPut", ep.Name)})
			file.WriteString(updateHandler)
		}
		if ep.Operation == "delete" {
			deletehandler := fmt.Sprintf("func %sHandlerDelete(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})
			file.WriteString(deletehandler)
		}
	}

	fmt.Println("Generating all routes . . .")
	file.WriteString("func Router(mux *http.ServeMux){\n")
	for _, route := range RouteList {
		fmt.Fprintf(file, "mux.HandleFunc(\"%s\", %s)\n", route.Route, route.Handler)
	}
	file.WriteString("}\n")

	fmt.Println("Writing the main server code . . ")
	file.WriteString("func main(){\nfmt.Println(\"API\")\nmux := http.NewServeMux()\nRouter(mux)\nfmt.Println(\"Server started at localhost:8000\")\nlog.Fatal(http.ListenAndServe(\":8000\", mux))}\n")
	fmt.Println("Finished")
}
