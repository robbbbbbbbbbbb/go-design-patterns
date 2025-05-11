package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type templateData struct {
	Data map[string]any
}

func (app *application) render(w http.ResponseWriter, t string, td *templateData) {
	var tmpl *template.Template
	// use cache?
	if app.config.useCache {
		if fd, ok := app.templateMap[t]; ok {
			tmpl = fd
			log.Println("template cache has been used")
		}
	}
	if tmpl == nil {
		newTmpl, err := app.buildTemplateFromDisk(t)
		if err != nil {
			log.Println("error building template from disk:", err)
			return
		}
		log.Println("building template from disk")
		tmpl = newTmpl
	}

	if td == nil {
		td = &templateData{}
	}

	if err := tmpl.ExecuteTemplate(w, t, td); err != nil {
		log.Println("error executing template: ", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func (app *application) buildTemplateFromDisk(t string) (*template.Template, error) {
	tmplSlice := []string{
		"../../templates/base.layout.gohtml",
		"../../templates/partials/navbar.layout.gohtml",
		"../../templates/partials/header.layout.gohtml",
		"../../templates/partials/footer.layout.gohtml",
		fmt.Sprintf("../../templates/%s", t),
	}

	tmpl, err := template.ParseFiles(tmplSlice...)
	if err != nil {
		return nil, err
	}

	app.templateMap[t] = tmpl
	return tmpl, nil
}
