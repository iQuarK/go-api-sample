package lib_test

import (
	"github.com/iquark/go-api-sample/lib"
	"testing"
)

func TestContainsPerson(t *testing.T) {
	alice := lib.NewPerson(0, "Alice", []lib.Person{})
	bob := lib.NewPerson(1, "Bob", []lib.Person{})
	carlos := lib.NewPerson(2, "Carlos", []lib.Person{})
	charlie := lib.NewPerson(4, "Charlie", []lib.Person{})
	eve := lib.NewPerson(7, "Eve", []lib.Person{})

	aTeam := lib.NewPerson(90, "The A-Team", []lib.Person{alice, bob, carlos})
	cTeam := lib.NewPerson(92, "The C-Team", []lib.Person{charlie, eve, aTeam})

	t.Log("It should be true if a Person is inside the array")

	array := []lib.Person{alice, bob, carlos, charlie, aTeam}

	var fibTests = []struct {
		n        lib.Person // input
		expected bool       // expected result
	}{
		{alice, true},
		{eve, false},
		{aTeam, true},
		{cTeam, false},
	}

	for _, tt := range fibTests {
		actual := lib.ContainsPerson(array, tt.n)
		if actual != tt.expected {
			t.Errorf("ContainsPerson(array, %s): expected %t, actual %t", tt.n.DisplayName, tt.expected, actual)
		}
	}
}
func TestContainsInt(t *testing.T) {
	t.Log("It should be true if a number is inside the array")

	array := []int32{1, 2, 3, 4, 5, 6}

	var fibTests = []struct {
		n        int32 // input
		expected bool  // expected result
	}{
		{3, true},
		{7, false},
		{0, false},
		{-1, false},
	}

	for _, tt := range fibTests {
		actual := lib.ContainsInt(array, tt.n)
		if actual != tt.expected {
			t.Errorf("ContainsInt(array, %d): expected %t, actual %t", tt.n, tt.expected, actual)
		}
	}
}
