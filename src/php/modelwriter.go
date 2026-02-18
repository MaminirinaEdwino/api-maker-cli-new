package php

import (
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/postgres"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"os"
)



func CreateDBFile(dbname string, projectname string, model []basetype.Model) {
	file, err := os.Create(projectname+"/src/config/database.php")
	utils.ErrorChecker(err)
	createQuery := ""
	for m := range model {
		createQuery+=postgres.Create(model[m])
		if m < len(model) {
			createQuery+="\n"
		}
	}
	fmt.Fprintf(file, `
<?php
class Database{
    private string $host="your_host";
    private string $db_name="%s";
    private string $user = "postgres_user";
    private string $password  ="postgres_pwd";
    private string $port = "5432";
    public $conn;

    public function getConnection(){
        $this->conn = null;
        try {
            $this->conn = pg_connect("host = $this->host port = $this->port dbname = $this->db_name user = $this->user password = $this->password ");
        } catch (\Throwable $error) {
            echo "Connection Error ". $error->getMessage();
        }
    }

    public function initDatabaseStructure(){
        $this->getConnection();
        
        $query = "
        %s
        ";
        $result = pg_query($this->conn, $query);
    }
}
	`, dbname, createQuery)
}