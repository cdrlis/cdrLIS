package surveying

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
)

// LAPoint Point
type LAPoint struct {
	common.VersionedObject

	InterpolationRole LAInterpolationType
	Monumentation     *LAMonumentationType
	OriginalLocation  *geometry.GMPoint
	PID               common.Oid
	PointType         LAPointType
	ProductionMethod  *LILineage
	TransAndResult    *LATransformation
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
