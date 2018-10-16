package main

import (
	"fmt"
	"github.com/IRValidator-go/validate"
)

func main() {
	// CardNumber check
	if err := validate.CardNumber("6395991166685826"); err != nil {
		fmt.Println(err)
	}

	// NationalID check
	if err := validate.NationalID("1111111111"); err != nil {
		fmt.Println(err)
	}

	// IBAN check
	if err := validate.IBAN("IR340570030280001175105001"); err != nil {
		fmt.Println(err)
	}

	// MobilePhone check
	if PhoneNumber, err := validate.MobilePhone("+989121111111"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(PhoneNumber)
	}
}
