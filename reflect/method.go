package reflect

import (
	"reflect"
)

func Methods(a any) (result []reflect.Method) {
	aType := TypeOf(a)
	for i := 0; i < aType.NumMethod(); i++ {
		method := aType.Method(i)
		result = append(result, method)
	}
	return
}
