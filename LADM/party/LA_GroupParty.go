package party

import "github.com/cdrlis/cdrLIS/LADM/common"

// LAGroupParty Group party
type LAGroupParty struct {
	LAParty

	GroupID common.Oid
	Type    LAGroupPartyType
	Parties []LAPartyMember
}

// LAGroupPartyType Group party type
type LAGroupPartyType int

const (
	// DefaultGroupParty Default Group party type
	DefaultGroupParty LAGroupPartyType = 0
)
