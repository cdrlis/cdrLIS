package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
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
	ReferencePoint  *geometry.GMPoint
	SuID            common.Oid
	SurfaceRelation *LASurfaceRelationType

	Baunit      []LABAunit                          // suBaunit
	SuHierarchy []LASpatialUnit                     // suHierarchy
	RelationSu  []LARequiredRelationshipSpatialUnit // relationSu
	Level       *LALevel                            // suLevel
	Whole       []LASpatialUnitGroup                // suSuGroup

	MinusBfs []LABoundaryFaceString // minus
	PlusBfs  []LABoundaryFaceString // plus
}

func (LASpatialUnit) TableName() string {
	return "LA_SpatialUnit"
}

func (su LASpatialUnit) AreaClosed() bool {
	multiSurface := su.CreateArea()
	closed, _ := multiSurface.AsGeometry().IsClosed()
	return closed
}

func (su LASpatialUnit) ComputeArea() LAAreaValue {
	var av LAAreaValue
	multiSurface := su.CreateArea()
	area, _ := multiSurface.AsGeometry().Area()
	av.AreaSize, av.Type = Area(area), CalculatedArea
	return av
}

func (su LASpatialUnit) CreateArea() geometry.GMMultiSurface {
	var ms geometry.GMMultiSurface
	return ms
}

// LAAreaValue Area value
type LAAreaValue struct {
	AreaSize Area
	Type     LAAreaType
}

// Area Area
type Area float64

//
// LA_AreaType: the LA_AreaType code list includes all the various area types, such as official or
// calculated, used in a specific land administration profile implementation. The LA_AreaType code list is
// required to implement the LA_AreaValue data type. The code list shall provide a complete list of all codes
// with a name and description.
//
type LAAreaType string

const (
	OfficialArea    LAAreaType = "officialArea"
	NonOfficialArea            = "nonOfficialArea"
	CalculatedArea             = "calculatedArea"
	SurveyedArea               = "surveyedArea"
)

//
// LA_DimensionType: the LA_DimensionType code list includes all the various dimension types, such as
// 2D or 3D, used in a specific land administration profile implementation. The LA_DimensionType code list
// is required only if the attribute dimension in LA_SpatialUnit class is implemented. The code list shall
// provide a complete list of all codes with a name and description.
//
type LADimensionType string

const (
	D0      LADimensionType = "0D"
	D1                      = "1D"
	D2                      = "2D"
	D3                      = "3D"
	Liminal                 = "laminal"
)

//
// LA_SurfaceRelationType: the LA_SurfaceRelationType code list includes all the various surface relation
// types, such as above or below surface, used in a specific land administration profile implementation. The
// LA_SurfaceRelationType code list is required only if the attribute surfaceRelation in LA_SpatialUnit class
// is implemented. The code list shall provide a complete list of all codes with a name and description.
//
type LASurfaceRelationType string

const (
	MixedSRT  LASurfaceRelationType = "mixed" // Due to conflict with LALevelContentType, Mixed can't be used.
	Below                           = "below"
	Above                           = "above"
	OnSurface                       = "onSurface"
)
