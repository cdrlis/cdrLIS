package surveying

import "github.com/cdrlis/cdrLIS/LADM/common"

// LABoundaryFaceString Boundary face string
type LABoundaryFaceString struct {
	common.VersionedObject

	BfsID          common.Oid
	Geometry       *GMMultiCurve
	LocationByText *string
}

// GMMultiCurve Multi curve
type GMMultiCurve string // TODO external package
