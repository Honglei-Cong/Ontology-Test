
package main

// int add(int a, int b) {
//      return a+b;
// }

/*
#include "foo.h"
*/
import "C"
import "fmt"

func main() {
        a := C.int(1)
        b := C.int(2)
        value := C.add(a, b)
        fmt.Printf("%v \n", value)
        fmt.Printf("%v \n", int(value))
}
