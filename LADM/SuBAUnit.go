package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"time"
)

//
// Administrative::LA_Right
//
// An instance of class LA_Right is a right. LA_Right is a subclass of class LA_RRR. A right may be associated
// to zero or more [0..*] mortgages (i.e. a mortgage is associated to the affected basic administrative unit but it
// may also be specifically associated to the right which is the object of the mortgage); see Figure 10.
type SuBAUnit struct {
	common.VersionedObject

	SUID                   string         `gorm:"column:su" json:"-"`
	SUBeginLifespanVersion time.Time      `gorm:"column:subeginlifespanversion" json:"-"`
	SU                     *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:SUID,SUBeginLifespanVersion" json:"su"`

	BaUnitID                   string    `gorm:"column:baunit" json:"-"`
	BaUnitBeginLifespanVersion time.Time `gorm:"column:baunitbeginlifespanversion" json:"-"`
	BaUnit                     *LABAUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:BaUnitID,BaUnitBeginLifespanVersion" json:"baunit"`
}

func (SuBAUnit) TableName() string {
	return "suBaunit"
}
