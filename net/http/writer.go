package http

import (
	"net/http"

	"github.com/SPSZerone/sps-go-zerone/net/http/header"
)

func WriteJsonUtf8(w http.ResponseWriter, jsonBytes []byte) (int, error) {
	header.SetContentTypeJsonUtf8(w)
	return w.Write(jsonBytes)
}
