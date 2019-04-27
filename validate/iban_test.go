package validate

import (
	"errors"
	"reflect"
	"testing"
)

func TestIBAN(t *testing.T) {
	tables := []struct {
		value string
		want  error
	}{
		{value: "UU27053000000100324200001", want: errors.New("")},
		{value: "DE27053000000100324200001", want: errors.New("")},
		{value: "IR۲۷۰۵۳۰۰۰۰۰۰۱۰۰۳۲۴۲۰۰۰۰۱", want: errors.New("")},
		{value: "IR530540103047000634411601", want: nil},
		{value: "IR340570030280001175105001", want: nil},
		{value: "Ir340570030280001175105001", want: nil},
	}
	for _, data := range tables {
		err := IBAN(data.value)
		if reflect.TypeOf(err) != reflect.TypeOf(data.want) {
			t.Errorf("IBAN was incorrect, got : %s, want: nil. value : %s ", err, data.value)
		}
	}
}
