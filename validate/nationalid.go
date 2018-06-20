package validate

import (
	"errors"
	"fmt"
	"strconv"
)

//NationalID national id validation
func NationalID(value string) error {
	valueCount := len(value)
	if valueCount < 8 {
		return errors.New("nationalid < 8")
	}

	// check len and add zero
	if valueCount >= 8 && valueCount < 10 {
		// add zero 0...
		for i := 0; i < 10-valueCount; i++ {
			value = fmt.Sprintf("0%s", value)
		}
	}

	// convert string to int slice
	var valueSlices []uint8
	for _, s := range value {
		i, err := strconv.ParseInt(string(s), 10, 8)
		if err != nil {
			return errors.New("value problem")
		}
		valueSlices = append(valueSlices, uint8(i))
	}

	// check #1 > *position
	s, err := calculateNationalIDNumbers(&valueSlices)
	if err != nil {
		return err
	}

	// check #2
	s %= 11
	l := valueSlices[len(value)-1]
	if (s < 2 && s != int(l)) || (s >= 2 && int(l) != 11-s) {
		return errors.New("nationalid is not valid")
	}

	return nil
}

//calculateNationalIDNumbers Yo
func calculateNationalIDNumbers(valueSlices *[]uint8) (sum int, err error) {
	var i uint8
	for i = 10; i >= 2; i-- {
		sum += int((i * (*valueSlices)[10-i]))
	}
	return sum, nil
}
