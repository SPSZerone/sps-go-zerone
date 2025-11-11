package http

import (
	"fmt"
	"io"
	"net/http"

	"github.com/SPSZerone/sps-go-zerone/net/http/header"
)

type (
	OnResponse func(response *http.Response) ([]byte, error)
)

var (
	OnResponseDefault = func(response *http.Response) ([]byte, error) {
		if response.Body == nil {
			return nil, fmt.Errorf("response body is nil")
		}
		return io.ReadAll(response.Body)
	}
)

func HandleResponse(response *http.Response, onResponse OnResponse) ([]byte, error) {
	defer func() {
		err := response.Body.Close()
		if err != nil {
			fmt.Printf("http response body close failed. err: %v\n", err)
		}
	}()

	if onResponse != nil {
		return onResponse(response)
	}

	return OnResponseDefault(response)
}

func WriteJsonUtf8(w http.ResponseWriter, jsonBytes []byte) (int, error) {
	header.SetContentTypeJsonUtf8(w)
	return w.Write(jsonBytes)
}
