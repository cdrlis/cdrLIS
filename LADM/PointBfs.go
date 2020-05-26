package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"time"
)

type PointBfs struct {
	common.VersionedObject

	PointID                   string                `gorm:"column:point" json:"-"`
	PointBeginLifespanVersion time.Time             `gorm:"column:pointbeginlifespanversion" json:"-"`
	Point                     *LAPoint              `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:PointID,PointBeginLifespanVersion" json:"point"`

	BfsID                     string                `gorm:"column:bfs" json:"-"`
	BfsBeginLifespanVersion   time.Time             `gorm:"column:bfsbeginlifespanversion" json:"-"`
	Bfs                       *LABoundaryFaceString `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:BfsID,BfsBeginLifespanVersion" json:"bfs"`

	Index                     int                   `gorm:"column:index" json:"index"`
}

func (PointBfs) TableName() string {
	return "pointBfs"
}
