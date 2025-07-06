package v3

import (
	"fmt"

	spssha256 "github.com/SPSZerone/sps-go-zerone/algorithm/crypto/sha256"
	spshttp "github.com/SPSZerone/sps-go-zerone/net/http"
	spshttpheader "github.com/SPSZerone/sps-go-zerone/net/http/header"
)

const (
	SignedHeadersRequire = "content-type;host"
)

func NewCanonicalRequest(host string) *CanonicalRequest {
	return &CanonicalRequest{
		HttpRequestMethod:    spshttp.MethodPost,
		CanonicalURI:         "/",
		CanonicalQueryString: "",
		SignedHeaders:        SignedHeadersRequire,
		CanonicalHeaders:     CanonicalHeadersRequire(host),
	}
}

func CanonicalHeadersRequire(host string) string {
	return fmt.Sprintf("content-type:%s\nhost:%s\n", spshttpheader.ContentTypeJsonUtf8, host)
}

type CanonicalRequest struct {
	HttpRequestMethod    string
	CanonicalURI         string
	CanonicalQueryString string
	CanonicalHeaders     string
	SignedHeaders        string
	HashedRequestPayload string
}

func (r *CanonicalRequest) Build() string {
	return spssha256.Hex(r.BuildCanonicalRequest())
}

func (r *CanonicalRequest) BuildCanonicalRequest() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n%s",
		r.HttpRequestMethod,
		r.CanonicalURI,
		r.CanonicalQueryString,
		r.CanonicalHeaders,
		r.SignedHeaders,
		r.HashedRequestPayload,
	)
}
