package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
)

//
// Administrative::LA_BAUnit
//
// An instance of class LA_BAUnit is a basic administrative unit. LA_BAUnit is associated to class LA_Party (a
// party may be a basic administrative unit). A basic administrative unit is associated to zero or more [0..*] spatial
// units. A basic administrative unit shall be associated to one or more [1..*] instances of right, restriction or
// responsibility (i.e. a basic administrative unit cannot exist if there is not at least one right, restriction or
// responsibility associated to it). A basic administrative unit can be spatially related, through a required
// relationship, to zero or more [0..*] other basic administrative units (i.e. create an explicit spatial relationship
// between two basic administrative units when the geometry is missing or inaccurate to provide reliable implicit
// results). Basic administrative units do not need to be related explicitly. However, if an explicit required
// relationship is specified, a basic administrative unit shall be associated to one or more [1..*] other basic
// administrative units. A basic administrative unit can be associated to zero or more [0..*] administrative sources
// (i.e. the basic administrative unit is usually described as the object affected by the right, restriction or
// responsibility in the administrative source). A basic administrative unit can be associated to zero or more [0..*]
// spatial sources (i.e. the extent – part of – of a basic administrative unit can be described on a spatial source).
// See Figure 10.

type LABAUnit struct {
	common.VersionedObject
	ID   string       `gorm:"column:id;primary_key" json:"-"`
	Name *string      `gorm:"-" json:"-"`// TODO: Add column to db
	Type LABAUnitType `gorm:"column:type" json:"type"`
	UID  common.Oid   `gorm:"column:uid" json:"uID"`

	Party []LAParty // baunitAsParty

	Rights           []LARight          `gorm:"foreignkey:UnitID,UnitBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"rights"`
	Responsibilities []LAResponsibility `gorm:"foreignkey:UnitID,UnitBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"responsibilities"`
	Restrictions     []LARestriction    `gorm:"foreignkey:UnitID,UnitBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"restrictions"`

	SU []SuBAUnit `gorm:"foreignkey:BaUnitID,BaUnitBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion;" json:"su"`
}

func (LABAUnit) TableName() string {
	return "LA_BAUnit"
}

// LABAUnitType BA unit type
type LABAUnitType string

const (
	BasicPropertyUnit LABAUnitType = "basicPropertyUnit"
	LeasedUnit        LABAUnitType = "leasedUnit"
	RightOfUseUnit    LABAUnitType = "rightOfUseUnit"
)
