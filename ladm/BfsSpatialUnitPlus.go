package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

type BfsSpatialUnitPlus struct {
	common.VersionedObject

	SuID                    string                `gorm:"column:su" json:"-"`
	SuBeginLifespanVersion  time.Time             `gorm:"column:subeginlifespanversion" json:"-"`
	Su                      *LASpatialUnit        `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:SuID,SuBeginLifespanVersion" json:"su"`

	BfsID                   string                `gorm:"column:bfs" json:"-"`
	BfsBeginLifespanVersion time.Time             `gorm:"column:bfsbeginlifespanversion" json:"-"`
	Bfs                     *LABoundaryFaceString `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:BfsID,BfsBeginLifespanVersion" json:"bfs"`
}

func (BfsSpatialUnitPlus) TableName() string {
	return "plus"
}
