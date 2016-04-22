package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type dbResponse struct {
	rows *sql.Rows
	err  error
}

type dbQuery struct {
	query     string
	arguments []interface{}
	response  chan *dbResponse
}

type dbClass struct {
	queryChan  chan *dbQuery
	connection *sql.DB
}

func (d *dbClass) answer() {
	for {
		q := <-d.queryChan
		r := new(dbResponse)
		r.rows, r.err = d.connection.Query(q.query, q.arguments...)
		q.response <- r
	}
}

func (d *dbClass) query(q string, args ...interface{}) (*sql.Rows, error) {
	qStruct := new(dbQuery)
	qStruct.query = q
	qStruct.arguments = args
	qStruct.response = make(chan *dbResponse, 1)
	d.queryChan <- qStruct
	response := <-qStruct.response
	return response.rows, response.err
}

var connection *dbClass

func ConnectToDatabase(user, database string) error {
	connectionString := fmt.Sprintf("user=%s dbname=%s host=/var/run/postgresql", user, database)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	connection = new(dbClass)
	connection.connection = db
	connection.queryChan = make(chan *dbQuery, 10)
	go connection.answer()
	return nil
}
