package lib

// Teams interface definition
type Teams interface {
	GetTeams(person Person, list []Person) ([]Person, error)
}

// GetTeams accepts a person and a list of all people/teams and returns a list of all the teams of which that person is a member
func GetTeams(person Person, list []Person) ([]Person, error) {
	var teams []Person

	for _, item := range list {
		if item.IsTeam {
			if ContainsPerson(GetMembers(item), person) {
				teams = append(teams, item)
			}
		}
	}

	return teams, nil
}

// const listLength = list.length;
// let teams = [];

// for (let i=0; i < listLength; i++) {
// 	const item = list[i];

// 	if (item.isTeam) {
// 		if (getMembers(item).indexOf(person) > -1) {
// 			teams = [ ...teams, item ];
// 		}
// 	}
// }

// return teams;
