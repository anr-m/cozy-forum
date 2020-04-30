package tpl

import (
	"html/template"
	"io"
)

var tpl *template.Template

func SetUp() {
	tpl = template.Must(template.ParseGlob("./static/templates/*"))
}

func ExecuteTemplate(wr io.Writer, name string, data interface{}) {
	tpl.ExecuteTemplate(wr, name, data)
}
