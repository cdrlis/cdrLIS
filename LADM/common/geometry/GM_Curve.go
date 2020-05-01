package geometry

import (
	"database/sql/driver"

	"github.com/paulsmith/gogeos/geos"
)

// GMCurve Point geometry type
type GMCurve GMObject

// Value converts the given GMCurve struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g GMCurve) Value() (driver.Value, error) {
	return (GMObject)(g).Value()
}

// Scan converts the hexadecimal representation of geometry into the given GMCurve
// struct. Implements Scanner interface for use with database operations.
func (g *GMCurve) Scan(value interface{}) error {
	return (*GMObject)(g).Scan(value)
}

// AsGeometry Returns underlying geometry type
func (g *GMCurve) AsGeometry() *geos.Geometry {
	return ((*GMObject)(g)).AsGeometry()
}
