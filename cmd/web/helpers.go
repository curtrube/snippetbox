package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/form/v4"
)

// serverError helper writes a log entry at Error level (including the request
// method and URI as attributes), then sends a generic 500 Internal Server Error
// response to the user.
func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError helper sends a specific status code and corresponding description
// to the user.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// newTemplateData helper creates a new templateData initialized with
// the current year.
func (app *application) newTemplateData(r *http.Request) templateData {
	return templateData{
		CurrentYear: time.Now().Year(),
		Flash:       app.sessionManager.PopString(r.Context(), "flash"),
	}
}

func (app *application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) error {
	// Retrieve the appropriate template set from the cache based on the page
	// name (like 'home.tmpl'). If no entry exists in the cache with the
	// provided name, then create a new error and call the serverError() helper
	// method that we made earlier and return.
	ts, ok := app.templateCache[page]
	if !ok {
		return fmt.Errorf("the template %s does not exist in the cache", page)
	}

	// Initalize a new buffer.
	buf := new(bytes.Buffer)

	// Write the template to the buffer, instead of straight to the
	// http.ResponseWriter. If there's an error, return it.
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return err
	}

	// If the template is written to the buffer without any errors, it's safe
	// to go ahead and write the HTTP status code to the http.ResponseWriter.
	w.WriteHeader(status)

	// Write the contents of the buffer to the http.ResponseWriter, returning any
	// error. Note: this is another time where we pass our http.ResponseWriter
	// to a function that takes an io.Writer.
	_, err = buf.WriteTo(w)
	return err
}

// decodePostForm helper calls r.ParseForm() on the current request. Next,
// calls app.formDecoder.Decode() to unpack the HTML form data to a target
// destination, Finally, checks fora  form.InvalidDecodeError error and
// triggers a panic if we ever see it.
func (app *application) decodePostForm(r *http.Request, dst any) error {
	// Call ParseForm() on the request, in the same way that we did in our
	// snippetCreatePost handler.
	err := r.ParseForm()
	if err != nil {
		return err
	}

	// Call Decode() on our decoder instance, passing the target destination as
	// the first parameter.
	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {
		// If er try to use an invalid target destination, the Decode() method
		// will return an error with the type form.InvalidDecodeError. We use
		// errors.AsType() to check for this and panic. At the end of this
		// chapter we'll talk about panicking versus returning errors, and
		// discuss why it's an appropriate thing to do in this specific situation.
		if _, ok := errors.AsType[*form.InvalidDecoderError](err); ok {
			panic(err)
		}

		// For all other errors, return them as normal.
		return err
	}

	return nil
}
