
package main

import "fmt"

func printArray(strs []string) {
    if strs == nil {
        fmt.Println("nil string array")
    }
	for i := 0; i < len(strs); i++ {
		fmt.Println(strs[i])
	}
    fmt.Println("")
}

func printMap(m map[string]string) {
    for k, v := range m {
        fmt.Println(k, v)
    }
    fmt.Println("")
}

func main() {
	a := []string{"x", "y"}
	b := []string{"a", "b"}

	printArray(append(append(a, b...), "c", "d"))

    idx := 0
    var s string
    for idx, s = range a {
        fmt.Println(s)
    }
    fmt.Printf("idx = %d\n\n", idx)

    x1 := "abc"
    printArray(append([]string{x1}, a...))

    m := make(map[string]string)
    m["test"] = "test"
    m["test2"] = "test2"
    defer printMap(m)
    delete(m, "test")

    m1 := byte(0x10)
    fmt.Printf("--%s--%d\n", string(m1), len(string(m1)))

    xx := make(map[string]interface{})
    xx["0"] = 2.0
    d, ok := xx["0"].(int64)
    fmt.Printf("%d, %v\n", d, ok)

    var x [10]byte
    fmt.Printf("x len: %d \n", len(x))
}
