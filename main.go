package main

import (
	"flag"
	"fmt"
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
	"github/mameinirinaedwino/api-maker-cli/module/postgres"
	"os"
	"strings"

	colortext "github.com/MaminirinaEdwino/colorText"
	"github.com/rivo/tview"
)

// Fonction qui scanne le saisie user
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
	var attr_type string = Scanner("Choose the field type\n1. int\n2. string\n3. bool\n4. float\n=> ")
	switch attr_type {
	case "string":
	case "int":
	case "bool":
	case "float":
		return "float32"
	case "relation":
	default:
		GetAttrType()
	}
	return attr_type
}

func GetEndPointAttributes() []basetype.Attribut {
	var attr_list []basetype.Attribut
	var temp basetype.Attribut

	temp.Nom = Scanner("Enter a attribute Name : ")
	for temp.Nom != "" {
		temp.Type = GetAttrType()
		attr_list = append(attr_list, temp)
		temp.Nom = Scanner("Enter a attribute Name : ")
	}

	return attr_list
}

func GetEndPoints() ([]basetype.EndPoint, []basetype.EndPoint) {
	var endPointDb []basetype.EndPoint
	var endPointNoDb []basetype.EndPoint
	var typeEndPoint string
	var nameEndPoint string
	var operation string
	var attribute []basetype.Attribut

	nameEndPoint = Scanner("Enter endPoint name : ")
	fmt.Println(nameEndPoint)

	for nameEndPoint != "" {
		typeEndPoint = GetEndPointType()
		fmt.Println(typeEndPoint)
		operation = GetEndPointOperation()
		fmt.Println(operation)
		attribute = GetEndPointAttributes()
		fmt.Println(attribute)
		endPoint := basetype.EndPoint{
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



func GetQueryWriter(ep_name string, sgbd string) string {
	var query string
	if sgbd == "postgresql" {
		query = fmt.Sprintf("select * from %s ", ep_name)
	}
	if sgbd == "mysql" {
		query = fmt.Sprintf("select * from %s ", ep_name)
	}
	return query
}

func GetByIDQueryWriter(ep_name string, sgbd string) string {
	return GetQueryWriter(ep_name, sgbd) + " where id = $1"
}

func PostQueryWriter(ep_name string, attrs []basetype.Attribut, sgbd string) string {
	var query string
	attr_list := ""
	prepare_params := ""
	nbr_separator := len(attrs) - 1
	for idx, attr := range attrs {
		attr_list += attr.Nom
		prepare_params += fmt.Sprintf("$%d", idx)
		if nbr_separator > 0 {
			attr_list += " ,"
			prepare_params += " ,"
			nbr_separator--
		}
	}
	if sgbd == "postgresql" {
		query = postgres.Insert(ep_name, attr_list, prepare_params)
	}
	return query
}

func PutQueryWriter(ep_name string, attrs []basetype.Attribut, sgbd string) string {
	query := ""
	attr_list := ""
	nbr_separator := len(attrs) - 1
	for idx, attr := range attrs {
		attr_list += fmt.Sprintf("%s = $%d", attr.Nom, idx)
		if nbr_separator > 0 {
			attr_list += " ,"
			nbr_separator--
		}
	}
	if sgbd =="postgresql" {
		query = postgres.Update(ep_name, attr_list, attrs)
	}
	
	return query
}

func DeleteQueryWriter(ep_name string, sgbd string) string {
	if sgbd == "postgresql" {
		return postgres.Delete(ep_name)
	}
	return ""
}

func WriteBodyType(endPoint basetype.EndPoint) string {
	var AttrList string
	for _, attr := range endPoint.Attribut {
		AttrList += fmt.Sprintf("%s %s `json:\"%s\"`\n", strings.ToUpper(attr.Nom), attr.Type, attr.Nom)
	}
	return fmt.Sprintf(`
	type %sbodyType struct{
		//ID
		%s
	}
	`, endPoint.Name, AttrList)
}

func WriteResponseType(endPoint basetype.EndPoint) string {
	responseBody := strings.Replace(WriteBodyType(endPoint), "//ID", "ID int `json:\"id\"`", 1)
	responseBody = strings.Replace(responseBody, "bodyType", "responseType", 1)
	return responseBody
}

func WriteErrorCheker(message string) string {
	return fmt.Sprintf("if err != nil {\nfmt.Println(\"%s\", err)\n}", message)
}

func WriteBodyDecodeur(endPoint string) string {
	return fmt.Sprintf("var body %sbodyType \ndecoder := json.NewDecoder(r.Body) \nerr := decoder.Decode(&body)\n%s", endPoint, WriteErrorCheker("Parsing Error"))
}

func WriteResponseWriter() string {
	return `
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(res)
	`
}

func dbCaller() string {
	return `
db, err := sql.Open("postgres", database_url)
if err != nil {
	log.Fatal(err)
}
defer db.Close()`
}

func QueryParamWriter(attrs []basetype.Attribut) string {
	params := ""
	nbr_params := len(attrs) - 1
	for _, attr := range attrs {
		params += fmt.Sprintf("body.%s", strings.ToUpper(attr.Nom))
		if nbr_params > 0 {
			params += " ,"
			nbr_params--
		}
	}
	return params
}

func ScanParamsWriter(endPoint basetype.EndPoint) string {
	attr_list := ""
	nbr_params := len(endPoint.Attribut) - 1
	for _, attr := range endPoint.Attribut {
		attr_list += fmt.Sprintf("&tmp.%s", strings.ToUpper(attr.Nom))
		if nbr_params > 0 {
			attr_list += " ,"
			nbr_params--
		}
	}
	return attr_list
}

func WriteCode(projectname string, sgbd string, db_name string, endPointDb []basetype.EndPoint, endPointNoDb []basetype.EndPoint) {
	var RouteList []basetype.Route
	project_dir := projectname
	fmt.Println("Start Writing the project code ")
	fmt.Println("Creating project folder . . .")
	os.Mkdir(project_dir, os.ModePerm)
	file, err := os.OpenFile(project_dir+"/main.go", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
		if err != nil {
		panic(err)
	}
	tableScript, err := os.OpenFile(project_dir+"/database.sql", os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
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
				insertHandler := fmt.Sprintf(`
	func %sHandlePost(w http.ResponseWriter, r *http.Request){
		%s
		%s
	res, err := db.Exec("%s", %s)
		%s
		%s
	}`+"\n",
				ep.Name,
				WriteBodyDecodeur(ep.Name),
				dbCaller(),
				PostQueryWriter(ep.Name, ep.Attribut, sgbd),
				QueryParamWriter(ep.Attribut),
				WriteErrorCheker("insert error"),
				WriteResponseWriter())

				RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("POST /%s", ep.Name), Handler: fmt.Sprintf("%sHandlePost", ep.Name)})

				selectHandler := fmt.Sprintf(`
	func %sHandleGetAll(w http.ResponseWriter, r *http.Request){
	%s
	var res []%sresponseType

	rows, err := db.Query("%s")

	for rows.Next(){
		var tmp %sresponseType
		rows.Scan(%s)
		res = append(res, tmp)
	}
	%s
	}
	`,
		ep.Name,
		dbCaller(),
		ep.Name,
		GetQueryWriter(ep.Name, sgbd),
		ep.Name,
		ScanParamsWriter(ep),
		WriteResponseWriter())

				RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s", ep.Name), Handler: fmt.Sprintf("%sHandleGetAll", ep.Name)})

				selectByIdHandler := fmt.Sprintf(`func %sHandlerGetById(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	var tmp %sresponseType
	%s
	rows,err := db.Query("%s", id)
	rows.Next()
	rows.Scan(%s)
	%s
				}
				`, ep.Name, ep.Name, dbCaller(), GetByIDQueryWriter(ep.Name, sgbd), ScanParamsWriter(ep), strings.Replace(WriteResponseWriter(), "res", "tmp", 1))
				RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("GET /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerGetById", ep.Name)})

				putHandler := fmt.Sprintf(`func %sHandlerPut(w http.ResponseWriter, r *http.Request){
	var body %sbodyType
	var tmp %sResponseType
	id := r.PathValue("id")
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	%s
	rows, err := db.Query("%s", id)
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	rows.Scan(%s)
	%s
	}

	`, ep.Name, ep.Name, ep.Name, dbCaller(), PutQueryWriter(ep.Name, ep.Attribut, sgbd), ScanParamsWriter(ep),strings.Replace(WriteResponseWriter(), "res", "tmp", 1))
				RouteList = append(RouteList, basetype.Route{Route: fmt.Sprintf("PUT /%s/{id}", ep.Name), Handler: fmt.Sprintf("%sHandlerPut", ep.Name)})

				deleteHandler := fmt.Sprintf(`func %sHandlerDelete(w http.ResponseWriter, r *http.Request){
	id := r.PathValue("id")
	type response struct{
		Message string
	}
	%s
	rows,err := db.Query('%s', id)
	%s
	rows.Next()
	tmp := response{
		Message: "users deleted",
	}
	%s
	}
	`, ep.Name, dbCaller(), DeleteQueryWriter(ep.Name, sgbd), WriteErrorCheker("erreur lors du suppression"),WriteResponseWriter())
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

func CreateProjectConfig() {
	var projectname string
	var sgbd string
	var db_name string
	var EndPointDB []basetype.EndPoint
	var EndPointNotDB []basetype.EndPoint

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
	fmt.Println(colortext.GreenString("API maker"))

	createProject := flag.Bool("create_project", false, "The command that create a new project Config")
	generateProject := flag.Bool("generate_project", false, "Generate the project from the conf file")
	showBox := flag.Bool("showBox", false, "Show a teste")
	flag.Parse()

	switch {
	case *createProject:
		fmt.Println("Create a project")
		CreateProjectConfig()
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
