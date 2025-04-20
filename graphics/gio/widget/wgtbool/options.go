package wgtbool

type Option func(*Bool)

func OptName(value string) Option {
	return func(o *Bool) {
		o.Name = value
	}
}

func OptDesc(value string) Option {
	return func(o *Bool) {
		o.Desc = value
	}
}

func OptValue(value bool) Option {
	return func(o *Bool) {
		o.Value = value
	}
}
