package num

// #cgo LDFLAGS: -L. -lstdc++
// #cgo CXXFLAGS: -std=c++14 -I.
// #include "fib.h"
import "C"
import _ "unsafe"

func Fib(n int) int64 {
	return int64(C.Fib(C.int(n)))
}
