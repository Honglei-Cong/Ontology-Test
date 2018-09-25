
package main

import (
        "fmt"
)

func main() {
        i := 0 
        defer fmt.Printf("i4 = %d \n", i)
        func (){
                defer fmt.Printf("i2 = %d \n", i)
                i++
                fmt.Printf("i1 = %d \n", i)
                i++
        }()
        i++
        fmt.Printf("i3 = %d \n", i)
        i++
}
