package utils

import (
	"reflect"
)

func GetType(data any) string {
	switch data.(type) {
	case int, float32, float64, uint, uint8, uint16, uint32, uint64:
		return "int"
	default:
		return reflect.TypeOf(data).String()
	}
	return ""
}
