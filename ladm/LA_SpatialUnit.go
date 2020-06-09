package ladm

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/cdrlis/cdrLIS/ladm/common/geometry"
	"github.com/cdrlis/cdrLIS/ladm/external"
	"github.com/paulsmith/gogeos/geos"
	"regexp"
	"strconv"
	"time"
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

	ID              string                 `gorm:"column:id;primary_key" json:"-"`
	ExtAddressID    *external.ExtAddress   `gorm:"column:extaddressid" json:"extAddressID"`
	Area            *LAAreaValue           `gorm:"column:area" json:"area"`
	Dimension       *LADimensionType       `gorm:"column:dimension" json:"dimension"`
	Label           *string                `gorm:"column:label" json:"label"`
	ReferencePoint  *geometry.GMPoint      `gorm:"column:referencepoint" json:"referencePoint"`
	SuID            common.Oid             `gorm:"column:suid" json:"suID"`
	SurfaceRelation *LASurfaceRelationType `gorm:"column:surfacerelation" json:"surfaceRelation"`

	// suLevel
	LevelID                   string    `gorm:"column:level" json:"-"`
	LevelBeginLifespanVersion time.Time `gorm:"column:levelbeginlifespanversion" json:"-"`
	Level                     *LALevel  `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:LevelID,LevelBeginLifespanVersion" json:"level"`

	BuildingUnit LARequiredRelationshipSpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"building,omitempty"`

	Baunit []SuBAUnit `gorm:"foreignkey:SUID,SUBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion;" json:"baunit"`

	SuHierarchy       *SuHierarchy                        `gorm:"foreignkey:ChildID,ChildBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"hierarchy,omitempty"` // suHierarchy
	RelationSu1       []LARequiredRelationshipSpatialUnit `gorm:"foreignkey:Su1ID,Su1BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"relation1,omitempty"`     // relationSu
	RelationSu2       []LARequiredRelationshipSpatialUnit `gorm:"foreignkey:Su2ID,Su2BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"relation2,omitempty"`     // relationSu
	SpatialUnitGroups []SuSuGroup                         `gorm:"foreignkey:PartID,PartBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"groups,omitempty"`      // suSuGroup
	MinusBfs          []BfsSpatialUnitMinus               `gorm:"foreignkey:SuID,SuBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion;" json:"bfsMinus"`                 // minus
	PlusBfs           []BfsSpatialUnitPlus                `gorm:"foreignkey:SuID,SuBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion;" json:"bfsPlus"`                  // plus
}

func (LASpatialUnit) TableName() string {
	return "LA_SpatialUnit"
}

func (su LASpatialUnit) AreaClosed() bool {
	closed, _ := geos.Must(su.Boundary()).IsClosed()
	return closed
}

func (su LASpatialUnit) ComputeArea() LAAreaValue {
	var av LAAreaValue
	multiSurface := su.CreateArea()
	area, _ := multiSurface.Area()
	av.AreaSize, av.Type = Area(area), CalculatedArea
	return av
}

func (su LASpatialUnit) CreateArea() *geometry.GMMultiSurface {
	msBoundary := geos.Must(su.Boundary())

	tempMultiSurface := geos.Must(geos.EmptyPolygon())
	var ms []*geos.Geometry

	nGeometry, _ := msBoundary.NGeometry()
	for i := 0; i < nGeometry; i++ {
		curve := geos.Must(msBoundary.Geometry(i))
		closed, _ := curve.IsClosed()
		if !closed {
			continue
		}
		surface := geos.Must(geos.NewPolygon(geos.MustCoords(curve.Coords())))
		ms = append(ms, surface)
		tempMultiSurface = geos.Must(tempMultiSurface.Union(surface))
	}
	multiSurface := geos.Must(tempMultiSurface.Clone())

	for _, surface := range ms {
		if related, _ := surface.RelatePat(multiSurface, "2FF1FF212"); related {
			multiSurface = geos.Must(multiSurface.Difference(surface))
		}
	}
	return &geometry.GMMultiSurface{GMObject: geometry.GMObject{Geometry: *multiSurface}}
}

func (su LASpatialUnit) Boundary() (*geos.Geometry, error) {
	msBoundary := geos.Must(geos.NewLineString())
	for _, bfs := range su.PlusBfs {
		var geom *geos.Geometry
		if bfs.Bfs.Geometry != nil {
			geom = &(bfs.Bfs.Geometry.GMObject.Geometry)
		} else {
			continue
		}
		curve := geos.Must(geos.NewLineString())
		nGeometry, err := geom.NGeometry()
		if err != nil {
			continue
		}
		for i := 0; i < nGeometry; i++ {
			curve := geos.Must(geom.Geometry(i))
			curve = geos.Must(curve.Union(geos.Must(geos.NewLineString(geos.MustCoords(curve.Coords())...))))
		}
		msBoundary = geos.Must(msBoundary.Union(curve))
	}
	for _, bfs := range su.MinusBfs {
		var geom *geos.Geometry
		if bfs.Bfs.Geometry != nil {
			geom = &(bfs.Bfs.Geometry.GMObject.Geometry)
		} else {
			continue
		}
		curve := geos.Must(geos.NewLineString())
		nGeometry, err := geom.NGeometry()
		if err != nil {
			continue
		}
		for i := 0; i < nGeometry; i++ {
			lineString := geos.Must(geom.Geometry(i))
			curve = geos.Must(curve.Union(geos.Must(geos.NewLineString(geos.MustCoords(lineString.Coords())...))))
		}
		msBoundary = geos.Must(msBoundary.Union(curve))
	}
	msBoundary = geos.Must(msBoundary.LineMerge())
	return msBoundary, nil
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
	NonOfficialArea LAAreaType = "nonOfficialArea"
	CalculatedArea  LAAreaType = "calculatedArea"
	SurveyedArea    LAAreaType = "surveyedArea"
)

func (area LAAreaValue) Value() (driver.Value, error) {
	return fmt.Sprintf("(%s,%f)",area.Type, area.AreaSize), nil
}

// Scan Reads Oid
func (area *LAAreaValue) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("cannot convert database value to area")
	}

	str := string(bytes)
	re := regexp.MustCompile("\\((.*?),(.*?)\\)")
	match := re.FindStringSubmatch(str)
	areaSize, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return err
	}
	areaValue := LAAreaValue{AreaSize: Area(areaSize), Type: LAAreaType(match[2])}
	*area = areaValue

	return nil
}

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
