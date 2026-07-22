package out

import "fmt"

// public变量
var EditorNameInfo string = "尚志远"

// private 变量 + 函数
var editorName string = "无名"

func SetEditor(editorNameNew string) {
	editorName = editorNameNew
}

func EditorName() string {
	return editorName
}

// struct 有public + private变量
type Song struct {
	Name   string
	Singer string
	author string
}

func (s *Song) Print() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Singer:", s.Singer)
	fmt.Println("author:", s.author)
}

func NewSong(name string, singer string, author string) *Song {
	return &Song{name, singer, author}
}

func (s *Song) Author() string {
	return s.author
}

func (s *Song) SetAuthor(author string) {
	s.author = author
}
