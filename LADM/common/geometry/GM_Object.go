package geometry

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/paulsmith/gogeos/geos"
)

//
const SRID = "3765" // Croatian SRID

// GMObject Point geometry type
type GMObject struct{
	geos.Geometry
}

// Value converts the given GMObject struct into WKT such that it can be stored in a
// database. Implements Valuer interface for use with database operations.
func (g GMObject) Value() (driver.Value, error) {

	str, err := g.ToWKT()
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

	geometry := GMObject{Geometry:*geom}
	*g = geometry

	return nil
}

func (g *GMObject) MarshalJSON() ([]byte, error) {
	wkt, err := g.ToWKT()
	if err != nil {
		return nil, err
	}
	return json.Marshal(wkt)
}

func (g *GMObject) UnmarshalJSON(data []byte) error {
	str := string(data)
	geom, err := geos.FromWKT(str)
	if err != nil {
		return err
	}
	geometry := GMObject{Geometry:*geom}
	*g = geometry
	return nil
}