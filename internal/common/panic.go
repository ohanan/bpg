package common

func PanicIf0(err error) {
	if err != nil {
		panic(err)
	}
}
func PanicIf1[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
