package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LASpatialUnit Spatil unit
type LASpatialUnit struct {
	ExtAddressID    *shared.Oid
	Area            *LAAreaValue
	Dimension       *LADimensionType
	Label           *string
	ReferencePoint  *GMPoint
	SuID            shared.Oid
	SurfaceRelation *LASurfaceRelationType
	Volume          *LAVolumeValue
}

func (su LASpatialUnit) areaClosed() bool {
	return true
}

func (su LASpatialUnit) volumeClosed() bool {
	return true
}

func (su LASpatialUnit) computeArea() string {
	return ""
}

func (su LASpatialUnit) computeVolume() string {
	return ""
}

func (su LASpatialUnit) createArea() string {
	return ""
}

func (su LASpatialUnit) createVolume() string {
	return ""
}

// GMPoint Point
type GMPoint string // TODO external package

// LADimensionType Dimension type
type LADimensionType int

const (
	// DefaultDimension Default dimension type
	DefaultDimension LADimensionType = 0
)

// LASurfaceRelationType Surface relation type
type LASurfaceRelationType int

const (
	// DefaultSurfaceRealtion Default surface realtion type
	DefaultSurfaceRealtion LASurfaceRelationType = 0
)
