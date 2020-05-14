package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"time"
)


type BfsSpatialUnitMinus struct {
	common.VersionedObject

	SuID                   string         `gorm:"column:su" json:"-"`
	SuBeginLifespanVersion time.Time      `gorm:"column:subeginlifespanversion" json:"-"`
	Su                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:SuID,SuBeginLifespanVersion" json:"su"`

	BfsID                   string                `gorm:"column:baunit" json:"-"`
	BfsBeginLifespanVersion time.Time             `gorm:"column:baunitbeginlifespanversion" json:"-"`
	Bfs                  *LABoundaryFaceString `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:BfsID,BfsBeginLifespanVersion" json:"bfs"`
}

func (BfsSpatialUnitMinus) TableName() string {
	return "bfsSpatialUnitMinus"
}
