package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
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

	InterpolationRole LAInterpolationType
	Monumentation     *LAMonumentationType
	OriginalLocation  *geometry.GMPoint
	PID               common.Oid
	PointType         LAPointType
	ProductionMethod  *LILineage
	TransAndResult    *LATransformation

	Bfs []LABoundaryFaceString // pointBfs
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
	Isolated                     = "isolated"
	Mid                          = "mid"
	MidArc                       = "midarc"
	Start                        = "start"
)

// LAMonumentationType Monumentation type
type LAMonumentationType string

const (
	Beacon       LAMonumentationType = "beacon"
	Cornserstone                     = "cornserstone"
	Marker                           = "marker"
	NotMarked                        = "notMarked"
)

// LAPointType Point type
type LAPointType string

const (
	Control  LAPointType = "control"
	NoSource             = "noSource"
)
