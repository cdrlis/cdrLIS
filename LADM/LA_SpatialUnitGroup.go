package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
)

//
// Spatial Unit::LA_SpatialUnitGroup
//
// Any number of spatial units (4.1.23), considered as an entity.
// An instance of class LA_SpatialUnitGroup is a spatial unit group. A spatial unit group is made of one or more
// [1..*] parts/elements (which can be spatial units, or spatial unit groups, or a combination of spatial units and
// spatial unit groups). A spatial unit group is part of zero or one [0..1] larger spatial unit group, which again can
// even be part of zero or one [0..1] larger spatial unit group, and so on. See Figure 11.

type LASpatialUnitGroup struct {
	common.VersionedObject

	ID             string            `gorm:"column:id;primary_key" json:"-"`
	HierarchyLevel int               `gorm:"column:hierarchylevel" json:"hierarchyLevel"`
	Label          *string           `gorm:"column:label" json:"label"`
	Name           *string           `gorm:"column:name" json:"name"`
	ReferencePoint *geometry.GMPoint `gorm:"column:referencepoint" json:"referencePoint"`
	SugID          common.Oid        `gorm:"column:sugid" json:"sugID"`

	SuGroupHierarchy *SuGroupHierarchy `gorm:"foreignkey:ElementID,ElementBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"hierarchy,omitempty"`

	SpatialUnits     []SuSuGroup       `gorm:"foreignkey:WholeID,WholeBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"units,omitempty"`
}

func (LASpatialUnitGroup) TableName() string {
	return "LA_SpatialUnitGroup"
}
