package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// check is exactly matches
	if r.URL.Path != "/" {
		app.notFound(w) // not found = 404 with notFound helper
		return
	}
	// paths
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	// parse files = to read template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // serverError helper
		return
	}
	// execute() =to write template content on response
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err)
	}

}
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	// extract id parameter from query string
	// Atoi = convert to an integer
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
