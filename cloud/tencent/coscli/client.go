package coscli

import (
	"net/http"
	"net/url"

	"github.com/tencentyun/cos-go-sdk-v5"
)

func NewClient() (*cos.Client, error) {
	cfg, err := GetConfig()
	if err != nil {
		return nil, err
	}

	bucketUrl, err := url.Parse(cfg.BucketUrl)
	if err != nil {
		return nil, err
	}

	cosBucketUrl := &cos.BaseURL{BucketURL: bucketUrl}
	httpClient := &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  cfg.SecretId,
			SecretKey: cfg.SecretKey,
		},
	}

	client := cos.NewClient(cosBucketUrl, httpClient)
	return client, nil
}
