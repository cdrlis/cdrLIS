package common

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// Fraction Simple mathematial fraction
type Fraction struct {
	Denominator int `json:"denominator"`
	Numerator   int `json:"numerator"`
}

// Value Returns Oid value
func (fraction Fraction) Value() (driver.Value, error) {
	return fmt.Sprint("(%d,%d)",fraction.Numerator, fraction.Denominator), nil
}

// Scan Reads Oid
func (fraction *Fraction) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot convert database value to geometry")
	}

	str := string(bytes)
	re := regexp.MustCompile("\\((.*?),(.*?)\\)") // TODO Better regex is needed
	match := re.FindStringSubmatch(str)

	tempNumerator, err := strconv.Atoi(match[1])
	if err != nil {
		return err
	}
	tempDenominator, err := strconv.Atoi(match[2])
	if err != nil {
		return err
	}
	temp := Fraction{Numerator: tempNumerator, Denominator: tempDenominator}
	*fraction = temp

	return nil
}