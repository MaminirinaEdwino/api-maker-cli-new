package postgres

import (
	"fmt"
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
)

func Insert(ep_name string, attr_list string, prepare_params string) string{
	return fmt.Sprintf("insert into %s (%s) values (%s) returning * ", ep_name, attr_list, prepare_params)
}

func Update(ep_name string, attr_list string, attrs []basetype.Attribut) string {
	return fmt.Sprintf("update %s set %s where id = $%d returning * ", ep_name, attr_list, len(attrs))
}

func Delete(ep_name string) string {
	return fmt.Sprintf("delete from %s where id = $1", ep_name)
}