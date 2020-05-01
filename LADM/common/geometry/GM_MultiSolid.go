package geometry

import (
	"database/sql/driver"

	"github.com/paulsmith/gogeos/geos"
)

// GMMultiSolid Point geometry type
type GMMultiSolid GMObject

// Value converts the given GMMultiSolid struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g GMMultiSolid) Value() (driver.Value, error) {
	return (GMObject)(g).Value()
}

// Scan converts the hexadecimal representation of geometry into the given GMMultiSolid
// struct. Implements Scanner interface for use with database operations.
func (g *GMMultiSolid) Scan(value interface{}) error {
	return (*GMObject)(g).Scan(value)
}

// AsGeometry Returns underlying geometry type
func (g *GMMultiSolid) AsGeometry() *geos.Geometry {
	return ((*GMObject)(g)).AsGeometry()
}
