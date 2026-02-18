package basetype


type Attribut struct {
	Nom  string
	Type string
}

type EndPoint struct {
	Name      string
	Operation string
	Attribut  []Attribut
}

type Route struct {
	Route   string
	Handler string
}

type Model struct{
	Name string
	Attribut []Attribut
}

type Page struct{
	
}