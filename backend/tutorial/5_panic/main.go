package main

import "fmt"

func main() {
	defer fmt.Println("justlastdance")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("哦吼吼吼吼哦豁")
	panic("多看一眼就会爆炸")
	fmt.Println("呀哈")
}
