package db

import (
	"database/sql"
)

type dbQuery struct {
	query string
	response chan string
}

type dbClass struct {
	queryChan chan *dbQuery
	connection *sql.DB
}
