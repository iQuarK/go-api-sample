# go-api-sample

## Install
Just clone this repository, and get some dependecies:
`$ go get -u github.com/gorilla/mux`
`$ go get -u github.com/gorilla/context`
`$ go get -u github.com/mattes/migrate`
`$ go get -u github.com/jmoiron/sqlx`
`$ go get github.com/go-sql-driver/mysql`

Set up the environment variables for the database:
`$ export RDS_DB_NAME=database_name`
`$ export RDS_USERNAME=user`
`$ export RDS_PASSWORD=password`
`$ export RDS_HOSTNAME=hostname`
`$ export RDS_PORT=8889`

To make it work first build:
`$ go build -o bin/application`

And run the binary:
`$ bin/application`
