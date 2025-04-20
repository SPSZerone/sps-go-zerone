package reflect

import (
	"reflect"

	"github.com/SPSZerone/sps-go-zerone/generic"
)

func NumericByName[T generic.Numeric](a any, name string, args ...any) (t T) {
	result := CallResultByName(a, name, args...)
	return Numeric[T](result)
}

func NumericByIndex[T generic.Numeric](a any, index int, args ...any) (t T) {
	result := CallResultByIndex(a, index, args...)
	return Numeric[T](result)
}

func Numeric[T generic.Numeric](v reflect.Value) (t T) {
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return T(v.Float())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return T(v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return T(v.Uint())
	default:
	}
	return
}
