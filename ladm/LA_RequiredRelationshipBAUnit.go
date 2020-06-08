package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

//
// Spatial Unit::LA_RequiredRelationshipSpatialUnit
//
// An instance of association class LA_RequiredRelationshipSpatialUnit is a required relationship between
// spatial units, see Figures 11 and 12. A required relationship between spatial units can be associated to zero
// or more [0..*] spatial sources to provide supporting documentation for the explicit relationhip.

type LARequiredRelationshipBAUnit struct {
	common.VersionedObject

	Unit1ID                   string         `gorm:"column:unit1" json:"-"`
	Unit1BeginLifespanVersion time.Time      `gorm:"column:unit1beginlifespanversion" json:"-"`
	Unit1                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:Unit1ID,Unit1BeginLifespanVersion" json:"unit1"`

	Unit2ID                   string         `gorm:"column:unit2" json:"-"`
	Unit2BeginLifespanVersion time.Time      `gorm:"column:unit2beginlifespanversion" json:"-"`
	Unit2                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:Unit2ID,Unit2BeginLifespanVersion" json:"unit2"`

	Relationship              string         `gorm:"column:relationship" json:"relationship"`
}

func (LARequiredRelationshipBAUnit) TableName() string {
	return "LA_RequiredRelationshipBAUnit"
}
