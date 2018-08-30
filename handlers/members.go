package handlers

import (
	"github.com/gorilla/mux"
	"github.com/iquark/go-api-sample/lib"
	"html"
	"log"
	"net/http"
)

// MembersHandlers interface definition
type MembersHandlers interface {
	GetMembersHandler(w http.ResponseWriter, r *http.Request)
}

// GetMembersHandler returns the list of teams a user belongs to
func GetMembersHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ou := vars["id"]
	log.Printf("teams!!, %q %s", html.EscapeString(r.URL.Path), ou)

	alice := lib.NewPerson(0, "Alice", []lib.Person{})
	bob := lib.NewPerson(1, "Bob", []lib.Person{})
	carlos := lib.NewPerson(2, "Carlos", []lib.Person{})
	charlie := lib.NewPerson(4, "Charlie", []lib.Person{})
	eve := lib.NewPerson(7, "Eve", []lib.Person{})

	aTeam := lib.NewPerson(90, "The A-Team", []lib.Person{alice, bob, carlos})
	cTeam := lib.NewPerson(92, "The C-Team", []lib.Person{charlie, eve, aTeam})

	members := lib.GetMembers(cTeam)

	JSONResponse(members, w)
}
