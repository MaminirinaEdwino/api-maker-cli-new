package main

import (
"fmt"
"log"
"net/http"
"encoding/json"
"database/sql"
// _ "github.com/lib/pq"
)
const database_url = "postgres://postgres:secret@localhost:5432/testepg?sslmode=disable"

	type usersbodyType struct{
		//ID
		NAME string `json:"name"`
AGE int `json:"age"`

	}
	
	type usersresponseType struct{
		ID int `json:"id"`
		NAME string `json:"name"`
AGE int `json:"age"`

	}
	

func usersHandlePost(w http.ResponseWriter, r *http.Request){
	var body usersbodyType 
decoder := json.NewDecoder(r.Body) 
err := decoder.Decode(&body)
if err != nil {
fmt.Println("Parsing Error", err)
}

	
db, err := sql.Open("postgres", database_url)
if err != nil {
	log.Fatal(err)
}
defer db.Close()
	res, err := db.Exec("insert into users (name ,age) values ($0 ,$1) returning *" , body.NAME ,body.AGE)
	if err != nil {
fmt.Println("insert error", err)
}
	
	
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(res)
	
}

func usersHandleGetAll(w http.ResponseWriter, r *http.Request){

db, err := sql.Open("postgres", database_url)
if err != nil {
	log.Fatal(err)
}
defer db.Close()			
var res []usersresponseType
rows, err := db.Query("select * from users" )
for rows.Next(){
	var tmp usersresponseType
	rows.Scan(&tmp.NAME ,&tmp.AGE)
	res = append(res, tmp)
}

w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusOK)
json.NewEncoder(w).Encode(res)
	
}
func usersHandlerGetById(w http.ResponseWriter, r *http.Request){}
func usersHandlerPut(w http.ResponseWriter, r *http.Request){}
func usersHandlerDelete(w http.ResponseWriter, r *http.Request){}
func Router(mux *http.ServeMux){
mux.HandleFunc("POST /users", usersHandlePost)
mux.HandleFunc("GET /users", usersHandleGetAll)
mux.HandleFunc("GET /users/{id}", usersHandlerGetById)
mux.HandleFunc("PUT /users/{id}", usersHandlerPut)
mux.HandleFunc("DELETE /users/{id}", usersHandlerDelete)
}
func main(){
fmt.Println("API")
mux := http.NewServeMux()
Router(mux)
fmt.Println("Server started at localhost:8000")
log.Fatal(http.ListenAndServe(":8000", mux))}
package main

package main

package main

package main

