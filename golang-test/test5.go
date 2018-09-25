
package main

import (
        "fmt"
        "encoding/json"
)

type X struct {
        U  string `json:"u"`
        S  string `json:"s"`
}

func setX(x X) {
        x.U = "test showX"
}

func setMap(m map[int]int) {
        m[1] = 1
}

func setMap2(m map[int]int) {
        // update m
        m = make(map[int]int)
        m[1] = 2
}

func main() {
        s := "{\"u\":\"test\"}"
        x := &X{}
        if err := json.Unmarshal([]byte(s), x); err != nil {
                fmt.Println(err.Error())
        }
        fmt.Printf("%s | %s\n", x.U, x.S)
        b, _ := json.Marshal(x)
        fmt.Println(string(b))

        err := fmt.Errorf("test")
        var err2 error
        fmt.Printf("print nil err: %s, %s \n", err, err2)

        x2 := X{"a", "b"}
        setX(x2)
        fmt.Printf("x: %s \n", x2.U)

        m := make(map[int]int)
        setMap(m)
        fmt.Printf("m: %d \n", m[1])
        setMap2(m)
        fmt.Printf("m: %d \n", m[1])
}

