package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

type SuGroupHierarchy struct {
	common.VersionedObject

	ElementID                   string              `gorm:"column:element" json:"-"`
	ElementBeginLifespanVersion time.Time           `gorm:"column:elementbeginlifespanversion" json:"-"`
	Element                     *LASpatialUnitGroup `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ElementID,ElementBeginLifespanVersion" json:"element"`

	SetID                       string              `gorm:"column:set" json:"-"`
	SetBeginLifespanVersion     time.Time           `gorm:"column:setbeginlifespanversion" json:"-"`
	Set                         *LASpatialUnitGroup `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:SetID,SetBeginLifespanVersion" json:"set"`
}

func (SuGroupHierarchy) TableName() string {
	return "suGroupHierarchy"
}
