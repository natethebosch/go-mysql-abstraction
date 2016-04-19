package sqlAdapter

import (
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native" // Native engine
	"fmt"
)

const bufSize = 20

type sqlAdapter struct{
	db mysql.Conn
	pipe chan mysql.Row
	query string
}

// setup the struct and return a channel to the db results
func Query(query string) (<-chan mysql.Row) {
	db := mysql.New("tcp", "", "127.0.0.1:3306", "root", "Theuser1", "test")

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	s := sqlAdapter{}

	s.db = db
	s.query = formatQuery(query)
	s.pipe = make(chan mysql.Row, bufSize)

	go s.bufferFetcher()

	return s.pipe
}

// set the query that the pipe will use to fetch results from
// in query ensure that LIMIT %d,%d can be appended to it for paging
func formatQuery(query string) (string) {
	_qry := query + " LIMIT %d,%d"
	fmt.Print("Created new query '" + _qry + "'\n")

	return _qry
}

func (s sqlAdapter) bufferFetcher() {

	pos := 0
	end := false

	for !end {

		rows, _, err := s.db.Query(fmt.Sprintf(s.query, pos, bufSize))
		if err != nil {
			panic(err)
		}

		for _, row := range rows {
			s.pipe<-row
			pos++
		}
	}
}
