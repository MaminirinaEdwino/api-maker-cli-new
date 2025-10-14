package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const database_url = "postgres://postgres:root@localhost:5432/testepg?sslmode=disable"

type usersbodyType struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func usersHandlePost(w http.ResponseWriter, r *http.Request) {

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

	res, err := db.Exec("insert into users (name, age) values($1, $2) returning * ", body.Name, body.Age)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}

type userResponseType struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
	AGE  int    `json:"age"`
}

func usersHandleGetAll(w http.ResponseWriter, r *http.Request) {
	var res []userResponseType
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from users")

	for rows.Next() {
		var tmp userResponseType
		rows.Scan(&tmp.ID, &tmp.NAME, &tmp.AGE)
		res = append(res, tmp)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
func usersHandlerGetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var tmp userResponseType
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select * from users where id  = $1", id)
	rows.Next()
	rows.Scan(&tmp.ID, &tmp.NAME, &tmp.AGE)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmp)
}
func usersHandlerPut(w http.ResponseWriter, r *http.Request) {
	var body usersbodyType
	var tmp userResponseType
	id := r.PathValue("id")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&body)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", database_url)
	if err != nil  {
		log.Fatal(err)
	}
	rows, err := db.Query("update users set name = $1, age = $2 where id = $3 returning * ", body.Name, body.Age, id)
	if err != nil {
		log.Fatal(err)
	}
	rows.Next()
	rows.Scan(&tmp.ID, &tmp.NAME, &tmp.AGE)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tmp)
}
func usersHandlerDelete(w http.ResponseWriter, r *http.Request) {}
func Router(mux *http.ServeMux) {
	mux.HandleFunc("POST /users", usersHandlePost)
	mux.HandleFunc("GET /users", usersHandleGetAll)
	mux.HandleFunc("GET /users/{id}", usersHandlerGetById)
	mux.HandleFunc("PUT /users/{id}", usersHandlerPut)
	mux.HandleFunc("DELETE /users/{id}", usersHandlerDelete)
}
func main() {
	fmt.Println("API")
	mux := http.NewServeMux()
	Router(mux)
	fmt.Println("Server started at localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
