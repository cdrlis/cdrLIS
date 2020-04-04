package party

import "github.com/cdrlis/cdrLIS/LADM/common"

//
// Party::LA_PartyMember
//
// An instance of class LA_PartyMember is a party member. Class LA_PartyMember is an optional association
// class between LA_Party and LA_GroupParty, see Figure 9.
type LAPartyMember struct {
	common.VersionedObject
	Share *common.Fraction
	Party *LAParty
	Group *LAGroupParty
}
