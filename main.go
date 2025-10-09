package main

import (
	"flag"
	"fmt"
	"os"
)

type Attribut struct {
	Nom  string
	Type string
}


/*
	RM_1 : Asina param ilay enddpoint specifique

*/


type EndPoint struct {
	Name      string
	Operation string
	Attribut  []Attribut
}

// Fonction qui scane le saisie user
func Scanner(label string) string {
	var tmp string
	fmt.Print(label)
	fmt.Scanln(&tmp)
	return tmp
}

func GetSGBD() string {
	var sgbd string
	fmt.Println("Choose your SGBD : ")
	fmt.Println("1. postgresql")
	fmt.Println("2. mysql")
	fmt.Scan(&sgbd)
	switch sgbd {
	case "postgresql":
		return sgbd
	case "mysql":
		return sgbd
	}
	return GetSGBD()
}

func GetEndPointType() string {
	typeEndPoint := Scanner("Choose the type of the endPoint \n" + "Endpoint that Interact with db or not (db / notdb) : ")
	switch typeEndPoint {
	case "db":
		return "db"
	case "notdb":
		return "notdb"
	}
	return GetEndPointType()
}

func GetEndPointOperation() string {
	operation := Scanner("Choose the operation for the endpoint : ")

	switch operation {
	case "create":
		return "create"
	case "read":
		return "read"
	case "update":
		return "update"
	case "delete":
		return "delete"
	case "crud":
		return "crud"
	}

	return GetEndPointOperation()
}

func GetAttrType() string {
	var attr_type string
	attr_type = Scanner("Choose the field type\n1. int\n2. string\n=> ")

	switch attr_type {
	case "string":
	case "int":
	default:
		GetAttrType()
	}
	return attr_type
}

func GetEndPointAttributes() []Attribut {
	var attr_list []Attribut
	var temp Attribut

	temp.Nom = Scanner("Enter a attribute Name : ")
	for temp.Nom != "" {
		temp.Type = GetAttrType()
		attr_list = append(attr_list, temp)
		temp.Nom = Scanner("Enter a attribute Name : ")
	}

	return attr_list
}

func GetEndPoints() ([]EndPoint, []EndPoint) {
	var endPointDb []EndPoint
	var endPointNoDb []EndPoint
	var typeEndPoint string
	var nameEndPoint string
	var operation string
	var attribute []Attribut

	nameEndPoint = Scanner("Enter endPoint name : ")
	fmt.Println(nameEndPoint)

	for nameEndPoint != "" {
		typeEndPoint = GetEndPointType()
		fmt.Println(typeEndPoint)
		operation = GetEndPointOperation()
		fmt.Println(operation)
		attribute = GetEndPointAttributes()
		fmt.Println(attribute)
		endPoint := EndPoint{
			Name:      nameEndPoint,
			Operation: operation,
			Attribut:  attribute,
		}
		if typeEndPoint == "db" {
			endPointDb = append(endPointDb, endPoint)
		} else {
			endPointNoDb = append(endPointNoDb, endPoint)
		}
		fmt.Println(endPoint)
		nameEndPoint = Scanner("Enter endPoint name : ")
	}

	//get endPoint
	return endPointDb, endPointNoDb

}

type Route struct{
	route string
	handler string
}

func GetQueryWriter(ep_name string,attrs []Attribut) string {
	

	return ""
}

func WriteCode(projectname string, sgbd string, db_name string, endPointDb []EndPoint, endPointNoDb []EndPoint) {
	var RouteList []Route
	project_dir := projectname
	fmt.Println("Start Writing the project code ")
	fmt.Println("Creating project folder . . .")
	os.Mkdir(project_dir, os.ModePerm)
	file, err := os.OpenFile(project_dir+"/main.go", os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)

	if err != nil {
		panic(err)
	}
	file.WriteString("package main\n\n")

	fmt.Println("Including all necessary package . . .")

	file.WriteString("import (\n")
	file.WriteString("\"fmt\"\n")
	file.WriteString("\"log\"\n")
	file.WriteString("\"net/http\"\n")
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
			database_url := fmt.Sprintf("const database_url = \"postgres://postgres:secret@localhost:5432/%s?sslmode=disable\"\n", db_name)
			file.WriteString(database_url)
		}
	}

	fmt.Println("Writing database migration code . . .")
	fmt.Println("writing all controller . . .")
	// wrting controller for endpoint db

	for _, ep := range endPointDb {
		if ep.Operation == "crud" {
			insertHandler := fmt.Sprintf(`
			func %sHandlePost(w http.ResponseWriter, r *http.Request){
				db, err := sql.Open("postgres", database_url)
				defer db.Close()



				if err != nil {
					log.Fatal(err)
				}
			}\n`, ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("POST /%s", ep.Name), handler: fmt.Sprintf("%sHandlePost", ep.Name)})

			selectHandler := fmt.Sprintf("func %sHandleGetAll(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s", ep.Name), handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})

			selectByIdHandler := fmt.Sprintf("func %sHandlerGetById(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})

			putHandler := fmt.Sprintf("func %sHandlerPut(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("PUT /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerPut", ep.Name)})

			deleteHandler := fmt.Sprintf("func %sHandlerDelete(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})

			file.WriteString(insertHandler)
			file.WriteString(selectHandler)
			file.WriteString(selectByIdHandler)
			file.WriteString(putHandler)
			file.WriteString(deleteHandler)
		}
		if ep.Operation == "create" {
			insertHandler := fmt.Sprintf("func %sHandlerPost(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("POST /%s", ep.Name), handler: fmt.Sprintf("%sHandlePost", ep.Name)})
			file.WriteString(insertHandler)
		}
		if ep.Operation == "read" {
			getAllHandler := fmt.Sprintf("func %sHandlerGetAll(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s", ep.Name), handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})
			getByIdHandler := fmt.Sprintf("func %sHandlerGetById(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})
			file.WriteString(getAllHandler)
			file.WriteString(getByIdHandler)
		}
		if ep.Operation == "update" {
			updateHandler := fmt.Sprintf("func %sHandlerPut(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("PUT /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerPut", ep.Name)})
			file.WriteString(updateHandler)
		}
		if ep.Operation == "delete" {
			deletehandler := fmt.Sprintf("func %sHandlerDelete(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})
			file.WriteString(deletehandler)
		}
	}

	for _, ep := range endPointNoDb {
		if ep.Operation == "crud" {
			insertHandler := fmt.Sprintf("func %sHandlePost(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("POST /%s", ep.Name), handler: fmt.Sprintf("%sHandlePost", ep.Name)})

			selectHandler := fmt.Sprintf("func %sHandleGetAll(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s", ep.Name), handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})

			selectByIdHandler := fmt.Sprintf("func %sHandlerGetById(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})

			putHandler := fmt.Sprintf("func %sHandlerPut(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("PUT /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerPut", ep.Name)})

			deleteHandler := fmt.Sprintf("func %sHandlerDelete(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})

			file.WriteString(insertHandler)
			file.WriteString(selectHandler)
			file.WriteString(selectByIdHandler)
			file.WriteString(putHandler)
			file.WriteString(deleteHandler)
		}
		if ep.Operation == "create" {
			insertHandler := fmt.Sprintf("func %sHandlerPost(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("POST /%s", ep.Name), handler: fmt.Sprintf("%sHandlePost", ep.Name)})
			file.WriteString(insertHandler)
		}
		if ep.Operation == "read" {
			getAllHandler := fmt.Sprintf("func %sHandlerGetAll(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s", ep.Name), handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})
			getByIdHandler := fmt.Sprintf("func %sHandlerGetById(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("GET /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})
			file.WriteString(getAllHandler)
			file.WriteString(getByIdHandler)
		}
		if ep.Operation == "update" {
			updateHandler := fmt.Sprintf("func %sHandlerPut(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("PUT /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerPut", ep.Name)})
			file.WriteString(updateHandler)
		}
		if ep.Operation == "delete" {
			deletehandler := fmt.Sprintf("func %sHandlerDelete(w http.ResponseWriter, r *http.Request){}\n", ep.Name)
			RouteList = append(RouteList, Route{route: fmt.Sprintf("DELETE /%s/{id}", ep.Name), handler: fmt.Sprintf("%sHandlerDelete", ep.Name)})
			file.WriteString(deletehandler)
		}
	}

	fmt.Println("Generating all routes . . .")
	file.WriteString("func Router(mux *http.ServeMux){\n")
	for _, route := range RouteList{
		fmt.Fprintf(file, "mux.HandleFunc(\"%s\", %s)\n", route.route, route.handler)
	}
	file.WriteString("}\n")


	fmt.Println("Writing the main server code . . ")
	file.WriteString("func main(){\nfmt.Println(\"API\")\nmux := http.NewServeMux()\nRouter(mux)\nfmt.Println(\"Server started at localhost:8000\")\nlog.Fatal(http.ListenAndServe(\":8000\", mux)))}\n")
	fmt.Println("Finished")
}

func CreateProjectConfig() {
	var projectname string
	var sgbd string
	var db_name string
	var EndPointDB []EndPoint
	var EndPointNotDB []EndPoint

	projectname = Scanner("Enter the name of your api : ")
	fmt.Println(projectname)
	sgbd = GetSGBD()
	fmt.Println("SGBD : " + sgbd)
	db_name = Scanner("Enter the name of the database :")
	fmt.Println(db_name)

	fmt.Println(len(EndPointDB))
	fmt.Println(len(EndPointNotDB))

	EndPointDB, EndPointNotDB = GetEndPoints()
	fmt.Println(EndPointDB)
	fmt.Println(EndPointNotDB)

	WriteCode(projectname, sgbd, db_name, EndPointDB, EndPointNotDB)
}

func main() {
	fmt.Println("API MAKER")

	createProject := flag.Bool("create_project", false, "The command that create a new project Config")
	generateProject := flag.Bool("generate_project", false, "Generate the project from the conf file")

	flag.Parse()

	switch {
	case *createProject:
		fmt.Println("Create a project")
		CreateProjectConfig()
	case *generateProject:
		fmt.Println("Generate Project")
	default:
		fmt.Println(`
Use these command : 
	- create_project : To create a new Project config
	- generate_project : To generate the code of the project 
		`)
	}

}
