package ladm

//
// Administrative::LA_Responsibility
//
// An instance of class LA_Responsibility is a responsibility. LA_Responsibility is a subclass of class LA_RRR.
// See Figure 10.

type LAResponsibility struct {
	LARRR

	Type LAResponsibilityType
}

// LAResponsibilityType Responsibility type
type LAResponsibilityType string

const (
	MonumentMaintenance LAResponsibilityType = "monumentMaintenance"
	WaterwayMaintenance                      = "waterwayMaintenance"
)
