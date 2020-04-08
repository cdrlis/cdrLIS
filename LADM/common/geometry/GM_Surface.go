package geometry

import (
	"database/sql/driver"

	"github.com/paulsmith/gogeos/geos"
)

// GMSurface Point geometry type
type GMSurface GMObject

// Value converts the given GMSurface struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g *GMSurface) Value() (driver.Value, error) {
	return (*GMObject)(g).Value()
}

// Scan converts the hexadecimal representation of geometry into the given GMSurface
// struct. Implements Scanner interface for use with database operations.
func (g *GMSurface) Scan(value interface{}) error {
	return (*GMObject)(g).Scan(value)
}

// AsGeometry Returns underlying geometry type
func (g *GMSurface) AsGeometry() *geos.Geometry {
	return ((*GMObject)(g)).AsGeometry()
}
