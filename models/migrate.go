package models

import (
	// "database/sql"
	"log"

	"github.com/mattes/migrate"
	// MySQL driver
	_ "github.com/mattes/migrate/database/mysql"
	// To get the migrations from the file system
	_ "github.com/mattes/migrate/source/file"
)

// NextMigration defines the current version the database is in
const NextMigration int = 4

// DoNextMigration updates the database with the last version
func DoNextMigration() bool {
	return updateMigration(NextMigration)
}

func updateMigration(limit int) bool {
	success := true
	m, err := migrate.New("file://models/migrations/", mysqlConnectionString())

	if err != nil {
		success = false
		panic(err)
	}

	current, dirty, _ := m.Version()

	log.Printf("MIGRATION: current in database %d (dirty %t), next %d", current, dirty, limit)
	if current != uint(limit) {
		step := int(uint(limit) - current)
		if err := m.Steps(step); err != nil {
			if err != migrate.ErrNoChange {
				log.Fatalln("Fatal:", err)
			} else {
				log.Println("Error:", err)
			}
			success = false
		}
	} else {
		if dirty {
			log.Println("MIGRATION: Current version is dirty, forcing current version.")
			if err := m.Force(limit); err != nil {
				if err != migrate.ErrNoChange {
					log.Fatalln("Force Fatal:", err)
				} else {
					log.Println("Force Error:", err)
				}
				success = false
			}
		} else {
			log.Println("MIGRATION: Same version, no need to migrate.")
		}

	}

	return success
}
