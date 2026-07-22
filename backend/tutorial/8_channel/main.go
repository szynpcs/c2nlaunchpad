package main

import "fmt"

func main() {
	fmt.Println("abcd")
	ch := make(chan int, 3)
	go func(a int) {
		ch <- a
	}(1)

	go func(a int) {
		ch <- a
	}(2)

	go func(a int) {
		ch <- a
	}(6)

	for i := 0; i < 3; i++ {
		result := <-ch
		fmt.Println(result)
	}
}
