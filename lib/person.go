package lib

// Person definition, it can be a team or no, depending if it has members or no
type Person struct {
	ID          int32
	DisplayName string
	Members     []Person
	IsTeam      bool
}

// People interface definition
type People interface {
	NewPerson(id int32, displayName string, members []Person) Person
}

// NewPerson instantiates a new Person
func NewPerson(id int32, displayName string, members []Person) Person {
	return Person{ID: id, DisplayName: displayName, Members: members, IsTeam: len(members) > 0}
}
