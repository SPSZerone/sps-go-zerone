package v3

import (
	"fmt"
	"time"
)

const (
	AlgorithmTC3HmacSha256 = "TC3-HMAC-SHA256"
)

func NewConfig(opts ...ConfigOption) *Config {
	c := &Config{
		Algorithm: AlgorithmTC3HmacSha256,
		Time:      time.Now(),
	}
	c.Update(opts...)
	return c
}

type Config struct {
	SecretId  string
	SecretKey string
	Host      string
	Algorithm string
	Service   string
	Version   string
	Action    string
	Region    string
	Time      time.Time
}

func (c *Config) Update(opts ...ConfigOption) {
	for _, opt := range opts {
		opt(c)
	}
}

func (c *Config) Build() (timestamp int64, date, credentialScope string) {
	timestamp, date = c.BuildTime()
	credentialScope = fmt.Sprintf("%s/%s/tc3_request", date, c.Service)
	return
}

func (c *Config) BuildTime() (timestamp int64, date string) {
	timestamp = c.Time.Unix()
	date = time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	return
}

type ConfigOption func(*Config)

func CfgOptSecretId(value string) ConfigOption {
	return func(c *Config) {
		c.SecretId = value
	}
}

func CfgOptSecretKey(value string) ConfigOption {
	return func(c *Config) {
		c.SecretKey = value
	}
}

func CfgOptHost(value string) ConfigOption {
	return func(c *Config) {
		c.Host = value
	}
}

func CfgOptAlgorithm(value string) ConfigOption {
	return func(c *Config) {
		c.Algorithm = value
	}
}

func CfgOptService(value string) ConfigOption {
	return func(c *Config) {
		c.Service = value
	}
}

func CfgOptVersion(value string) ConfigOption {
	return func(c *Config) {
		c.Version = value
	}
}

func CfgOptAction(value string) ConfigOption {
	return func(c *Config) {
		c.Action = value
	}
}

func CfgOptRegion(value string) ConfigOption {
	return func(c *Config) {
		c.Region = value
	}
}

func CfgOptTime(value time.Time) ConfigOption {
	return func(c *Config) {
		c.Time = value
	}
}
