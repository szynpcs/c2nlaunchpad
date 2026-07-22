package animal

type Dog struct {
	Name string
}

func (dog Dog) Speak() string {
	return "旺旺"
}
