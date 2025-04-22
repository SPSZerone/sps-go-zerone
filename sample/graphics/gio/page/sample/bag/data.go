package bag

func NewData() Data {
	return Data{}
}

type Data struct {
	Properties []Property
}

func (d *Data) AddProperties(properties ...Property) {
	d.Properties = append(d.Properties, properties...)
}

func (d *Data) GetProperty(key PropertyKey) *Property {
	if key < 0 || int(key) >= len(d.Properties) {
		return nil
	}
	return &d.Properties[key]
}
