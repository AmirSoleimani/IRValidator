package validate

import (
	"errors"
	"reflect"
	"testing"
)

func TestNationalID(t *testing.T) {
	tables := []struct {
		value string
		want  error
	}{
		{value: "2130396217", want: errors.New("")},
		{value: "0000000001", want: errors.New("")},
		{value: "abcd1234", want: errors.New("")},
		{value: "امیرجونامیرجون", want: errors.New("")},
		{value: "۲۱۳۰۳۹۶۲۱۶", want: errors.New("")},
		{value: "", want: errors.New("")},
		{value: "-123456789", want: errors.New("")},
		{value: "-1234567890", want: errors.New("")},
		{value: "4608968882", want: nil},
		{value: "1111111111", want: nil},
		{value: "939092001", want: nil},
		{value: "0939092001", want: nil},
	}
	for _, data := range tables {
		err := NationalID(data.value)
		if reflect.TypeOf(err) != reflect.TypeOf(data.want) {
			t.Errorf("nationalid was incorrect, got : %s, want: nil. value : %s ", err, data.value)
		}
	}
}
