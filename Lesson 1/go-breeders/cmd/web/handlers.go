package main

import (
	"fmt"
	"net/http"
)

func (app *application) ShowHome(w http.ResponseWriter, r *http.Request) {
	// values := r.URL.Query()
	var msg string = "Hello World!"
	if rmsg := r.URL.Query().Get("msg"); rmsg != "" {
		msg = rmsg
	}
	fmt.Fprint(w, msg)
}
