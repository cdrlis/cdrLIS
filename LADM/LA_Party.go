package ladm

import (
	"database/sql/driver"
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/lib/pq"
)

//
// Party::LA_Party
//
// An instance of class LA_Party is a party. A party is associated to zero or more [0..*] instances of a subclass of
// LA_RRR. LA_Party is also associated to LA_BAUnit, to cater for the fact that a basic administrative unit can
// be a party (e.g. a basic administrative unit holding an easement on another basic administrative unit). A party
// may be associated to zero or more [0..*] administrative sources (i.e. the author of a transfer document is
// defined as a party playing the role of conveyancer in a source). A party may be associated to zero or more
// [0..*] spatial sources (i.e. the author of a survey document is defined as a party playing the role of surveyor in
// a source); see Figure 9.
type LAParty struct {
	common.VersionedObject
	ID     string               `gorm:"column:id;primary_key" json:"-"`
	ExtPid *common.Oid          `gorm:"column:extpid" json:"extPID"`
	Name   *string              `gorm:"column:name" json:"name"`
	PID    common.Oid           `gorm:"column:pid;type:Oid" json:"pID"`
	Role   LAPartyRoleTypeArray `gorm:"column:role" json:"role"`
	Type   LAPartyType          `gorm:"column:type" json:"type"`

	Groups []LAPartyMember `gorm:"foreignkey:PartyID,PartyBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"groups"`

	Unit []LABAUnit // baunitAsParty

	// rrrParty
	Rights           []LARight          `gorm:"foreignkey:PartyID,PartyBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"rights"`
	Responsibilities []LAResponsibility `gorm:"foreignkey:PartyID,PartyBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"responsibilities"`
	Restrictions     []LARestriction    `gorm:"foreignkey:PartyID,PartyBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"restrictions"`
}

func (LAParty) TableName() string {
	return "LA_Party"
}

// LAPartyType Party type
type LAPartyType string

const (
	BAUnit           LAPartyType = "baunit"
	Group            LAPartyType = "group"
	NaturalPerson    LAPartyType = "naturalPerson"
	NonNaturalPerson LAPartyType = "nonNaturalPerson"
)

type LAPartyRoleTypeArray pq.StringArray

func (a *LAPartyRoleTypeArray) Scan(src interface{}) error {
	return (*pq.StringArray)(a).Scan(src)
}

func (a LAPartyRoleTypeArray) Value() (driver.Value, error) {
	return (pq.StringArray)(a).Value()
}

// LAPartyRoleType Party role type
type LAPartyRoleType string

const (
	Bank               LAPartyRoleType = "bank"
	CertifiedSurveyor  LAPartyRoleType = "certifiedSurveyor"
	Citizen            LAPartyRoleType = "citizen"
	Conveyancer        LAPartyRoleType = "conveyancer"
	Employee           LAPartyRoleType = "employee"
	Farmer             LAPartyRoleType = "farmer"
	MoneyProvider      LAPartyRoleType = "moneyProvider"
	Notary             LAPartyRoleType = "notary"
	StateAdministrator LAPartyRoleType = "stateAdministrator"
	Surveyor           LAPartyRoleType = "surveyor"
	Writer             LAPartyRoleType = "writer"
)
