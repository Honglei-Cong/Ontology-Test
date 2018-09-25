
package main

/*
#cgo CFLAGS: -I./foo
#cgo LDFLAGS: -L./foo -lfoo
#include "foo.h"
#include <stdlib.h>
*/
import "C"

import (
        "fmt"
        "reflect"
        "unsafe"
)

func pass() {
        in := []C.int{1, 2, 3, 4}
        inPointer := unsafe.Pointer(&in[0])
        inC := (*C.int)(inPointer)

        value := C.pass_array_pointer(inC, 4)
        fmt.Println(value)
}

func get() {
        n := 10
        outC := C.return_array_pointer(C.int(n))
        defer C.free(unsafe.Pointer(outC))

        sh := reflect.SliceHeader{uintptr(unsafe.Pointer(outC)), n, n}
        out := *(*[]C.int)(unsafe.Pointer(&sh))
        for _, v := range out {
                fmt.Println(v)
        }
}

func main() {
        pass()
        get()
}

