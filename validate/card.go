package validate

import (
	"errors"
	"strconv"
)

var firstNumValid = [3]uint8{
	6,
	5,
	4,
}

//CardNumber card number validation
func CardNumber(value string) error {
	// check length
	if len(value) != 16 {
		return errors.New("card length < 16 :)")
	}

	// convert string to int slice
	var valueSlices []uint8
	for _, s := range value {
		i, err := strconv.ParseInt(string(s), 10, 8)
		if err != nil {
			return errors.New("string problem")
		}
		valueSlices = append(valueSlices, uint8(i))
	}

	// check #1 > first number
	err := checkFirstCardNumber(valueSlices[0])
	if err != nil {
		return err
	}

	// check #2 > calcualte *slices < 10
	sum := calculateCardNumbers(&valueSlices)

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
func calculateCardNumbers(valueSlices *[]uint8) (sum uint8) {
	for i := 0; i < len(*valueSlices); i++ {
		var x uint8
		if i%2 == 0 {
			x = (*valueSlices)[i] * 2
			if x > 10 {
				x -= 9
			}
		} else {
			x = (*valueSlices)[i]
		}
		sum += x
	}

	return sum
}
