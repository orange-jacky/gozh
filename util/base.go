package util

import (
	"reflect"
)

////////////////////////////////////////////////////////////////////////////
func GetStructName(t interface{}) (name string) {
	ty := reflect.TypeOf(t)
	switch ty.Kind() {
	case reflect.Ptr:
		name = ty.Elem().Name()
	}
	return name
}
