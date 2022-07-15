package panics

import (
	"fmt"
	"strings"
)

func If0(err error) {
	if err != nil {
		panic(err)
	}
}
func If1[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}
func If2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return v1, v2
}
func If3[T1, T2, T3 any](v1 T1, v2 T2, v3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3
}
func If4[T1, T2, T3, T4 any](v1 T1, v2 T2, v3 T3, v4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return v1, v2, v3, v4
}

func buildError(err error, args ...interface{}) error {
	if err == nil {
		return nil
	}
	if len(args) == 0 {
		return err
	}
	if v, ok := args[0].(string); ok && strings.Contains(v, "%") {
		args = append(args, err)
		return fmt.Errorf(v+", error: %w", args[1:]...)
	}
	return fmt.Errorf(fmt.Sprint(args...)+", error: %w", err)
}
