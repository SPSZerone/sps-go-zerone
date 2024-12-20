package http

import (
	"io"
	"net/http"
	"net/url"

	"github.com/SPSZerone/sps-go-zerone/net/http/header"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

func PostFormDefault(url string, data url.Values) ([]byte, error) {
	return PostForm(url, data, nil)
}

func PostForm(url string, data url.Values, onResponse OnResponse) ([]byte, error) {
	response, err := http.PostForm(url, data)
	if err != nil {
		return nil, err
	}
	return HandleResponse(response, onResponse)
}

func PostJsonUtf8Request(url string, body io.Reader, opts ...RequestOption) ([]byte, error) {
	opts = append(opts, ReqOptHeaderAdd(header.ContentType, header.ContentTypeJsonUtf8))
	return Request(MethodPost, url, body, nil, opts...)
}

func PostRequest(url string, body io.Reader, opts ...RequestOption) ([]byte, error) {
	return Request(MethodPost, url, body, nil, opts...)
}

func Request(method, url string, body io.Reader, onResponse OnResponse, opts ...RequestOption) ([]byte, error) {
	reqOpts := NewRequestOptions(method, url, body, onResponse, opts...)
	return DoRequest(reqOpts)
}

func DoRequest(opts *RequestOptions) ([]byte, error) {
	request, err := http.NewRequest(opts.Method, opts.Url, opts.Body)
	if err != nil {
		return nil, err
	}
	for key, value := range opts.Header {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return HandleResponse(response, opts.OnResponse)
}
