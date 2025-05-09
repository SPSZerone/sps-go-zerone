package json

import (
	"net/http"

	"google.golang.org/protobuf/proto"

	spshttp "github.com/SPSZerone/sps-go-zerone/net/http"
	spsheader "github.com/SPSZerone/sps-go-zerone/net/http/header"
)

func NewHttpWriter(msg proto.Message, opts ...MarshalOption) *HttpWriter {
	return &HttpWriter{
		Msg:  msg,
		Opts: opts,
	}
}

type HttpWriter struct {
	Msg  proto.Message
	Opts []MarshalOption
}

func (w *HttpWriter) Render(writer http.ResponseWriter) error {
	_, err := WriteJsonUtf8(writer, w.Msg, w.Opts...)
	return err
}

func (w *HttpWriter) WriteContentType(writer http.ResponseWriter) {
	spsheader.SetContentTypeJsonUtf8(writer)
}

func WriteJsonUtf8(w http.ResponseWriter, msg proto.Message, opts ...MarshalOption) (int, error) {
	jsonBytes, err := Marshal(msg, opts...)
	if err != nil {
		return 0, err
	}
	return spshttp.WriteJsonUtf8(w, jsonBytes)
}
