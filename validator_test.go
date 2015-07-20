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

	person := new(P1)
	if err := Validate(person, values); err != nil {
		t.Errorf("%v validate error due to %s", *person, err)
	}
}
