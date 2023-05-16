package main

import (
	"fmt"
	"time"
)

func main() {

	// channel
	ch := make(chan int)
	// 每个管道都有一个类型，并指定一个buffer size，如果不指定，默认为1

	// write into channel
	// will block if channel is full
	ch <- 1

	// read from channel
	n := <-ch
	fmt.Println(n)

	go fmt.Println("hello")
	go fmt.Println("world")
	time.Sleep(100 * time.Millisecond)
}
