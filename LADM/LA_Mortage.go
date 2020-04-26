package ladm

//
// Administrative::LA_Mortgage
//
// An instance of class LA_Mortgage is a mortgage. LA_Mortgage is a subclass of LA_Restriction. LA_Mortgage
// is associated to class LA_Right (the right that is the basis for the mortgage). A mortgage can be associated to
// zero or more [0..*] rights (i.e. a mortgage can be associated specifically to the right which is the object of the
// mortgage). In all cases, the mortgage is associated, through LA_Restriction and LA_RRR, to the basic
//administrative unit which is affected by the mortgage; see Figure 10.

type LAMortgage struct {
	LARestriction

	Amount       *Currency
	InterestRate *float32
	Ranking      *int
	Type         *LAMortgageType

	Rights []LARight // mortageRight
}

func (LAMortgage) TableName() string {
	return "LA_Mortgage"
}

type Currency struct {
	Amount float32
	Code   iso4217
}

// Currency based on ISO 4217
type iso4217 string

const (
	AED iso4217 = "AED"
	// ...
	EUR = "EUR"
	// ...
	USD = "USD"
	// ...
	ZWL = "ZWL"
)

// LAMortageType Mortage type
type LAMortgageType string

const (
	LevelPayment LAMortgageType = "levelPayment"
	Linear                      = "linear"
	Microcredit                 = "microcredit"
)
