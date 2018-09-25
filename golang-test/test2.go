
package main

import (
	"fmt"
	"time"
)

func prod(c chan int) {
	defer close(c)
	defer func() { fmt.Println("prod done") }()

	for i := 0; i < 5; i++ {
		c <- i
	}

}

func consm(c chan int) {
	time.Sleep(1 * time.Second)
	for x := range(c) {
		fmt.Println(x)
	}
}

func main() {
	c := make(chan int, 10)
	go prod(c);
	go consm(c);

	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
