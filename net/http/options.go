package http

import "io"

func NewRequestOptions(method, url string, body io.Reader, onResponse OnResponse, opts ...RequestOption) *RequestOptions {
	r := &RequestOptions{
		Method:     method,
		Url:        url,
		Body:       body,
		OnResponse: onResponse,
	}
	r.Update(opts...)
	return r
}

type RequestOptions struct {
	Method     string
	Url        string
	Body       io.Reader
	Header     map[string]string
	OnResponse OnResponse
}

func (r *RequestOptions) Update(opts ...RequestOption) {
	for _, opt := range opts {
		opt(r)
	}
}

func (r *RequestOptions) HeaderAdd(key, value string) {
	if r.Header == nil {
		r.Header = make(map[string]string)
	}
	r.Header[key] = value
}

type RequestOption func(*RequestOptions)

func ReqOptMethod(value string) RequestOption {
	return func(r *RequestOptions) {
		r.Method = value
	}
}

func ReqOptUrl(value string) RequestOption {
	return func(r *RequestOptions) {
		r.Url = value
	}
}

func ReqOptBody(value io.Reader) RequestOption {
	return func(r *RequestOptions) {
		r.Body = value
	}
}

func ReqOptHeader(value map[string]string) RequestOption {
	return func(r *RequestOptions) {
		r.Header = value
	}
}

func ReqOptHeaderAdd(key, value string) RequestOption {
	return func(r *RequestOptions) {
		r.HeaderAdd(key, value)
	}
}

func ReqOptOnResponse(value OnResponse) RequestOption {
	return func(r *RequestOptions) {
		r.OnResponse = value
	}
}
