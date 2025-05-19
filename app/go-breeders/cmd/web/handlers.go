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
	page := chi.URLParam(r, "page")
	err := app.render(w, page+".page.gohtml", nil)
	if err != nil {
		app.render(w, "error.page.gohtml", nil)
	}
}

func (app *application) TestPatterns(w http.ResponseWriter, r *http.Request) {
	app.render(w, "test.page.gohtml", nil)
}

func (app *application) NewDogFromFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	np := pets.New("Dog")

	tools.WriteJSON(w, http.StatusOK, np)
}

func (app *application) NewCatFromFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	np := pets.New("Cat")

	tools.WriteJSON(w, http.StatusOK, np)
}

func (app *application) NewDogFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	tools.WriteJSON(w, http.StatusOK, dog)
}

func (app *application) NewCatFromAbstractFactory(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
	tools.WriteJSON(w, http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(w http.ResponseWriter, r *http.Request) {
	var tools toolbox.Tools
	dogBreeds, err := app.Models.DogBreed.All()
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	err = tools.WriteJSON(w, http.StatusOK, dogBreeds)
	if err != nil {
		tools.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}
}
