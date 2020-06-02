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

type LARequiredRelationshipSpatialUnit struct {
	common.VersionedObject

	Su1ID                   string         `gorm:"column:su1" json:"-"`
	Su1BeginLifespanVersion *time.Time     `gorm:"column:su1beginlifespanversion" json:"-"`
	Su1                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:Su1ID,Su1BeginLifespanVersion" json:"su1"`

	Su2ID                   string         `gorm:"column:su2" json:"-"`
	Su2BeginLifespanVersion *time.Time     `gorm:"column:su2beginlifespanversion" json:"-"`
	Su2                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:Su2ID,Su2BeginLifespanVersion" json:"su2"`

	Relationship ISO19125Type `gorm:"column:relationship" json:"relationship"`
}

func (LARequiredRelationshipSpatialUnit) TableName() string {
	return "LA_RequiredRelationshipSpatialUnit"
}

type ISO19125Type string

const (
	ST_Equals     ISO19125Type = "ST_Equals"
	ST_Disjoint   ISO19125Type = "ST_Disjoint"
	ST_Intersects ISO19125Type = "ST_Intersects"
	ST_Touches    ISO19125Type = "ST_Touches"
	ST_Crosses    ISO19125Type = "ST_Crosses"
	ST_Within     ISO19125Type = "ST_Within"
	ST_Contains   ISO19125Type = "ST_Contains"
	ST_Overlaps   ISO19125Type = "ST_Overlaps"
)
