// Package helpers contains helper or generic functions
package helpers

func IsUint(value interface{}) bool {
	_, ok := value.(uint)
	return ok
}
