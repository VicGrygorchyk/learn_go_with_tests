package reflect

import (
	"reflect"
)

func Walk(x interface{}, fn func(inp string)) {
	val := reflect.ValueOf(x)

	fieldsNum := val.NumField()

	for i :=0; i < fieldsNum; i++ {
		field := val.Field(i)

		if field.Kind() == reflect.String {
            fn(field.String())
        }
		if field.Kind() == reflect.Struct {
            Walk(field.Interface(), fn)
        }
	}
}