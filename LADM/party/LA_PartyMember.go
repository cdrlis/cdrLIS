package party

import "github.com/cdrlis/cdrLIS/LADM/common"

// LAPartyMember Party member
type LAPartyMember struct {
	common.VersionedObject

	Share *common.Fraction
	Party *LAParty
	Group *LAGroupParty
}
