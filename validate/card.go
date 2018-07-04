package validate

import (
	"errors"
	"strconv"
)

var firstNumValid = [3]uint8{
	54, //6
	53, //5
	42, //4
}

//CardNumber card number validation
func CardNumber(value string) error {
	// check length
	if len(value) != 16 {
		return errors.New("card length < 16 :)")
	}

	// check #1 > first number
	err := checkFirstCardNumber(value[0])
	if err != nil {
		return err
	}

	// check #2 > calcualte *slices < 10
	sum := calculateCardNumbers(&value)

	// check #3 > %10 = 0
	if sum%10 != 0 {
		return errors.New("card number is not valid")
	}

	return nil
}

//checkFirstCardNumber first number in cardnum
func checkFirstCardNumber(firstNum uint8) error {
	for _, v := range firstNumValid {
		if v == firstNum {
			return nil
		}
	}

	return errors.New("first number isn`t valid")
}

//calculateCardNumbers Yo
func calculateCardNumbers(valueSlices *string) (sum uint64) {
	for i := 0; i < len(*valueSlices); i++ {
		v, _ := strconv.ParseUint(string((*valueSlices)[i]), 10, 8)
		var x uint64
		if i%2 == 0 {
			if x = v * 2; x >= 10 {
				x -= 9
			}
		} else {
			x = v
		}
		sum += x
	}

	return sum
}
