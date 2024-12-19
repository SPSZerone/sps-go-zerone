package http

import (
	"fmt"
	"io"
	"net/http"
)

const (
	MethodGet  = "GET"
	MethodPost = "POST"
)

func DoRequest(r *Request) error {
	request, err := http.NewRequest(r.Method, r.Url, r.Body)
	if err != nil {
		return err
	}
	for key, value := range r.Header {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {

		}
	}()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("StatusCode:%d", response.StatusCode)
	}

	if r.OnSuccess != nil {
		r.OnSuccess(response)
	}

	return nil
}

func NewRequest(method, url string, body io.Reader, onSuccess func(*http.Response), opts ...RequestOption) *Request {
	r := &Request{
		Method:    method,
		Url:       url,
		Body:      body,
		OnSuccess: onSuccess,
	}
	r.Update(opts...)
	return r
}

type Request struct {
	Method    string
	Url       string
	Body      io.Reader
	Header    map[string]string
	OnSuccess func(response *http.Response)
}

func (r *Request) Update(opts ...RequestOption) {
	for _, opt := range opts {
		opt(r)
	}
}

func (r *Request) HeaderAdd(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}

type RequestOption func(*Request)

func ReqOptMethod(value string) RequestOption {
	return func(r *Request) {
		r.Method = value
	}
}

func ReqOptUrl(value string) RequestOption {
	return func(r *Request) {
		r.Url = value
	}
}

func ReqOptBody(value io.Reader) RequestOption {
	return func(r *Request) {
		r.Body = value
	}
}

func ReqOptHeader(value map[string]string) RequestOption {
	return func(r *Request) {
		r.Header = value
	}
}

func ReqOptHeaderAdd(key, value string) RequestOption {
	return func(r *Request) {
		r.HeaderAdd(key, value)
	}
}

func ReqOptOnSuccess(onSuccess func(response *http.Response)) RequestOption {
	return func(r *Request) {
		r.OnSuccess = onSuccess
	}
}
