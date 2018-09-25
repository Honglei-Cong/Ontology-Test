
package main

import (
        "fmt"
        "time"
)

func main() {

        duration, err := time.ParseDuration("5s")
        if err != nil {
                fmt.Println("init dur failed", err)
                return
        }

        fmt.Println("start")
        timechan := time.After(duration)
        time.Sleep(4 * time.Second)
        fmt.Println("after 3s")
        select {
        case <- timechan:
                fmt.Println("5s timeout")
        }

        t2 := time.NewTimer(1 * time.Second)
        cnt := 5
        for cnt > 0 {
                select {
                case <-t2.C:
                        t2.Reset(1 * time.Second)
                }
                fmt.Printf("cnt: %d \n", cnt)
                cnt--
        }

        n := 10
        C := make(chan int)
        var ticker *time.Timer
        ticker = time.AfterFunc(time.Second, func() {
                C <- n
                fmt.Printf("n: %d\n", n)
                ticker.Reset(time.Second)
        })

        n += 1
        done := make(chan struct{})
        go func() {
                defer close(done)
                for {
                        select {
                        case i := <-C:
                                if i == n {
                                        return
                                }
                                n += 1
                        }
                }
        }()

        <-done
        fmt.Println("done")
}

