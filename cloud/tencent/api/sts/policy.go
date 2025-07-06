package sts

import (
	"encoding/json"
)

func NewPolicy(stmts ...*Statement) *Policy {
	policy := &Policy{
		Version: "2.0",
	}
	policy.AppendStatement(stmts...)
	return policy
}

type Policy struct {
	Version   string       `json:"version"`
	Statement []*Statement `json:"statement"`
}

func (p *Policy) AppendStatement(stmts ...*Statement) {
	p.Statement = append(p.Statement, stmts...)
}

func (p *Policy) ToJson() string {
	policyJsonBytes, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(policyJsonBytes)
}

func (p *Policy) MarshalJSON() ([]byte, error) {
	return json.Marshal(p)
}
