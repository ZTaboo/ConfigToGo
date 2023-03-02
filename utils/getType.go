package utils

import (
	"math"
	"reflect"
)

func GetType(data any) string {
	switch data.(type) {
	case int, uint, uint8, uint16, uint32, uint64:
		return "int"
	case float64, float32:
		// 判断是否真的为float
		if data.(float64) == math.Trunc(data.(float64)) {
			return "int"
		} else {
			return "float64"
		}
	default:
		return reflect.TypeOf(data).String()
	}
}
