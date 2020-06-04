package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
)

//
// Administrative::LA_Restriction
//
// An instance of class LA_Restriction is a restriction. LA_Restriction is a subclass of class LA_RRR.
// LA_Mortgage is a specialization of LA_Restriction (6.4.6); see Figure 10.

type LARestriction struct {
	common.VersionedObject
	ID  string     `gorm:"column:id;primary_key" json:"-"`
	RID common.Oid `gorm:"column:rid" json:"rID"`
	RRR *LARRR     `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"rrr,omitempty"`

	PartyRequired *bool             `gorm:"column:partyrequired" json:"partyRequired"`
	Type          LARestrictionType `gorm:"column:type" json:"type"`

	Mortgage *LAMortgage `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"mortgage,omitempty"`
}

func (LARestriction) TableName() string {
	return "LA_Restriction"
}

// LARestrictionType Restriction type
type LARestrictionType string

const (
	AdminPublicServitude LARestrictionType = "adminPublicServitude"
	Monument             LARestrictionType = "monument"
	MonumentPartly       LARestrictionType = "monumentPartly"
	NoBuilding           LARestrictionType = "noBuilding"
	Servitude            LARestrictionType = "servitude"
	ServitudePartly      LARestrictionType = "servitudePartly"
)
