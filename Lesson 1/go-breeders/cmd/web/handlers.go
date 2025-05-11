package main

import (
	"fmt"
	"go-breeders/pets"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tsawler/toolbox"
)

func (app *application) Health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "home.page.gohtml", nil)
}

func (app *application) ShowPage(w http.ResponseWriter, r *http.Request) {
	factoryRoutes := map[string]int{
		"new-dog": 1,
		"new-cat": 1,
	}
	page := chi.URLParam(r, "page")
	if _, ok := factoryRoutes[page]; !ok {
		app.render(w, page+".page.gohtml", nil)
	}
}

func (app *application) NewDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	np := pets.New("Dog")

	tools.WriteJSON(w, http.StatusOK, np)

	// jsonData, err := json.Marshal(np)
	// if err != nil {
	// 	http.Error(w, "cannot marshal data", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// _, err = w.Write(jsonData)
	// if err != nil {
	// 	http.Error(w, "cannot write data", http.StatusInternalServerError)
	// 	return
	// }
}

func (app *application) NewCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	np := pets.New("Cat")

	tools.WriteJSON(w, http.StatusOK, np)

	// jsonData, err := json.Marshal(np)
	// if err != nil {
	// 	http.Error(w, "cannot marshal data", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// _, err = w.Write(jsonData)
	// if err != nil {
	// 	http.Error(w, "cannot write data", http.StatusInternalServerError)
	// 	return
	// }
}
