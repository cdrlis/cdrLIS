package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/common"

//
// Spatial Unit::LA_RequiredRelationshipSpatialUnit
//
// An instance of association class LA_RequiredRelationshipSpatialUnit is a required relationship between
// spatial units, see Figures 11 and 12. A required relationship between spatial units can be associated to zero
// or more [0..*] spatial sources to provide supporting documentation for the explicit relationhip.

type LARequiredRelationshipSpatialUnit struct {
	common.VersionedObject
	su1          *LASpatialUnit
	su2          *LASpatialUnit
	relationship iso19125Type
}

type iso19125Type int

const (
	ST_Equals iso19125Type = iota
	ST_Disjoint
	ST_Intersects
	ST_Touches
	ST_Crosses
	ST_Within
	ST_Contains
	ST_Overlaps
)
