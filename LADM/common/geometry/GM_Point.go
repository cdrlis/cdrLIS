package geometry

import (
	"database/sql/driver"
)

// GMPoint Point geometry type
type GMPoint GMObject

// Value converts the given GMPoint struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g *GMPoint) Value() (driver.Value, error) {
	return (*GMObject)(g).Value()
}

// Scan converts the hexadecimal representation of geometry into the given GMPoint
// struct. Implements Scanner interface for use with database operations.
func (g *GMPoint) Scan(value interface{}) error {
	return (*GMObject)(g).Scan(value)
}
