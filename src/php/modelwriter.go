package php

import (
	"strings"
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/postgres"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"os"
)

func CreateDBFile(dbname string, projectname string, model []basetype.Model) {
	file, err := os.Create(projectname+"/src/config/database.php")
	utils.ErrorChecker(err)
	var createQuery strings.Builder
	for m := range model {
		createQuery.WriteString(postgres.Create(model[m]))
		if m < len(model) {
			createQuery.WriteString("\n")
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
	`, dbname, createQuery.String())
}

func CreateMigrateDBFile(projectName string){
	file, err := os.Create(projectName+"/src/config/migrationDatabase.php")
	utils.ErrorChecker(err)
	fmt.Fprint(file, `
<?php

require "./database.php";
$database = new Database();
$database->getConnection();
$database->initDatabaseStructure();
	`)
}

func UpperCaseFirstLetter(chain string) string {
	tab := strings.Split(chain, "")
	tab[0] = strings.ToUpper(tab[0])
	return strings.Join(tab, "")
}

func CreateModelFile(projectName string, model basetype.Model){
	file, err := os.Create(projectName+"/src/models/"+model.Name+".php")
	utils.ErrorChecker(err)
	file.WriteString("<?php\n")
	file.WriteString(fmt.Sprintf("class %s {\n", UpperCaseFirstLetter(model.Name)))

	for _,attr := range model.Attribut {
		fmt.Fprintf(file, "private %s $%s;\n", attr.Type ,attr.Nom )
	}

	fmt.Fprint(file, "public function __construct(){}")
	for _, attr := range model.Attribut {
		fmt.Fprintf(file, `
	public function get%s(){
		return $this->%s; 
	}
	public function set%s($%s){
		$this->%s = $%s;
		return $this;
	}
	`, UpperCaseFirstLetter(attr.Nom), attr.Nom, UpperCaseFirstLetter(attr.Nom), attr.Nom, attr.Nom, attr.Nom)
	}
	file.WriteString("\n}\n")
}