package mysql

import (
	"testing"
	"github.com/ziutek/mymysql/mysql"
	"time"
)

func TestStreaming(t *testing.T){
	setup()

	stream := BulkQuery("SELECT id,col1 FROM `test` WHERE id=1 OR id=2 ORDER BY id ASC")

	for i := 0; i < 2; i++ {

		// receive rows
		select {
		case <-stream:
		case <-time.After(10 * time.Millisecond):
			t.Fail()
		}
	}

	// check that there isn't any extra values
	select {
	case <-stream:
		t.Fail()
	case <-time.After(10 * time.Millisecond):
	}
}

func TestStreamingSequence(t *testing.T){
	setup()

	stream := BulkQuery("SELECT id,col1 FROM `test` WHERE id=1 OR id=2 ORDER BY id ASC")
	var rw mysql.Row

	for i := 1; i <= 2; i++ {
		rw = <-stream

		if rw.Int(0) != i {
			t.Fail()
		}
	}
}

func TestStreamingMultiSegment(t *testing.T){
	setup()

	stream := BulkQuery("SELECT id,col1 FROM `test` WHERE id < 26 ORDER BY id ASC")
	var rw mysql.Row

	for i := 1; i < 26; i++ {
		rw = <-stream

		if rw.Int(0) != i {
			t.Fail()
		}
	}
}

func TestStreamingInvalidQuery(t *testing.T){
	setup()

	stream := BulkQuery("SELECT INVALID QUERY")

	select{
	case r := <-stream:
		if r != nil {
			t.Fail()
		}
	case <-time.After(500 * time.Millisecond):
	}
}
