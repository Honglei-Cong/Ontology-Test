
package main

import (
        "math/big"
        "fmt"
)

func main() {
        e := big.NewFloat(100)
        r := big.NewFloat(10000.0)
        // s := big.NewFloat(500000.0)
        // f := big.NewFloat(0.2)

        t1 := new(big.Float).Add(new(big.Float).Quo(e, r), (big.NewFloat(1.0)))
        fmt.Println("t1", t1)
}

