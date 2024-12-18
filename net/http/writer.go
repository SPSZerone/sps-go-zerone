package http

import (
	"net/http"
)

func WriteJsonUtf8(w http.ResponseWriter, jsonBytes []byte) (int, error) {
	HeaderSetContentTypeJsonUtf8(w)
	return w.Write(jsonBytes)
}
