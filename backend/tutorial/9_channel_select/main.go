package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("hello world")

	ch := make(chan string)

	go func() {
		time.Sleep(10 * time.Second)
		ch <- "希望我们有再见得那一天把"
	}()

	select {
	case msg := <-ch:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("你们不会有再见的那一天了")
	}
}
