package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"github.com/curtrube/snippetbox/assets"
	"github.com/curtrube/snippetbox/internal/models"
)

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
// to it as the project grows.
type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func newTemplateCache() (map[string]*template.Template, error) {
	// Initalize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Use the fs.Glob() to get a slice of filepaths in the assets.Files embedded
	// filesystem which match the pattern 'html/pages/*.tmpl'. This essentially
	// gives us a slice of all the 'page' templates for the application, just
	// like before.
	pages, err := fs.Glob(assets.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		//Extract the file name from the full filepath.
		name := filepath.Base(page)

		// Create a slice containing the filepath patterns for the templates we
		// want to parse.
		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		// Use the ParseFS() instead of ParseFiles() to parse the template files
		// from the assets.Files embedded filesystem.
		ts, err := template.New(name).Funcs(functions).ParseFS(assets.Files, patterns...)
		if err != nil {
			return nil, err
		}

		// Add the template set to the map, using the name of the page
		// (like 'home.tmpl') as the key.
		cache[name] = ts
	}

	return cache, nil
}

// humanDate returns a nicely formatted string representation of a time.Time.
func humanDate(t time.Time) string {
	// Return the empty string if the time has the zero value.
	if t.IsZero() {
		return ""
	}

	// Convert the time to UTC before formatting it.
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap value and store it in a global variable. This is
// essentially a string-keyed map which acts as a lookup table mapping names to
// functions.
var functions = template.FuncMap{
	"humanDate": humanDate,
}
