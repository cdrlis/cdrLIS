package geometry

import (
	"database/sql/driver"

	"github.com/paulsmith/gogeos/geos"
)

// GMMultiCurve Point geometry type
type GMMultiCurve GMObject

// Value converts the given GMMultiCurve struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g *GMMultiCurve) Value() (driver.Value, error) {
	return (*GMObject)(g).Value()
}

// Scan converts the hexadecimal representation of geometry into the given GMMultiCurve
// struct. Implements Scanner interface for use with database operations.
func (g *GMMultiCurve) Scan(value interface{}) error {
	return (*GMObject)(g).Scan(value)
}

// AsGeometry Returns underlying geometry type
func (g *GMMultiCurve) AsGeometry() *geos.Geometry {
	return ((*GMObject)(g)).AsGeometry()
}
