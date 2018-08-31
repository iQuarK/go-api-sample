package models

import (
	"github.com/iquark/go-api-sample/lib"
)

// Person DB definition, row returned from person table, to be moved to the normal Person struct
type Person struct {
	ID          uint   `json:"id"`
	DisplayName string `json:"display_name"`
	Member      uint   `json:"member"`
}

// PeopleDataStore interface definition for people
type PeopleDataStore interface {
	GetPeople() ([]lib.Person, error)
	GetPerson(id string) (Person, error)
}

// GetPeople retrieves the full list of people
func (db *DB) GetPeople() ([]lib.Person, error) {
	people := []Person{}

	// fetch all people with its members from the db
	err := db.Select(&people, "SELECT id, display_name, m.member_id as member FROM person p LEFT JOIN member m ON p.id=m.person_id ORDER BY p.id")

	if err != nil {
		panic(err)
	}

	var peopleMap map[uint]lib.Person

	// maps DB rows into a list of people
	for i := 0; i < len(people); i++ {
		p := people[i]

		if _, exists := peopleMap[p.ID]; !exists {
			peopleMap[p.ID] = lib.Person{ID: p.ID, DisplayName: p.DisplayName}
		}
	}

	// add members into the mapped people
	for i := 0; i < len(people); i++ {
		p := people[i]

		// convert DB rows into list of people
		if person, exists := peopleMap[p.Member]; exists {
			personMap := peopleMap[p.ID]
			personMap.Members = append(personMap.Members, person)
			peopleMap[p.ID] = personMap
		}

	}

	var result []lib.Person
	for _, value := range peopleMap {
		result = append(result, value)
	}

	switch {
	case err != nil:
		return nil, err
	default:
		return result, nil
	}
}

// GetPerson retrieves a Person with all its members
func (db *DB) GetPerson(id string) (Person, error) {
	people, err := db.GetPeople()

	if err != nil {
		panic(err)
	}

	if person, err := people[]

	return person, err
}
