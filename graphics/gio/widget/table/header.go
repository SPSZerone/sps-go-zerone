package table

func NewHeader(text string) Header {
	return Header{
		Text: text,
	}
}

type Header struct {
	Text string
}
