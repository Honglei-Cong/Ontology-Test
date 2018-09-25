
package main

import(
        "fmt"
        "time"
        "encoding/json"
)

type User struct {
        ID      int
        Name    string
}

func main() {
        user1 := &User{ID: 123, Name: "user1"}
        jsonstr, _ := json.Marshal(user1)
        fmt.Println(string(jsonstr))

        fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

        layout := "2006-01-02 15:04:05.000"
        str := "2014-11-12 11:45:26.123"
        t1, err := time.Parse(layout, str)
        if err != nil {
                fmt.Println(err)
        }

        str2 := "2015-11-12 11:45:26.456"
        t2, err := time.Parse(layout, str2)
        if err != nil {
                fmt.Println(err)
        }
        
        if t1.Before(t2) {
                fmt.Println(t1)
        }
}

