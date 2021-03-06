package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

type SuSuGroup struct {
	common.VersionedObject

	PartID                    string         `gorm:"column:part;primary_key" json:"-"`
	PartBeginLifespanVersion  time.Time      `gorm:"column:partbeginlifespanversion" json:"-"`
	Part                      *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PartID,PartBeginLifespanVersion" json:"part"`

	WholeID                   string         `gorm:"column:whole;primary_key" json:"-"`
	WholeBeginLifespanVersion time.Time      `gorm:"column:wholebeginlifespanversion" json:"-"`
	Whole                     *LASpatialUnitGroup `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:WholeID,WholeBeginLifespanVersion" json:"whole"`
}

func (SuSuGroup) TableName() string {
	return "suSuGroup"
}
