
package main

import ("fmt")

func main() {
        s := []int{1, 2, 3}
        x := make([]int, 0, len(s))
        x = append(x, s...)
        s = []int{}
        fmt.Printf("%v, %v\n", s, x)
}
