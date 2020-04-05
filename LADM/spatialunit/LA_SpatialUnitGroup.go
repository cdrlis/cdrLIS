package spatialunit

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/paulsmith/gogeos/geos"
)

//
// Spatial Unit::LA_SpatialUnitGroup
//
// Any number of spatial units (4.1.23), considered as an entity.
// An instance of class LA_SpatialUnitGroup is a spatial unit group. A spatial unit group is made of one or more
// [1..*] parts/elements (which can be spatial units, or spatial unit groups, or a combination of spatial units and
// spatial unit groups). A spatial unit group is part of zero or one [0..1] larger spatial unit group, which again can
// even be part of zero or one [0..1] larger spatial unit group, and so on. See Figure 11.

type LASpatialUnitGroup struct {
	common.VersionedObject
	HierarchyLevel int
	Label          *string
	Name           *string
	ReferencePoint *GMPoint
	SugID          common.Oid
}
