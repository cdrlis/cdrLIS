package administrative

//
// Administrative::LA_Right
//
// An instance of class LA_Right is a right. LA_Right is a subclass of class LA_RRR. A right may be associated
// to zero or more [0..*] mortgages (i.e. a mortgage is associated to the affected basic administrative unit but it
// may also be specifically associated to the right which is the object of the mortgage); see Figure 10.
type LARight struct {
	LARRR

	Type LARightType

	Mortages []LAMortgage // mortageRight
}

// LARightType Right type
type LARightType int

const (
	// DefaultRight Default right type
	DefaultRight LARightType = iota
	AgriActivity
	BelowTheDepth
	BoatHarbour
	CommonwealthAcquisition
	Covenant
	Easement
	ExcludedArea
	Forest
	Freeholding
	Grazing
	HousingLand
	IndustrialState
	LandsLease
	Lease
	MainRoad
	MarinePark
	MineTenure
	NationalPark
	Occupation
	Ownership
	PortAuthority
	ProfitPrendre
	Railway
	Reserve
	StateForest
	StateLand
	TimberReserve
	TransferredProperty
	WaterResource
	WaterRights
)
