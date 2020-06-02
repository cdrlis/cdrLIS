package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/cdrlis/cdrLIS/ladm/common/geometry"
)

//
// Surveying and Representation::LA_BoundaryFaceString
//
// boundary face string: boundary (4.1.3) forming part of the outside of a spatial unit (4.1.23).
// NOTE Boundary face strings are used to represent the boundaries of spatial units by means of line strings in 2D.
// This 2D representation is a 2D boundary in a 2D land administration (4.1.10) system. In a 3D land administration system
// it represents a series of vertical boundary faces (4.1.4) where an unbounded volume is assumed, surrounded by
// boundary faces which intersect the Earth’s surface (such as traditionally depicted in the cadastral map).
// An instance of class LA_BoundaryFaceString is a boundary face string. LA_BoundaryFaceString is associated
// to class LA_Point and class LA_SpatialSource to document the origin of the geometry. In the case of a
// location by text, a boundary face string would not be defined by points. However, in all other cases, a
// boundary face string shall be defined by two or more [2..*] points (i.e. as a minimum a boundary starts and
// ends at a point, i.e. a straight line).

type LABoundaryFaceString struct {
	common.VersionedObject
	ID string `gorm:"column:id;primary_key" json:"-"`

	BfsID          common.Oid             `gorm:"column:bfsid" json:"bfsID"`
	Geometry       *geometry.GMMultiCurve `gorm:"column:geometry" json:"geometry"`
	LocationByText *string                `gorm:"column:locationbytext" json:"locationByText"`

	point   []LAPoint             // pointBfs
	MinusSu []BfsSpatialUnitMinus `gorm:"foreignkey:BfsID,BfsBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion;" json:"suMinus"`
	PlusSu  []BfsSpatialUnitPlus  `gorm:"foreignkey:BfsID,BfsBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion;" json:"suPlus"`
}

func (LABoundaryFaceString) TableName() string {
	return "LA_BoundaryFaceString"
}
