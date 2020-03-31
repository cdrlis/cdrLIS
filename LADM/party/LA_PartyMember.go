package party

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LAPartyMember Party member
type LAPartyMember struct {
	shared.VersionedObject

	Share *shared.Fraction
	Party *LAParty
	Group *LAGroupParty
}
