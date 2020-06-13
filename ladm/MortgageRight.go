package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

type MortgageRight struct {
	common.VersionedObject

	MortgageID                   string         `gorm:"column:mortgage;primary_key" json:"-"`
	MortgageBeginLifespanVersion time.Time      `gorm:"column:mortgagebeginlifespanversion" json:"-"`
	Mortgage                     *LAMortgage    `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:MortgageID,MortgageBeginLifespanVersion" json:"mortgage"`

	RightID                      string         `gorm:"column:right_;primary_key" json:"-"`
	RightBeginLifespanVersion    time.Time      `gorm:"column:right_beginlifespanversion" json:"-"`
	Right                        *LARight       `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:RightID,RightBeginLifespanVersion" json:"right"`
	
	Index                        int            `gorm:"column:index" json:"index"`
}

func (MortgageRight) TableName() string {
	return "mortgageRight"
}
