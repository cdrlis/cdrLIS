package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"time"
)

type BAUnitAsParty struct {
	common.VersionedObject

	UnitID                   string    `gorm:"column:unit" json:"-"`
	UnitBeginLifespanVersion time.Time `gorm:"column:unitbeginlifespanversion" json:"-"`
	Unit                     *LABAUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:BAUnitID,BAUnitBeginLifespanVersion" json:"baunit"`

	PartyID                    string    `gorm:"column:party" json:"-"`
	PartyBeginLifespanVersion  time.Time `gorm:"column:partybeginlifespanversion" json:"-"`
	Party                      *LAParty  `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PartyID,PartyBeginLifespanVersion" json:"party"`
}

func (BAUnitAsParty) TableName() string {
	return "baunitAsParty"
}
