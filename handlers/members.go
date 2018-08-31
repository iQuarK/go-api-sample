package handlers

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/iquark/go-api-sample/lib"
	"github.com/iquark/go-api-sample/middlewares"
	"net/http"
)

// MembersHandlers interface definition
type MembersHandlers interface {
	GetMembersHandler(w http.ResponseWriter, r *http.Request)
}

// GetMembersHandler returns the list of direct and indirect members of a team
func GetMembersHandler(w http.ResponseWriter, r *http.Request) {
	// parameters
	vars := mux.Vars(r)
	id := vars["id"]

	// database
	db := context.Get(r, "DB").(*middlewares.DB)
	members, _ := db.GetPerson(id)

	alice := lib.NewPerson(0, "Alice", []lib.Person{})
	bob := lib.NewPerson(1, "Bob", []lib.Person{})
	carlos := lib.NewPerson(2, "Carlos", []lib.Person{})
	charlie := lib.NewPerson(4, "Charlie", []lib.Person{})
	eve := lib.NewPerson(7, "Eve", []lib.Person{})

	aTeam := lib.NewPerson(90, "The A-Team", []lib.Person{alice, bob, carlos})
	cTeam := lib.NewPerson(92, "The C-Team", []lib.Person{charlie, eve, aTeam})

	listMembers := lib.GetMembers(cTeam)

	JSONResponse(listMembers, w)
}
