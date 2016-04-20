package mysql

import "testing"

func TestSingleRowQuery(t *testing.T) {
	r := QueryOneRow("SELECT id,col1 FROM `test` where id=1");

	if r == nil {
		t.Fail()
	}
}

func TestSingleRowQueryMultiValue(t *testing.T) {
	r := QueryOneRow("SELECT id,col1 FROM `test` where id=1 OR id=2");

	if r == nil {
		t.Fail()
	}

	// should return first line
	if r.Int(0) == 2 {
		t.Fail()
	}
}

func TestSingleRowQueryNoValue(t *testing.T) {
	r := QueryOneRow("SELECT id,col1 FROM `test` where id=-1");

	if r != nil {
		t.Fail()
	}
}

func TestSingleRowQueryInvalidSyntax(t *testing.T) {
	r := QueryOneRow("SELECT INVALID SYNTAX");

	if r != nil {
		t.Fail()
	}
}
