package models

// Member DB definition, row returned from member table
type Member struct {
	PersonID uint `json:"person_id"`
	MemberID uint `json:"member_id"`
}

// MembersDataStore interface definition for members
type MembersDataStore interface {
	GetMembers(teamID string) ([]Member, error)
}

// GetMembers retrieves the full list of members of a team
func (db *DB) GetMembers(teamID string) ([]Member, error) {
	members := []Member{}

	// fetch all members from the db
	err := db.Select(&members, "SELECT person_id, member_id FROM member m JOIN person p ON p.id=person_id")

	if err != nil {
		panic(err)
	}

	switch {
	case err != nil:
		return nil, err
	default:
		return members, nil
	}
}
