package party

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LAGroupParty Group party
type LAGroupParty struct {
	LAParty

	GroupID shared.Oid
	Type    LAGroupPartyType
	Parties []LAPartyMember
}

// LAGroupPartyType Group party type
type LAGroupPartyType int

const (
	// DefaultGroupParty Default Group party type
	DefaultGroupParty LAGroupPartyType = 0
)
