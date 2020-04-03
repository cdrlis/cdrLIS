package surveying

import "github.com/cdrlis/cdrLIS/LADM/common"

// LABoundaryFace Boundary face string
type LABoundaryFace struct {
	common.VersionedObject

	BfID           common.Oid
	Geometry       *GMMultiSurface
	LocationByText *string
}

// GMMultiSurface Multi surface
type GMMultiSurface string // TODO external package
