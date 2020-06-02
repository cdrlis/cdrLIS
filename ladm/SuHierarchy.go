package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

type SuHierarchy struct {
	common.VersionedObject

	ChildID                    string         `gorm:"column:child" json:"-"`
	ChildBeginLifespanVersion  time.Time      `gorm:"column:childbeginlifespanversion" json:"-"`
	Child                      *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ChildID,ChildBeginLifespanVersion" json:"child"`

	ParentID                   string         `gorm:"column:parent" json:"-"`
	ParentBeginLifespanVersion time.Time      `gorm:"column:parentbeginlifespanversion" json:"-"`
	Parent                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ParentID,ParentBeginLifespanVersion" json:"parent"`
}

func (SuHierarchy) TableName() string {
	return "suHierarchy"
}
