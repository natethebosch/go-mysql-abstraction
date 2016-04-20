package mysql

import (
	"testing"
	"github.com/ziutek/mymysql/mysql"
	"time"
	"errors"
)

func setup() {
	// nothing
	SetConnectionInfo("127.0.0.1:3306", "root", "Theuser1", "blue-giraffe")
}

func teardown(){
	isInit = false

	// empty conn before closing
	done := false
	for !done {
		select {
		case <-conn:
		case <- time.After(100 * time.Millisecond):
			done = true
		}
	}

	close(conn)
}

func ensureConnectionValid(conn mysql.Conn, t *testing.T) (bool){
	if !conn.IsConnected(){
		t.Error(errors.New("Connection is not connected"))
		return false
	}

	return true
}

func TestConnectionPoolHasLimit(t *testing.T) {
	setup()

	var cn mysql.Conn

	for i := 0; i < connectionPool; i++ {
		cn = getConnection()
		ensureConnectionValid(cn, t)
	}

	// if we can get extra connections that means that it's not stopping at the limit
	c := make(chan mysql.Conn)
	go func(){
		c <- getConnection()
	}()

	select {
	case cn = <-c:
		t.Error("Should not be able to get connection after connectionPool is exhausted")
		t.Fail()
	case <-time.After(300 * time.Millisecond):

	}

	teardown()
}

func TestConnectionPoolWillRegenerate(t *testing.T){
	setup()

	for i := 0; i < connectionPool * 3; i++ {

		conn := getConnection()

		if !ensureConnectionValid(conn, t) {
			t.Fail()
			return
		}

		conn.Close()

		returnConnection(conn)
	}

	teardown()
}

func TestConnectionPoolWillReuseOrRegenerate(t *testing.T){
	setup()

	for i := 0; i < connectionPool * 3; i++ {

		conn := getConnection()

		if !ensureConnectionValid(conn, t) {
			t.Fail()
			return
		}

		// close every other connection
		if i % 2 == 0 {
			conn.Close()
		}

		returnConnection(conn)
	}

	teardown()
}