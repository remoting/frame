package util

import "github.com/remoting/frame/util/conv"

func Int(value interface{}) int {
	return conv.Int(value)
}
func String(value interface{}) string {
	return conv.String(value)
}
func Int64(value interface{}) int64 {
	return conv.Int64(value)
}
func Float32(value interface{}) float32 {
	return conv.Float32(value)
}
func Float64(value interface{}) float64 {
	return conv.Float64(value)
}
