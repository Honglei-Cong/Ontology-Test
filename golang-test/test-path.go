
package main

import (
        "fmt"
        "path/filepath"
)

func main() {
        path := "/Users/conghonglei/Workspace/golang/test"
        fmt.Println(filepath.Base(path))

        path = "/Users/conghonglei/Workspace/golang/"
        fmt.Println(filepath.Base(path))
}
