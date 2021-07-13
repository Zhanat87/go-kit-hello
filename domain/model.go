package domain

const Greeting = "Hi, "

type Model struct {
	Name string
}

func (s *Model) SayHi() string {
	return Greeting + s.Name
}
