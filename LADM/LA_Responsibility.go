package ladm

//
// Administrative::LA_Responsibility
//
// An instance of class LA_Responsibility is a responsibility. LA_Responsibility is a subclass of class LA_RRR.
// See Figure 10.

type LAResponsibility struct {
	LARRR

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
