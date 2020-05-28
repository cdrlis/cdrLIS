package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"time"
)

//
// Party::LA_PartyMember
//
// An instance of class LA_PartyMember is a party member. Class LA_PartyMember is an optional association
// class between LA_Party and LA_GroupParty, see Figure 9.
type LAPartyMember struct {
	common.VersionedObject

	Share *common.Fraction `gorm:"column:fraction" json:"fraction"`

	PartyID                   string    `gorm:"column:parties;primary_key" json:"-"`
	PartyBeginLifespanVersion time.Time `gorm:"column:partiesbeginlifespanversion;primary_key" json:"-"`
	Party                     *LAParty  `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PartyID,PartyBeginLifespanVersion" json:"party"`

	GroupID                   string        `gorm:"column:groups;primary_key" json:"-"`
	GroupBeginLifespanVersion time.Time     `gorm:"column:groupsbeginlifespanversion;primary_key" json:"-"`
	Group                     *LAGroupParty `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:GroupID,GroupBeginLifespanVersion" json:"group"`
}

func (LAPartyMember) TableName() string {
	return "LA_PartyMember"
}
