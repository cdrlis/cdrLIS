package common

import (
	"database/sql/driver"
	"errors"
	"regexp"
)

// Oid Object identifier
type Oid struct {
	LocalID   string
	Namespace string
}

// Value Returs Oid value
func (id *Oid) Value() (driver.Value, error) {
	return "(" + id.LocalID + ", " + id.Namespace + ")", nil
}

// Scan Reads Oid
func (id *Oid) Scan(value interface{}) error {

	if value == nil {
		id = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot convert database value to geometry")
	}

	str := string(bytes)
	re := regexp.MustCompile("\\((.*?),(.*?)\\)") // TODO Better regex is needed
	match := re.FindStringSubmatch(str)

	oid := Oid{Namespace: match[1], LocalID: match[2]}
	*id = oid

	return nil
}
