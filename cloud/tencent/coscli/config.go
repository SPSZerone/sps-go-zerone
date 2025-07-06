package coscli

import (
	"fmt"
	"os"
)

const (
	EnvTencentCosBucketUrl = "TENCENT_COS_BUCKET_URL"
	EnvTencentCosSecretId  = "TENCENT_COS_SECRET_ID"
	EnvTencentCosSecretKey = "TENCENT_COS_SECRET_KEY"
)

func GetConfig() (*Config, error) {
	bucketUrl := GetBucketUrl()
	secretId := GetSecretId()
	secretKey := GetSecretKey()
	if bucketUrl == "" ||
		secretId == "" ||
		secretKey == "" {
		return nil, fmt.Errorf("tencent cos config invalid")
	}
	cfg := &Config{
		BucketUrl: bucketUrl,
		SecretId:  secretId,
		SecretKey: secretKey,
	}
	return cfg, nil
}

type Config struct {
	BucketUrl string
	SecretId  string
	SecretKey string
}

func GetBucketUrl() string {
	return os.Getenv(EnvTencentCosBucketUrl)
}

func GetSecretId() string {
	return os.Getenv(EnvTencentCosSecretId)
}

func GetSecretKey() string {
	return os.Getenv(EnvTencentCosSecretKey)
}
