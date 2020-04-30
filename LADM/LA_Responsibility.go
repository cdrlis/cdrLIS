package ladm

import "time"

//
// Administrative::LA_Responsibility
//
// An instance of class LA_Responsibility is a responsibility. LA_Responsibility is a subclass of class LA_RRR.
// See Figure 10.

type LAResponsibility struct {
	LARRR

	Type LAResponsibilityType `gorm:"column:type" json:"type"`

	PartyID                   string           `gorm:"column:party" json:"-"`
	PartyBeginLifespanVersion time.Time        `gorm:"column:partybeginlifespanversion" json:"-"`
	Party                     *LAParty         `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PartyID,PartyBeginLifespanVersion" json:"party"`

	UnitID                   string           `gorm:"column:baunit" json:"-"`
	UnitBeginLifespanVersion time.Time        `gorm:"column:baunitbeginlifespanversion" json:"-"`
	Unit                     *LABAUnit         `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:UnitID,UnitBeginLifespanVersion" json:"unit"`
}

func (LAResponsibility) TableName() string {
	return "LA_Responsibility"
}

// LAResponsibilityType Responsibility type
type LAResponsibilityType string

const (
	MonumentMaintenance LAResponsibilityType = "monumentMaintenance"
	WaterwayMaintenance LAResponsibilityType = "waterwayMaintenance"
)
