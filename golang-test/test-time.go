
package main

import (
        "fmt"
        "time"
)

func main() {
        fmt.Println(time.Now().Format("20060102150405"))
        fmt.Println(time.Now().Unix())

        // timeStr := "20170112030405"
        timeStr := "20160501150405"
        tm, _ := time.Parse("20060102150405", timeStr)
        fmt.Println(tm.Unix())

        tm2 := time.Unix(10, 0)

        fmt.Println(tm.Format("2006-01-02 15:04:05"))
        fmt.Println(tm2.Format("2006-01-02 15:04:05"))


}
