package animal

type Bird struct {
	Name string
}

func (b *Bird) Speak() string {
	return "咕咕级"
}
