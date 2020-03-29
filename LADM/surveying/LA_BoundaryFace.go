package surveying

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LABoundaryFace Boundary face string
type LABoundaryFace struct {
	shared.VersionedObject

	BfID           shared.Oid
	Geometry       *GMMultiSurface
	LocationByText *string
}

// GMMultiSurface Multi surface
type GMMultiSurface string // TODO external package
