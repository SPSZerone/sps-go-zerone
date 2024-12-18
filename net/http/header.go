package http

import "net/http"

const (
	HeaderContentType = "Content-Type"

	HeaderContentTypeJsonUtf8   = "application/json; charset=utf-8"
	HeaderContentTypeJsonUtf8JS = "application/javascript; charset=utf-8"
	HeaderContentTypeJsonAscii  = "application/json"
)

func HeaderSetContentTypeJsonUtf8(w http.ResponseWriter) {
	HeaderSetContentType(w, []string{HeaderContentTypeJsonUtf8})
}
func HeaderSetContentTypeJsonUtf8JS(w http.ResponseWriter) {
	HeaderSetContentType(w, []string{HeaderContentTypeJsonUtf8JS})
}
func HeaderSetContentTypeJsonAscii(w http.ResponseWriter) {
	HeaderSetContentType(w, []string{HeaderContentTypeJsonAscii})
}

func HeaderSetContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header[HeaderContentType]; len(val) == 0 {
		header[HeaderContentType] = value
	}
}
