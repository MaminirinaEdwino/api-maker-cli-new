package postgres

import (
	"fmt"
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
	"github/mameinirinaedwino/api-maker-cli/module/utils"
	"os"
)

func WriteTableColumns(columns []basetype.Attribut) string {
	script := ""
	for idx, attr := range columns {
		script += fmt.Sprintf(`%s %s%s`, attr.Nom, DatabaseTypeConverter(attr.Type), utils.AddComaOrNot(idx, len(columns)-1))
	}
	return script
}

func TableGeneratorPg(endPoint []basetype.EndPoint, file *os.File) {
	fmt.Println("Generating table creation script ...")
	file.WriteString("-- script de Generation de table --")
	for _, endpoint := range endPoint {
		fmt.Fprintf(file, `
CREATA TABLE %s(
id SERIAL PRIMARY KEY,
%s
);
`, endpoint.Name, WriteTableColumns(endpoint.Attribut))
	}

}