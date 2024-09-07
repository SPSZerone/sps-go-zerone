package error

import (
	"fmt"
)

type Option func(e *Error)

type Error struct {
	FormatPrefix string
	ArgsPrefix   []any

	FormatSuffix string
	ArgsSuffix   []any
}

func (e *Error) Reset(opts ...Option) {
	e.FormatPrefix = ""
	e.ArgsPrefix = nil

	e.FormatSuffix = ""
	e.ArgsSuffix = nil

	e.Update(opts...)
}

func (e *Error) Update(opts ...Option) {
	for _, opt := range opts {
		opt(e)
	}
}

func (e *Error) Errorf(format string, args ...any) error {
	argsFinal := make([]any, 0, len(e.ArgsPrefix)+len(args)+len(e.ArgsSuffix))
	formatFinal := fmt.Sprintf("%s%s%s", e.FormatPrefix, format, e.FormatSuffix)
	argsFinal = append(argsFinal, e.ArgsPrefix...)
	argsFinal = append(argsFinal, args...)
	argsFinal = append(argsFinal, e.ArgsSuffix...)
	return fmt.Errorf(formatFinal, argsFinal...)
}

func OptFormatPrefix(value string) Option {
	return func(e *Error) {
		e.FormatPrefix = value
	}
}

func OptArgsPrefix(value ...any) Option {
	return func(e *Error) {
		e.ArgsPrefix = value
	}
}

func OptFormatSuffix(value string) Option {
	return func(e *Error) {
		e.FormatSuffix = value
	}
}

func OptArgsSuffix(value ...any) Option {
	return func(e *Error) {
		e.ArgsSuffix = value
	}
}
