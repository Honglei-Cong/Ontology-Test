
package main

import (
        "fmt"
)

func shiftSlice(x []int, n int) []int {
        n = n % len(x)
        return append(x[n:], x[:n]...)
}

func main() {
        x := []int{1, 2, 3, 4, 5, 6, 7}
        fmt.Printf("%v \n", shiftSlice(x, 0))
        fmt.Printf("%v \n", shiftSlice(x, 1))
        fmt.Printf("%v \n", shiftSlice(x, 2))
        fmt.Printf("%v \n", shiftSlice(x, 3))
        fmt.Printf("%v \n", shiftSlice(x, 3))
        fmt.Printf("%v \n", shiftSlice(x, 4))
        fmt.Printf("%v \n", shiftSlice(x, 5))
        fmt.Printf("%v \n", shiftSlice(x, 6))
        fmt.Printf("%v \n", shiftSlice(x, 7))
        fmt.Printf("%v \n", shiftSlice(x, 8))
}


