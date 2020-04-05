package spatialunit

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/external"
)

//
// Spatial Unit::LA_SpatialUnit - single area (or multiple areas) of land (4.1.9) and/or water, or a single volume
// (or multiple volumes) of space
//
// An instance of class LA_SpatialUnit is a spatial unit. A spatial unit may be associated to zero or more [0..*]
// basic administrative units (i.e. the spatial unit may be used to describe the extent – part of – a basic
// administrative unit). A spatial unit may be associated to zero or one [0..1] levels (i.e. a spatial unit can be
// associated to a property level). A spatial unit cannot be associated to more than one level. A spatial unit may
// be associated to zero or more [0..*] spatial unit groups (i.e. a spatial unit can be associated to a subdivision
// and also to school district). A spatial unit can be spatially related, through a required relationship, to zero or
// more [0..*] other spatial units (i.e. creates an explicit spatial relationship between two spatial units when the
// geometry is missing or inaccurate to provide reliable implicit results). Spatial units do not need to be related
// explicitly. A spatial unit can be associated to zero or more [0..*] spatial sources. A spatial unit can form part of
// v0..1 other spatial unit. A spatial unit can include 0..* other spatial units. Spatial units can be further
// specialized
// into building units (6.5.3) or utility networks (6.5.4); see Figure 11

type LASpatialUnit struct {
	common.VersionedObject
	ExtAddressID    *external.ExtAddress
	Area            *LAAreaValue
	Dimension       *LADimensionType
	Label           *string
	ReferencePoint  *GMPoint // TODO
	SuID            common.Oid
	SurfaceRelation *LASurfaceRelationType
}

// Should they be defined in the interface LASpatialUniter ?!

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

// LAAreaValue Area value
type LAAreaValue struct {
	AreaSize Area
	Type     LAAreaType
}

// Area Area
type Area int64

// LAAreaType Area type
type LAAreaType int

const (
	// DefaultArea Default Area type
	DefaultArea LAAreaType = iota
	OfficialArea
	NonOfficialArea
	CalculatedArea
	SurveyedArea
)

// LADimensionType Dimension type
type LADimensionType int

const (
	D0 LADimensionType = iota
	D1
	D2
	D3
	Liminal
)

//
// LA_SurfaceRelationType: the LA_SurfaceRelationType code list includes all the various surface relation
// types, such as above or below surface, used in a specific land administration profile implementation. The
// LA_SurfaceRelationType code list is required only if the attribute surfaceRelation in LA_SpatialUnit class
// is implemented. The code list shall provide a complete list of all codes with a name and description.
//
type LASurfaceRelationType int

const (
	Mixed LASurfaceRelationType = iota
	Below
	Above
	OnSurface
)
