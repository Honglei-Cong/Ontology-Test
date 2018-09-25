
package main
import "fmt"

func main() {
        a := make(map[int]string)
        a[10] = "10"

        b := a
        b[10] = "100"
        fmt.Println(a[10])
}

