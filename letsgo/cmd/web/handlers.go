package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)


type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	}
func (app *application)handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// order maters
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	// Use the template.ParseFiles() function to read the template file into a
	// template set. If there's an error, we log the detailed error message and use
	// the http.Error() function to send a generic 500 Internal Server Error
	// response to the user.
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
		return
	}
	// We then use the Execute() method on the template set to write the template
	// content as the response body. The last parameter to Execute() represents any
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Use the serverError() helper.
	}
}

// Add a showSnippet handler function.
func (app *application)showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // Use the notFound() helper.
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a createSnippet handler function.
func (app *application)createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		// w.WriteHeader(405)
		// w.Write([]byte("use post method"))
		app.clientError(w, http.StatusMethodNotAllowed) // Use the clientError() helper.
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"name":"Alex", "message":"Create a new snippet..."}`))
}
