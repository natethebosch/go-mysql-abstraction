package mysql

import (
	"github.com/ziutek/mymysql/mysql"
)


// set the query that the pipe will use to fetch results from
// in query ensure that LIMIT %d,%d can be appended to it for paging
func formatQuery(query string) (string) {
	_qry := query + " LIMIT %d,%d"

	return _qry
}

// setup the struct and return a channel to the db results.
// creates a new connection
func BulkQuery(query string) (<-chan mysql.Row) {


	s := sqlAdapter{}

	s.db = getConnection()
	s.query = formatQuery(query)
	s.pipe = make(chan mysql.Row, bufSize)

	go s.stream()

	return s.pipe
}