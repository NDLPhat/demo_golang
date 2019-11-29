package student

type Student struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
}

func (s Student) GetFullName() string {
	return s.FirstName + " " + s.LastName
}

func (s *Student) SetEmail(e string) {
	s.Email = e
}
