package models

import (
	"fmt"
	"log"
	"os"

	// MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DataStore defines the interface to access the data from the DB
type DataStore interface {
	PeopleDataStore
	MembersDataStore
}

// DB to call the Database
type DB struct {
	*sqlx.DB
}

// DBAccessData to access the Database
type DBAccessData struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string
}

func dbAccessDetails() DBAccessData {
	return DBAccessData{os.Getenv("RDS_USERNAME"), os.Getenv("RDS_PASSWORD"), os.Getenv("RDS_HOSTNAME"), os.Getenv("RDS_PORT"), os.Getenv("RDS_DB_NAME")}
}

func connectionString() string {
	data := dbAccessDetails()

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.Database)
}

func mysqlConnectionString() string {
	data := dbAccessDetails()

	return fmt.Sprintf("mysql://%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", data.User, data.Password, data.Host, data.Port, data.Database)
}

/*
NewDB creates an instance to access the DB
Environment variables to access the DataBase:
 - DB_HOST_API: host to access the API database.
 - DB_PORT_API: port of the host.
 - DB_USER_API: user.
 - DB_PASSWORD_API: password.
 - DB_NAME_API: database name.
*/
func NewDB() (*DB, error) {
	db, err := sqlx.Open("mysql", connectionString())

	if err != nil {
		log.Println("Error connecting to database")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Database not reachable")
		return nil, err
	}

	return &DB{db}, nil
}
