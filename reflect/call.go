package reflect

import (
	"reflect"
)

func CallResultByIndex(a any, index int, args ...any) reflect.Value {
	return CallResult(ValueOf(a).Method(index), args...)
}

func CallByIndex(a any, index int, args ...any) []reflect.Value {
	return Call(ValueOf(a).Method(index), args...)
}

func CallResultByName(a any, name string, args ...any) reflect.Value {
	return CallResult(ValueOf(a).MethodByName(name), args...)
}

func CallByName(a any, name string, args ...any) []reflect.Value {
	return Call(ValueOf(a).MethodByName(name), args...)
}

func CallResult(m reflect.Value, args ...any) reflect.Value {
	result := Call(m, args...)
	if len(result) > 0 {
		return result[0]
	}
	return reflect.Value{}
}

func Call(m reflect.Value, args ...any) []reflect.Value {
	in := make([]reflect.Value, len(args))
	for i, _ := range args {
		in[i] = reflect.ValueOf(args[i])
	}
	return m.Call(in)
}
