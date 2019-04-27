package validate

import (
	"errors"
	"math/big"
	"strconv"
	"strings"
)

//IBAN Check IBAN Number
func IBAN(value string) error {

	value = strings.ToUpper(value)

	if len(value) >= 4 {
		//get the first four
		firstFour := value[:4]

		//replace the first four
		value = strings.Replace(value, firstFour, "", -1)

		//set char number to the first four
		for _, r := range firstFour[:2] {
			s := string(r)
			firstFour = strings.Replace(firstFour, s, strconv.Itoa(int([]rune(s)[0])-55), -1)
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
