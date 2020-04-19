package geometry

import (
	"database/sql/driver"
	"errors"

	"github.com/paulsmith/gogeos/geos"
)

//
const SRID = "3765" // Croatian SRID

// GMObject Point geometry type
type GMObject geos.Geometry

// Value converts the given GMObject struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g GMObject) Value() (driver.Value, error) {
	geometry := geos.Geometry(g)

	str, err := geometry.ToWKT()
	if err != nil {
		return nil, err
	}

	return "SRID=" + SRID + ";" + str, nil
}

// Scan converts the hexadecimal representation of geometry into the given GMObject
// struct. Implements Scanner interface for use with database operations.
func (g *GMObject) Scan(value interface{}) error {

	if value == nil {
		g = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot convert database value to geometry")
	}

	str := string(bytes)

	geom, err := geos.FromHex(str)
	if err != nil {
		return errors.New("cannot get geometry from hex")
	}

	geometry := GMObject(*geom)
	*g = geometry

	return nil
}

// AsGeometry Returns underlying geometry type
func (g *GMObject) AsGeometry() *geos.Geometry {
	return (*geos.Geometry)(g)
}
