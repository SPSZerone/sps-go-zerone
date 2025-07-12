package json

import (
	"google.golang.org/protobuf/proto"
)

func Format(msg proto.Message, opts ...MarshalOption) string {
	return NewMarshalOptions(opts...).Format(msg)
}

func Marshal(msg proto.Message, opts ...MarshalOption) ([]byte, error) {
	return NewMarshalOptions(opts...).Marshal(msg)
}
