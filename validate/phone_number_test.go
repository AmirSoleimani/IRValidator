package validate

import (
	"errors"
	"reflect"
	"testing"
)

func TestMobilePhone(t *testing.T) {
	tables := []struct {
		value       string
		want        error
		PhoneNumber string
	}{
		{value: "2130396217", want: errors.New(""), PhoneNumber: ""},
		{value: "+9893909200111", want: errors.New(""), PhoneNumber: ""},
		{value: "0098939092011", want: errors.New(""), PhoneNumber: ""},
		{value: "09390920011", want: nil, PhoneNumber: "9390920011"},
		{value: "+989390920011", want: nil, PhoneNumber: "9390920011"},
		{value: "00989390920011", want: nil, PhoneNumber: "9390920011"},
	}
	for _, data := range tables {
		PhoneNumber, err := MobilePhone(data.value)
		if reflect.TypeOf(err) != reflect.TypeOf(data.want) &&  PhoneNumber == data.PhoneNumber{
			t.Errorf("nationalid was incorrect, got : %s, want: nil. value : %s ", err, data.value)
		}
	}
}
