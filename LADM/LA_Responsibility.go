package ladm

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
)

//
// Administrative::LA_Responsibility
//
// An instance of class LA_Responsibility is a responsibility. LA_Responsibility is a subclass of class LA_RRR.
// See Figure 10.

type LAResponsibility struct {
	common.VersionedObject
	ID  string `gorm:"column:id;primary_key" json:"-"`
	RRR *LARRR `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"rrr,omitempty"`

	Type LAResponsibilityType `gorm:"column:type" json:"type"`
}

func (LAResponsibility) TableName() string {
	return "LA_Responsibility"
}

// LAResponsibilityType Responsibility type
type LAResponsibilityType string

const (
	MonumentMaintenance LAResponsibilityType = "monumentMaintenance"
	WaterwayMaintenance LAResponsibilityType = "waterwayMaintenance"
)
