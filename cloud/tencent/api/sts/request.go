package sts

import (
	"fmt"
	"io"

	spstencentsignv3 "github.com/SPSZerone/sps-go-zerone/cloud/tencent/api/public/signature/v3"
	spshttp "github.com/SPSZerone/sps-go-zerone/net/http"
	ssphttpheader "github.com/SPSZerone/sps-go-zerone/net/http/header"
)

var (
	Domain  = "sts.tencentcloudapi.com"
	Service = "sts"
	Version = "2018-08-13"

	HttpsUrl = fmt.Sprintf("https://%s/", Domain)
)

// Request sts default request
func Request(param *spstencentsignv3.ParamPublic, body io.Reader, onResponse spshttp.OnResponse) ([]byte, error) {
	opts := NewRequestOptions(param, body, onResponse)
	return spshttp.DoRequest(opts)
}

func NewRequestOptions(param *spstencentsignv3.ParamPublic, body io.Reader, onResponse spshttp.OnResponse) *spshttp.RequestOptions {
	headerOpts := map[string]string{
		ssphttpheader.Authorization: param.Authorization,
		ssphttpheader.ContentType:   ssphttpheader.ContentTypeJsonUtf8,
		ssphttpheader.Host:          param.Host,
		"X-TC-Action":               param.Action,
		"X-TC-Timestamp":            param.TimestampStr(),
		"X-TC-Version":              param.Version,
		"X-TC-Region":               param.Region,
	}
	return spshttp.NewRequestOptions(spshttp.MethodPost, HttpsUrl, body, onResponse, spshttp.ReqOptHeader(headerOpts))
}
