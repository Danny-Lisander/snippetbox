package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// sends generic 500 internal server error
// use debug stack to get stack trace for current goroutine
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// sends 400 "bad request"
// use http status text to automatically generate a human-friendly text
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// not found helper sends 404
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
