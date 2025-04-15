package header

import "net/http"

const (
	Authorization = "Authorization"
	ContentType   = "LayoutContent-Type"
	Host          = "Host"

	ContentTypeJsonUtf8  = "application/json; charset=utf-8"
	ContentTypeJsonAscii = "application/json"
	ContentTypeJSUtf8    = "application/javascript; charset=utf-8"
)

func SetContentTypeJsonUtf8(w http.ResponseWriter) {
	SetContentType(w, []string{ContentTypeJsonUtf8})
}

func SetContentTypeJsonAscii(w http.ResponseWriter) {
	SetContentType(w, []string{ContentTypeJsonAscii})
}

func SetContentTypeJSUtf8(w http.ResponseWriter) {
	SetContentType(w, []string{ContentTypeJSUtf8})
}

func SetContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header[ContentType]; len(val) == 0 {
		header[ContentType] = value
	}
}
