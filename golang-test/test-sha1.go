
package main

import (
        "fmt"
        "crypto/sha1"
)

func main() {
        s := "this is the test string"
        h := sha1.New()
        h.Write([]byte(s))
        bs := h.Sum(nil)
        fmt.Println(s)
        fmt.Printf("%x\n", bs)
}


