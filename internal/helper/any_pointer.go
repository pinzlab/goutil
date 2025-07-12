package helper

func Pointer[V any](value V) *V {
	return &value
}
