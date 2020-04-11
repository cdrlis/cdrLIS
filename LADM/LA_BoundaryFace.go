package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
)

// LABoundaryFace Boundary face string
type LABoundaryFace struct {
	common.VersionedObject

	BfID           common.Oid
	Geometry       *geometry.GMMultiSurface
	LocationByText *string
}
