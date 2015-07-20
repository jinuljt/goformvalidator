package goformvalidator

import (
	"github.com/gorilla/schema"
	"gopkg.in/validator.v2"
)

var decoder = schema.NewDecoder()

func Validate(dst interface{}, values map[string][]string) error {
	decoder.Decode(dst, values)
	return validator.Validate(dst)
}
