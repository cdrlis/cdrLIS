package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LASpatialUnitGroup Spatial unit group
type LASpatialUnitGroup struct {
	SugID          shared.Oid
	HierarchyLevel int
	Label          string
	Name           string
	ReferencePoint string
}
