package administrative

// LAResponsibility Responsibility
type LAResponsibility struct {
	LARRR
	Type LAResponsibilityType
}

// LAResponsibilityType Responsibility type
type LAResponsibilityType int

const (
	// DefaultResponsibility Default responsibility type
	DefaultResponsibility LAResponsibilityType = 0
)
