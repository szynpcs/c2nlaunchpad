package main

import (
	"c2n/tutorial/2_pointer/song"
	"fmt"
)

func main() {
	fmt.Println("等一轮风声")
	a := 10
	fmt.Println(a)
	p := &a
	*p = 213
	fmt.Println(*p)
	fmt.Println(a)
	changeIn(&a)
	fmt.Println(a)

	fmt.Println("====================================")

	s := song.Song{Name: "吹梦到西洲"}
	fmt.Println(s)

	author := song.Author{}
	author.SetAge(30).SetName("继国缘一")
	fmt.Println(author)
	s.Author = &author
	fmt.Println(s)
	fmt.Println(s.Author)
	fmt.Println(*s.Author)
	author.SetAge(99)
	fmt.Println(s)
	fmt.Println(s.Author)
	fmt.Println(*s.Author)

}

func changeIn(a *int) {
	*a = *a + 1
}
