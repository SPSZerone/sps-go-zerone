package reflect

import (
	"reflect"
)

func TypeOf(a any) reflect.Type {
	value := reflect.ValueOf(a)
	return value.Type()
}
