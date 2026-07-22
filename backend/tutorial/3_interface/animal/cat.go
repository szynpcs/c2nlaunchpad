package animal

type Cat struct {
	Name string
}

func (cat Cat) Speak() string {
	return "喵咪咪"
}
