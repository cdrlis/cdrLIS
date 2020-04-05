package party

import (
	"github.com/cdrlis/cdrLIS/LADM/administrative"
	"github.com/cdrlis/cdrLIS/LADM/common"
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
	// LA_Party
	ExtPid *common.Oid
	Name   *string
	Pid    common.Oid
	Role   *LAPartyRoleType
	Type   LAPartyType
	Groups []LAPartyMember

	Unit []administrative.LABAunit // baunitAsParty
	RRR  []administrative.LARRR    // rrrParty
}

// LAPartyType Party type
type LAPartyType int

const (
	DefaultPartyType LAPartyType = iota
	BAUnit
	Group
	NaturalPerson
	NonNaturalPerson
)

// LAPartyRoleType Party role type
type LAPartyRoleType int

const (
	// DefaultPartyRole Default Party role type
	DefaultPartyRole LAPartyRoleType = iota
	Bank
	CertifiedSurveyor
	Citizen
	Conveyancer
	Employee
	Farmer
	MoneyProvider
	Notary
	StateAdministrator
	Surveyor
	Writer
)
