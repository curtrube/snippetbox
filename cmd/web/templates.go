package main

import "github.com/curtrube/snippetbox/internal/models"

// Include a Snippets fields in the templateData struct.
type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
