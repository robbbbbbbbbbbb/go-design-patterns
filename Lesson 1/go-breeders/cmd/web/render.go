package main

import (
	"html/template"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template

	// // use cache?
	// if app.config.useCache {
	// 	if fd, ok := app.templateMap["template"]; ok {

	// 	}
	// }
}
