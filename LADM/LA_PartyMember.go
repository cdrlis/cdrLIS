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
	PartyID                   string           `gorm:"column:parties" json:"-"`
	PartyBeginLifespanVersion time.Time        `gorm:"column:partiesbeginlifespanversion" json:"-"`
	GroupID                   string           `gorm:"column:groups" json:"-"`
	GroupBeginLifespanVersion time.Time        `gorm:"column:groupsbeginlifespanversion" json:"-"`
	Share                     *common.Fraction `gorm:"column:fraction" json:"fraction"`
	Party                     *LAParty         `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PartyID,PartyBeginLifespanVersion;PRELOAD:false" json:"party"`
	Group                     *LAGroupParty    `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:GroupID,GroupBeginLifespanVersion;PRELOAD:false" json:"group"`
}

func (LAPartyMember) TableName() string {
	return "LA_PartyMember"
}
