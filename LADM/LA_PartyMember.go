package ladm

import "github.com/cdrlis/cdrLIS/LADM/common"

//
// Party::LA_PartyMember
//
// An instance of class LA_PartyMember is a party member. Class LA_PartyMember is an optional association
// class between LA_Party and LA_GroupParty, see Figure 9.
type LAPartyMember struct {
	common.VersionedObject
	Share *common.Fraction	`gorm:"column:fraction" json:"fraction"`
	Party *LAParty			`gorm:"column:parties" json:"party"`
	Group *LAGroupParty		`gorm:"column:groups" json:"group"`
}

func (LAPartyMember) TableName() string {
	return "LA_PartyMember"
}