package sts

func NewStatement() *Statement {
	return &Statement{
		Principal: &Principal{},
	}
}

type Statement struct {
	Principal *Principal                   `json:"principal"`
	Effect    string                       `json:"effect"`
	Action    []string                     `json:"action"`
	Resource  []string                     `json:"resource"`
	Condition map[string]map[string]string `json:"condition"`
}

func (s *Statement) AppendAction(actions ...string) {
	s.Action = append(s.Action, actions...)
}

func (s *Statement) AppendResource(resource ...string) {
	s.Resource = append(s.Resource, resource...)
}

type Principal struct {
	Qcs []string `json:"qcs"`
}

func (p *Principal) AppendQcs(qcs ...string) {
	p.Qcs = append(p.Qcs, qcs...)
}
