package main

import (
	"errors"
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
