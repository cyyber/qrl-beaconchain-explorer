package handlers

import (
	"net/http"

	"github.com/theQRL/zond-beaconchain-explorer/utils"
)

/**
base.go provides neat little helper functions for creating more abstract handlers.
By using these, you can write handlers that are able to handle both web and json api calls.

To migrate/write such handlers, replace X with Y:
	w.Header.Set()                     -> setAutoContentType()
	FormValue()	                       -> FormValueOrJSON()
	http.Error()                       -> ErrorOrJSONResponse()
	utils.SetFlash(); http.Redirect()  -> FlashRedirectOrJSONErrorResponse()
	http.Redirect()                    -> RedirectOrJSONOKResponse()
	w.WriteHeader(http.StatusOK)       -> OKResponse()
*/

// SetAutoContentType text/html for web OR application/json for API
func SetAutoContentType(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
}

// ErrorOrJSONResponse http.Error for web OR json error for API
func ErrorOrJSONResponse(w http.ResponseWriter, r *http.Request, errorText string, statusCode int) {
	http.Error(w, errorText, statusCode)
}

// FormValueOrJSON FormValue for web OR json value for API
func FormValueOrJSON(r *http.Request, key string) string {
	return r.FormValue(key)
}

// FlashRedirectOrJSONErrorResponse Set a flash message and redirect for web OR send an json error response with
// value as its data for API
func FlashRedirectOrJSONErrorResponse(w http.ResponseWriter, r *http.Request, name, value, url string, code int) {
	utils.SetFlash(w, r, name, value)
	http.Redirect(w, r, url, code)
}

// RedirectOrJSONOKResponse Redirect for web OR send an OK json response with empty data for API
func RedirectOrJSONOKResponse(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, url, code)
}

// OKResponse writeHeader(200) for web OR writeHeader(200) + empty json OK response
func OKResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
