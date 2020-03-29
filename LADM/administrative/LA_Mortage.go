package administrative

// LAMortage Mortage
type LAMortage struct {
	LARestriction

	Amount       *Currency
	InterestRate *float64
	Ranking      *int
	Type         LAMortageType
}

// Currency Currency
type Currency string

// LAMortageType Mortage type
type LAMortageType int

const (
	// DefaultMortage Default mortage type
	DefaultMortage LAMortageType = 0
)
