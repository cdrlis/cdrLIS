package ladm

import "github.com/cdrlis/cdrLIS/ladm/common"

//
// Party::LA_GroupParty
//
// An instance of class LA_GroupParty is a group party. Class LA_GroupParty is a subclass of LA_Party, thus
// allowing instances of class LA_GroupParty to have an association with instances of class LA_RRR (and
// thereby also to class LA_BAUnit). A group party consists of two or more [2..*] parties, but also of other group
// parties (that is to say, a group party of group parties). Conversely, a party is a member of zero or more [0..*]
// group parties, see Figure 9.

type LAGroupParty struct {
	common.VersionedObject
	ID    string   `gorm:"column:id;primary_key" json:"-"`
	Party *LAParty `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"party,omitempty"`

	PID  common.Oid       `gorm:"column:pid" json:"pID"`
	Type LAGroupPartyType `gorm:"column:type" json:"type"`

	Parties []LAPartyMember `gorm:"foreignkey:PID,GroupBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"parties"`
}

func (LAGroupParty) TableName() string {
	return "LA_GroupParty"
}

// LAGroupPartyType Group party type
type LAGroupPartyType string

const (
	Association LAGroupPartyType = "association"
	BAUnitGroup LAGroupPartyType = "baunitGroup"
	Family      LAGroupPartyType = "family"
	Tribe       LAGroupPartyType = "tribe"
)
