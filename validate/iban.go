package validate

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

//charNumber , char number
var charNumber = map[string]int{
	"A": 10,
	"B": 11,
	"C": 12,
	"D": 13,
	"E": 14,
	"F": 15,
	"G": 16,
	"H": 17,
	"I": 18,
	"J": 19,
	"K": 20,
	"L": 21,
	"M": 22,
	"N": 23,
	"O": 24,
	"P": 25,
	"Q": 26,
	"R": 27,
	"S": 28,
	"T": 29,
	"U": 30,
	"V": 31,
	"W": 32,
	"X": 33,
	"Y": 34,
	"Z": 35,
}

//IBAN Check IBAN Number
func IBAN(value string) error {
	if len(value) >= 4 {
		//get the first four
		firstFour := value[:4]

		//replace the first four
		value = strings.Replace(value, firstFour, "", -1)

		//set char number to the first four
		for _, r := range firstFour[:2] {
			s := string(r)
			firstFour = strings.Replace(firstFour, s, strconv.Itoa(charNumber[s]), -1)
		}

		//add end of string
		value = value + firstFour

		//convert string to bigInt
		i, err := big.NewInt(0).SetString(value, 10)
		if !err {
			return errors.New("IBAN > convert problem")
		}

		//set divider value
		j := big.NewInt(97)

		//div->mod
		k := big.NewInt(0).Mod(i, j)

		//result
		if k.Cmp(big.NewInt(1)) != 0 { //Correct
			return errors.New("IBAN > problem")
		}

	}

	return nil
}
