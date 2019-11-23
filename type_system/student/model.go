package student

type Student struct {
	firstName string `json:"first_name" bson: "full_name" validate: "required"`
	lastName  string
	Age       int
	Email     string
}

func (s Student) GetFullName() string {
	return s.firstName + " " + s.lastName
}

func (s *Student) SetEmail(e string) {
	s.Email = e
}
