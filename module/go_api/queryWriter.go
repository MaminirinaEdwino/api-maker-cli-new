package goapi

import (
	"fmt"
	basetype "github/mameinirinaedwino/api-maker-cli/module/baseType"
	"github/mameinirinaedwino/api-maker-cli/module/postgres"
	"strings"
)

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