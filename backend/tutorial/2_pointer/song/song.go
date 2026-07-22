package song

type Song struct {
	Name   string
	Author *Author
}

type Author struct {
	Name string
	Age  int
}

func (a *Author) SetName(name string) *Author {
	a.Name = name
	return a
}

func (a *Author) SetAge(age int) *Author {
	a.Age = age
	return a
}
