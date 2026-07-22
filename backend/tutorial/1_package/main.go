package main

import (
	"fmt"
	"packagetutorial/out"
)

func main() {
	fmt.Print("你是谁")
	fmt.Println(out.EditorNameInfo)

	fmt.Println(out.EditorName())
	out.SetEditor("继国缘一")
	fmt.Println(out.EditorName())

	song := out.NewSong("花开不败", "尚志远", "继国缘一")
	song.Print()

	song.SetAuthor("继国衍生")
	song.Print()

}
