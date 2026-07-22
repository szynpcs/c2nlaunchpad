package main

import (
	"fmt"
	"sync"
)

func main() {
	//fmt.Println("abc")
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()

		fmt.Println("wg1")
	}()

	go func() {
		defer wg.Done()

		//defer wg.Done()
		fmt.Println("wg2")
		//wg.Done()
	}()

	go func() {
		defer wg.Done()

		//defer wg.Done()
		fmt.Println("wg3")
		//wg.Done()
	}()

	wg.Wait()
	fmt.Println("over")
}
