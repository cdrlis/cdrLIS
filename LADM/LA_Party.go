package ladm

import (
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
	ID     string         `gorm:"column:id;primary_key"`
	ExtPid *common.Oid    `gorm:"column:extpid"`
	Name   *string        `gorm:"column:name"`
	Pid    common.Oid     `gorm:"column:pid"`
	Role   pq.StringArray `gorm:"type:varchar(100)[]"`
	Type   string         `gorm:"column:type"`
	Groups []LAPartyMember

	Unit []LABAunit // baunitAsParty
	RRR  []LARRR    // rrrParty
}

func (LAParty) TableName() string {
	return "LA_Party"
}

// LAPartyType Party type
type LAPartyType string // TODO Temporary not used (GORM debuging)

const (
	BAUnit           LAPartyType = "baunit"
	Group                        = "group"
	NaturalPerson                = "naturalPerson"
	NonNaturalPerson             = "nonNaturalPerson"
)

// LAPartyRoleType Party role type
type LAPartyRoleType string // TODO Temporary not used (GORM debuging)

const (
	Bank               LAPartyRoleType = "bank"
	CertifiedSurveyor                  = "certifiedSurveyor"
	Citizen                            = "citizen"
	Conveyancer                        = "conveyancer"
	Employee                           = "employee"
	Farmer                             = "farmer"
	MoneyProvider                      = "moneyProvider"
	Notary                             = "notary"
	StateAdministrator                 = "stateAdministrator"
	Surveyor                           = "surveyor"
	Writer                             = "writer"
)
