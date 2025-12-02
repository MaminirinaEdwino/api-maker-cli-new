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


func GetByIDQueryWriter(ep_name string, sgbd string) string {
	return GetQueryWriter(ep_name, sgbd) + " where id = $1"
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