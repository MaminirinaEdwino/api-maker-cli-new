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

type Model struct {
	Name     string
	Attribut []Attribut
}

type HtmlTag struct {
	Type string
}

func (tag *HtmlTag) GetHtmlTag() string {
	tagString := ""
	switch tag.Type {
	case "div":
	case "h1":
	}
	return tagString
}

type Page struct {
	Name    string
	Content []string
}

type Component struct {
	Name string
	Content []string
}