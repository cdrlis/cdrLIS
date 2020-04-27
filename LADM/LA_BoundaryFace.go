package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
)

// LABoundaryFace Boundary face string
type LABoundaryFace struct {
	common.VersionedObject

	BfID           common.Oid               `gorm:"column:bfid" json:"bfID"`
	Geometry       *geometry.GMMultiSurface `gorm:"column:geometry" json:"geometry"`
	LocationByText *string                  `gorm:"column:locationbytext" json:"locationByText"`
}
