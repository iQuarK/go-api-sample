package handlers

import (
	"github.com/gorilla/mux"
	"github.com/iquark/go-api-sample/lib"
	"html"
	"log"
	"net/http"
)

// TeamsHandlers interface definition
type TeamsHandlers interface {
	GetTeamsHander(w http.ResponseWriter, r *http.Request)
}

// GetTeamsHandler returns the list of teams a user belongs to
func GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ou := vars["id"]
	log.Printf("teams!!, %q %s", html.EscapeString(r.URL.Path), ou)

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

	teams, error := lib.GetTeams(alice, people)

	switch {
	case error != nil:
		http.Error(w, error.Error(), http.StatusInternalServerError)
	default:
		JSONResponse(teams, w)
	}
}
