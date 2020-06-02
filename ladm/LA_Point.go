package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/cdrlis/cdrLIS/ladm/common/geometry"
)

//
// Surveying and Representation::LA_Point
//
// An instance of class LA_Point is a point. A point may be associated to zero or one [0..1] spatial units (i.e. the
// point may be used as the reference point to describe the position of a spatial unit). A point may be associated
// to zero or more [0..*] boundary faces (i.e. a point may be used to define a vertex of the side of a 3D parcel). A
// point may be associated to zero or more [0..*] boundary face strings (i.e. a point can be used to define the
// start, end or vertex of a boundary). A point should be associated to zero or more [0..*] spatial sources. See
// Figure 12.
//
type LAPoint struct {
	common.VersionedObject

	InterpolationRole LAInterpolationType  `gorm:"column:interpolationrole" json:"interpolationRole"`
	Monumentation     *LAMonumentationType `gorm:"column:monumentation" json:"monumentation"`
	OriginalLocation  *geometry.GMPoint    `gorm:"column:originallocation" json:"originalLocation"`
	PID               common.Oid           `gorm:"column:pid" json:"pID"`
	PointType         LAPointType          `gorm:"column:pointtype" json:"pointType"`
	ProductionMethod  *LILineage           `gorm:"column:productionmethod" json:"productionMethod"`
	TransAndResult    *LATransformation    `gorm:"column:transandresult" json:"transAndResult"`

	Bfs               []PointBfs           `gorm:"foreignkey:PointID,PointBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"bfs,omitempty"`
}

func (LAPoint) TableName() string {
	return "LA_Point"
}

// LILinage Linage
type LILineage string // TODO external package

// LAInterpolationType Interpolation role type
type LAInterpolationType string

const (
	End      LAInterpolationType = "end"
	Isolated LAInterpolationType = "isolated"
	Mid      LAInterpolationType = "mid"
	MidArc   LAInterpolationType = "midarc"
	Start    LAInterpolationType = "start"
)

// LAMonumentationType Monumentation type
type LAMonumentationType string

const (
	Beacon       LAMonumentationType = "beacon"
	Cornserstone LAMonumentationType = "cornserstone"
	Marker       LAMonumentationType = "marker"
	NotMarked    LAMonumentationType = "notMarked"
)

// LAPointType Point type
type LAPointType string

const (
	Control  LAPointType = "control"
	NoSource LAPointType = "noSource"
)
