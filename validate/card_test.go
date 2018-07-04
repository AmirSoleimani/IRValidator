package validate

import (
	"errors"
	"reflect"
	"testing"
)

func TestCard(t *testing.T) {
	tables := []struct {
		value string
		want  error
	}{
		{value: "abcdefghi1234555", want: errors.New("")},
		{value: "622106104944", want: errors.New("")},
		{value: "3221061049447982", want: errors.New("")},
		{value: "6221061049447983", want: errors.New("")},
		{value: "۶۲۲۱۰۶۱۰۴۹۴۴۷۹۸۲", want: errors.New("")},
		{value: "6221061049447982", want: nil},
		{value: "6395991166685826", want: nil},
	}
	for _, data := range tables {
		err := CardNumber(data.value)
		if reflect.TypeOf(err) != reflect.TypeOf(data.want) {
			t.Errorf("cardnumber was incorrect, got : %s, want: nil. value : %s ", err, data.value)
		}
	}
}
