package hello

import "testing"

func Test_Name(t *testing.T) {

	if hello() != "Hello!" {
		t.Errorf("Expected %s received %s", "Hello!", hello())
	}
}
