package json

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoregistry"
)

func NewMarshalOptions(opts ...MarshalOption) protojson.MarshalOptions {
	o := protojson.MarshalOptions{
		Multiline:         true,
		EmitDefaultValues: true,
	}
	UpdateMarshalOptions(&o, opts...)
	return o
}

func UpdateMarshalOptions(o *protojson.MarshalOptions, opts ...MarshalOption) {
	for _, opt := range opts {
		opt(o)
	}
}

type MarshalOption func(*protojson.MarshalOptions)

func MOptMultiline(value bool) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.Multiline = value
	}
}

func MOptIndent(value string) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.Indent = value
	}
}

func MOptAllowPartial(value bool) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.AllowPartial = value
	}
}

func MOptUseProtoNames(value bool) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.UseProtoNames = value
	}
}

func MOptUseEnumNumbers(value bool) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.UseEnumNumbers = value
	}
}

func MOptEmitUnpopulated(value bool) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.EmitUnpopulated = value
	}
}

func MOptEmitDefaultValues(value bool) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.EmitDefaultValues = value
	}
}

func MOptResolver(value Resolver) MarshalOption {
	return func(o *protojson.MarshalOptions) {
		o.Resolver = value
	}
}

type Resolver interface {
	protoregistry.ExtensionTypeResolver
	protoregistry.MessageTypeResolver
}
