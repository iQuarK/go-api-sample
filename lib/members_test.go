package lib_test

import (
	"github.com/iquark/go-api-sample/lib"
	"reflect"
	"testing"
)

func TestGetMembers(t *testing.T) {
	alice := lib.NewPerson(0, "Alice", []lib.Person{})
	bob := lib.NewPerson(1, "Bob", []lib.Person{})
	carlos := lib.NewPerson(2, "Carlos", []lib.Person{})
	charlie := lib.NewPerson(4, "Charlie", []lib.Person{})
	eve := lib.NewPerson(7, "Eve", []lib.Person{})

	aTeam := lib.NewPerson(90, "The A-Team", []lib.Person{alice, bob, carlos})
	cTeam := lib.NewPerson(92, "The C-Team", []lib.Person{charlie, eve, aTeam})

	t.Log("It should return the list of members for a given team")

	list := lib.GetMembers(cTeam)

	expected := []lib.Person{charlie, eve, alice, bob, carlos}

	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected %v, but it was %v instead.", expected, list)
	}
}
