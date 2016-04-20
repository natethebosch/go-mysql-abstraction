package mysql

import (
	"github.com/ziutek/mymysql/mysql"
	"fmt"
)

type sqlAdapter struct{
	db mysql.Conn
	pipe chan mysql.Row
	query string
}

// worker thread to stream the results
func (s sqlAdapter) stream() {

	pos := 0
	end := false

	for !end {

		rows, _, err := s.db.Query(fmt.Sprintf(s.query, pos, bufSize))

		if err != nil {
			returnConnection(s.db)
			close(s.pipe)
			return
		}

		for _, row := range rows {
			s.pipe<-row
			pos++
		}

		if len(rows) == 0 {
			end = true
		}
	}

	returnConnection(s.db)
}