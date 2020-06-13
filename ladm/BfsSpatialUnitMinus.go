package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"time"
)

type BfsSpatialUnitMinus struct {
	common.VersionedObject

	SuID                    string                `gorm:"column:su;primary_key" json:"-"`
	SuBeginLifespanVersion  time.Time             `gorm:"column:subeginlifespanversion" json:"-"`
	Su                      *LASpatialUnit        `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:SuID,SuBeginLifespanVersion" json:"su"`

	BfsID                   string                `gorm:"column:bfs;primary_key" json:"-"`
	BfsBeginLifespanVersion time.Time             `gorm:"column:bfsbeginlifespanversion" json:"-"`
	Bfs                     *LABoundaryFaceString `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:BfsID,BfsBeginLifespanVersion" json:"bfs"`
}

func (BfsSpatialUnitMinus) TableName() string {
	return "minus"
}
