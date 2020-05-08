package tpl

import (
	"html/template"
	"io"

	"../errorhandle"
)

var tpl *template.Template

// SetUp ...
func SetUp() {
	var err error
	tpl, err = template.ParseGlob("./static/templates/*")
	errorhandle.Fatal(err)
}

// ExecuteTemplate ...
func ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return tpl.ExecuteTemplate(wr, name, data)
}
