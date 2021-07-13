package domain

type Model struct {
	Name string
}

func (s *Model) SayHi() string {
	return "Hi, " + s.Name
}
