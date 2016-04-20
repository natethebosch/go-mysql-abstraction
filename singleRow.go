package mysql

import (
	"github.com/ziutek/mymysql/mysql"
)

// fetches one row from the database
// creates a connection if one is not cached
func QueryOneRow(query string) (mysql.Row) {


	var conn mysql.Conn = getConnection()
	rws, _, err := conn.Query(query)

	if err != nil {
		return nil
	}

	if len(rws) == 0 {
		return nil
	}

	returnConnection(conn)

	return rws[0]
}