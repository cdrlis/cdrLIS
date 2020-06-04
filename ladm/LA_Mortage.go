package ladm

import "github.com/cdrlis/cdrLIS/ladm/common"

//
// Administrative::LA_Mortgage
//
// An instance of class LA_Mortgage is a mortgage. LA_Mortgage is a subclass of LA_Restriction. LA_Mortgage
// is associated to class LA_Right (the right that is the basis for the mortgage). A mortgage can be associated to
// zero or more [0..*] rights (i.e. a mortgage can be associated specifically to the right which is the object of the
// mortgage). In all cases, the mortgage is associated, through LA_Restriction and LA_RRR, to the basic
//administrative unit which is affected by the mortgage; see Figure 10.

type LAMortgage struct {
	common.VersionedObject
	ID          string         `gorm:"column:id;primary_key" json:"-"`
	RID         common.Oid     `gorm:"column:rid" json:"rID"`
	Restriction *LARestriction `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"restriction,omitempty"`

	Amount       *Currency       `gorm:"column:amount" json:"amount"`
	InterestRate *float32        `gorm:"column:interestRate" json:"interestrate"`
	Ranking      *int            `gorm:"column:ranking" json:"ranking"`
	Type         *LAMortgageType `gorm:"column:type" json:"type"`

	Rights []MortgageRight `gorm:"foreignkey:MortgageID,MortgageBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion;" json:"rights"`
}

func (LAMortgage) TableName() string {
	return "LA_Mortgage"
}

type Currency struct {
	Amount float32
	Code   ISO4217Type
}

// Currency based on ISO 4217
type ISO4217Type string

const (
	AED ISO4217Type = "AED"
	// ...
	EUR ISO4217Type = "EUR"
	// ...
	USD ISO4217Type = "USD"
	// ...
	ZWL ISO4217Type = "ZWL"
)

// LAMortageType Mortage type
type LAMortgageType string

const (
	LevelPayment LAMortgageType = "levelPayment"
	Linear       LAMortgageType = "linear"
	Microcredit  LAMortgageType = "microcredit"
)
