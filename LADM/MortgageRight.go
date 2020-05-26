package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"time"
)

type MortgageRight struct {
	common.VersionedObject

	MortgageID                   string         `gorm:"column:mortgage" json:"-"`
	MortgageBeginLifespanVersion time.Time      `gorm:"column:mortgagebeginlifespanversion" json:"-"`
	Mortgage                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:MortgageID,MortgageBeginLifespanVersion" json:"mortgage"`

	RightID                      string         `gorm:"column:right" json:"-"`
	RightBeginLifespanVersion    time.Time      `gorm:"column:rightbeginlifespanversion" json:"-"`
	Right                        *LARight       `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:RightID,RightBeginLifespanVersion" json:"right"`
	
	Index                        int            `gorm:"column:index" json:"index"`
}

func (MortgageRight) TableName() string {
	return "mortgageRight"
}
