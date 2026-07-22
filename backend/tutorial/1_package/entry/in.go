package main

import (
	"fmt"
	"packageStudy/out"
)

func main() {
	fmt.Println("等一轮风声")
	fmt.Println("熟透了夜深")
	song := out.Song{
		Name: "怜城辞",
	}
	song.Print()
	song.Singer = "鲁哈"
	song.Print()

	song.SetAuthor("谢怜")
	fmt.Println(song.Author())
	song.Print()

	newSong := out.NewSong("无别", "张信哲", "天官赐福")
	newSong.Print()
	newSong.Name = "不再别"
	newSong.Name = "情歌王子张信哲"
	newSong.SetAuthor("天官赐福，百无禁忌")
	newSong.Print()

	fmt.Println(out.EditorNameInfo)

	fmt.Println(out.EditorName())
	out.SetEditor("继国缘一")
	fmt.Println(out.EditorName())

}
