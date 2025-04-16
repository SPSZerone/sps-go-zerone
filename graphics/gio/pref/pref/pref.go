package pref

func NewPref() Pref {
	return Pref{
		TableStyle: NewTableStyle(),
	}
}

type Pref struct {
	TableStyle TableStyle
}
