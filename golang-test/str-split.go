
package main

import (
        "fmt"
        "strings"
)

func main() {
        name := "user0"
        fmt.Println(strings.Split(name, "user")[1])

        sql := " select   * from xxx  ; "
        for _, s := range strings.SplitN(sql, ";", 2) {
                fmt.Printf("'%s'\n", s)
        }
        fmt.Println("===========")
        for _, s := range strings.Split(sql, " ") {
                fmt.Printf("'%s'\n", s)
        }
        fmt.Println("===========")
        for _, s := range strings.SplitN(sql, ";", 100) {
                fmt.Printf("'%s'\n", s)
        }

        fmt.Printf("split test: %v \n", len(strings.SplitN("test", ".", 2)))


        x := 200
        fmt.Printf("%016x \n", x)
}

