package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

type BAUnitAsParty struct {
	common.VersionedObject

	UnitID                   string    `gorm:"column:unit;primary_key" json:"-"`
	UnitBeginLifespanVersion time.Time `gorm:"column:unitbeginlifespanversion" json:"-"`
	Unit                     *LABAUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:BAUnitID,BAUnitBeginLifespanVersion" json:"baunit"`

	PartyID                    string    `gorm:"column:party;primary_key" json:"-"`
	PartyBeginLifespanVersion  time.Time `gorm:"column:partybeginlifespanversion" json:"-"`
	Party                      *LAParty  `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PartyID,PartyBeginLifespanVersion" json:"party"`
}

func (BAUnitAsParty) TableName() string {
	return "baunitAsParty"
}
