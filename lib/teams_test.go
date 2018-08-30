package lib_test

import (
	"github.com/iquark/go-api-sample/lib"
	"reflect"
	"testing"
)

func TestGetTeams(t *testing.T) {
	alice := lib.NewPerson(0, "Alice", []lib.Person{})
	bob := lib.NewPerson(1, "Bob", []lib.Person{})
	carlos := lib.NewPerson(2, "Carlos", []lib.Person{})
	carol := lib.NewPerson(3, "Carol", []lib.Person{})
	charlie := lib.NewPerson(4, "Charlie", []lib.Person{})
	chuck := lib.NewPerson(5, "Chuck", []lib.Person{})
	dave := lib.NewPerson(6, "Dave", []lib.Person{})
	eve := lib.NewPerson(7, "Eve", []lib.Person{})
	mallory := lib.NewPerson(8, "Mallory", []lib.Person{})
	peggy := lib.NewPerson(9, "Peggy", []lib.Person{})
	trent := lib.NewPerson(10, "Trent", []lib.Person{})
	victor := lib.NewPerson(11, "Victor", []lib.Person{})
	walter := lib.NewPerson(12, "Walter", []lib.Person{})

	aTeam := lib.NewPerson(90, "The A-Team", []lib.Person{alice, bob, carlos})
	bTeam := lib.NewPerson(91, "The B-Team", []lib.Person{peggy, trent, victor, bob})
	cTeam := lib.NewPerson(92, "The C-Team", []lib.Person{charlie, eve, aTeam})

	people := []lib.Person{alice, bob, carlos, carol, charlie, chuck, dave, eve, mallory, peggy, trent, victor, walter, aTeam, bTeam, cTeam}

	t.Log("It should return the list of teams for a given user")

	list, error := lib.GetTeams(alice, people)

	expected := []lib.Person{aTeam, cTeam}

	switch {
	case error != nil:
		t.Error("Expected [alice], but it was error instead.")
	case !reflect.DeepEqual(list, expected):
		t.Errorf("Expected /n%v,/n but it was /n%v/n instead.", expected, list)
	}
}
