package geometry

import (
	"database/sql/driver"
)

// GMMultiSurface Point geometry type
type GMMultiSurface GMObject

// Value converts the given GMMultiSurface struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g *GMMultiSurface) Value() (driver.Value, error) {
	return (*GMObject)(g).Value()
}

// Scan converts the hexadecimal representation of geometry into the given GMMultiSurface
// struct. Implements Scanner interface for use with database operations.
func (g *GMMultiSurface) Scan(value interface{}) error {
	return (*GMObject)(g).Scan(value)
}