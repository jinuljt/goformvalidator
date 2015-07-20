package goformvalidator

import "testing"

type P1 struct {
	Name string `schema:name`
	Age  int    `schema:age`
}

type P2 struct {
	Name string `validate:"max=2" schema:name`
	Age  int    `validate:"max=50" schema:age`
}

func TestP1(t *testing.T) {
	values := map[string][]string{
		"name": {"John"},
		"age":  {"999"},
	}

	p1 := new(P1)
	if err := Validate(p1, values); err != nil {
		t.Errorf("%v validate error due to %s", *p1, err)
	}

	p2 := new(P2)
	if err := Validate(p2, values); err == nil {
		t.Errorf("%v validate should error", *p2)
	}
}
