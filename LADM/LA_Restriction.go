package ladm

import "time"

//
// Administrative::LA_Restriction
//
// An instance of class LA_Restriction is a restriction. LA_Restriction is a subclass of class LA_RRR.
// LA_Mortgage is a specialization of LA_Restriction (6.4.6); see Figure 10.

type LARestriction struct {
	LARRR

	PartyRequired *bool             `gorm:"column:partyrequired" json:"partyRequired"`
	Type          LARestrictionType `gorm:"column:type" json:"type"`

	PartyID                   string           `gorm:"column:party" json:"-"`
	PartyBeginLifespanVersion time.Time        `gorm:"column:partybeginlifespanversion" json:"-"`
	Party                     *LAParty         `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PartyID,PartyBeginLifespanVersion" json:"party"`

	UnitID                   string           `gorm:"column:baunit" json:"-"`
	UnitBeginLifespanVersion time.Time        `gorm:"column:baunitbeginlifespanversion" json:"-"`
	Unit                     *LABAUnit         `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:UnitID,UnitBeginLifespanVersion" json:"unit"`
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
