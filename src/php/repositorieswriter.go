package php

import (
	"fmt"
	basetype "github/MaminirinaEdwino/api-maker-cli/src/baseType"
	"github/MaminirinaEdwino/api-maker-cli/src/utils"
	"os"
)

func writeInsertFunction(file *os.File, model basetype.Model) {
	fmt.Fprintf(file, "public function insert(%s $%s){", UpperCaseFirstLetter(model.Name), model.Name)
	query := "$query = \"INSERT INTO %s (%s) VALUES(%s)\";"
	attrString := ""
	attrValue := ""
	for i, attr := range model.Attribut {
		attrString += fmt.Sprintf("%s", attr.Nom)
		attrValue += fmt.Sprintf("%d", i+1)
		if i < len(model.Attribut)-1 {
			attrString += ","
			attrValue += ","
		}
	}
	fmt.Fprintf(file, query, model.Name, attrString, attrValue)
	execString := `
		pg_prepare($this->database->conn, "", $query);
        pg_execute($this->database->conn, "", [
            %s
        ]);
	`
	attrExecString := ""
	for i, attr := range model.Attribut {
		attrExecString+= fmt.Sprintf("\"%s\" => $%s->get%s()", attr.Nom, attr.Nom, UpperCaseFirstLetter(attr.Nom))
		if i < len(model.Attribut) - 1  {
			attrExecString+=",\n"
		}else{
			attrExecString+="\n"
		}
	} 
	fmt.Fprintf(file, execString, attrExecString)
	file.WriteString("}\n")
		
}

func CreateRepositoriesWriter(ProjectName string, model basetype.Model) {
	file, err := os.Create(ProjectName + "/src/repositories/" + model.Name + "Repositories.php")
	utils.ErrorChecker(err)
	file.WriteString("<?php\n")
	file.WriteString("require_once \"./src/config/database.php\";\n")
	fmt.Fprintf(file, "require_once \"./src/models/%s.php\";\n", model.Name)
	fmt.Fprintf(file, "class %sRepositories {\n", UpperCaseFirstLetter(model.Name))
	file.WriteString("private Database $database;\n")
	file.WriteString(`
	public function __construct()
    {
        $this->database = new Database;
        $this->database->getConnection();
    }
	`)

}
