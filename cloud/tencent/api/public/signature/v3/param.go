package v3

import (
	"fmt"
	"time"
)

func NewParamPublic(action string, region string, authorization string, opts ...ParamPublicOption) *ParamPublic {
	p := &ParamPublic{
		Action:        action,
		Region:        region,
		Authorization: authorization,
		Version:       "2018-08-13",
		Timestamp:     time.Now(),
	}
	p.Update(opts...)
	return p
}

type ParamPublicOption func(p *ParamPublic)

// ParamPublic
//
// see [DOC]
//
// [DOC]: https://cloud.tencent.com/document/api/1312/48201
type ParamPublic struct {
	Action        string // e.g. GetFederationToken
	Region        string // e.g. ap-beijing https://cloud.tencent.com/document/api/1312/48201#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8
	Authorization string
	Version       string
	Timestamp     time.Time
	Token         string
	Language      string // zh-CN en-US
	Host          string
}

func (p *ParamPublic) Update(opts ...ParamPublicOption) {
	for _, opt := range opts {
		opt(p)
	}
}

func (p *ParamPublic) TimestampStr() string {
	return fmt.Sprintf("%d", p.Timestamp.Unix())
}

func PPOptAction(value string) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Action = value
	}
}

func PPOptRegion(value string) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Region = value
	}
}

func PPOptTimestamp(value time.Time) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Timestamp = value
	}
}

func PPOptVersion(value string) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Version = value
	}
}

func PPOptAuthorization(value string) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Authorization = value
	}
}

func PPOptToken(value string) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Token = value
	}
}

func PPOptLanguage(value string) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Language = value
	}
}

func PPOptHost(value string) ParamPublicOption {
	return func(p *ParamPublic) {
		p.Host = value
	}
}
