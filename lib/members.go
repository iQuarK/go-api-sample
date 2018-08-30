package lib

// Members interface definition
type Members interface {
	GetMembers(person Person) []Person
}

// GetMembers facade of getMembersVisited that returns all the people (not teams) who are direct and indirect members of a team
func GetMembers(team Person) []Person {
	return getMembersVisited(team, []uint{})
}

func getMembersVisited(team Person, visited []uint) []Person {
	members := []Person{}

	for _, member := range team.Members {
		if !member.IsTeam {
			members = append(members, member)
		} else {
			if !ContainsNumber(visited, member.ID) {
				visited = append(visited, member.ID)
				members = append(members, getMembersVisited(member, visited)...)
			}
		}
	}

	return members
}
