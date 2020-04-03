package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/common"

// LASpatialUnitGroup Spatial unit group
type LASpatialUnitGroup struct {
	HierarchyLevel int
	Label          *string
	Name           *string
	ReferencePoint *GMPoint
	SugID          common.Oid
}
