package main

import (
	"c2n/tutorial/3_interface/animal"
	"fmt"
)

func main() {
	fmt.Println("INTERFACE的学习")

	var a animal.Animal = &animal.Dog{"旺财"}
	fmt.Println(a.Speak())

	var b animal.Animal = &animal.Bird{Name: "蓝孔雀"}
	fmt.Println(b.Speak())

	var b2 animal.Animal = &animal.Bird{Name: "绿宝石"}
	fmt.Println(b2.Speak())

	var c animal.Animal = animal.Cat{Name: "雪球馒头子"}
	fmt.Println(c.Speak())
}
