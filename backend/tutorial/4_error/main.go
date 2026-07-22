package main

import (
	"c2n/tutorial/4_error/customerror"
	"errors"
	"fmt"
)

func main() {
	fmt.Println("错误定义")
	err := fmt.Errorf("大脑宕机了")
	fmt.Println(err)
	err1 := errors.New("大脑宕机了牙牙乐")
	fmt.Println(err1)
	brainError := customerror.BrainError{"大脑极度充沛了"}
	fmt.Println(brainError)
}
