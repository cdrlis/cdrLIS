package ladm

//
// Administrative::LA_Right
//
// An instance of class LA_Right is a right. LA_Right is a subclass of class LA_RRR. A right may be associated
// to zero or more [0..*] mortgages (i.e. a mortgage is associated to the affected basic administrative unit but it
// may also be specifically associated to the right which is the object of the mortgage); see Figure 10.
type LARight struct {
	LARRR

	Type LARightType

	mortgage []LAMortgage // mortgageRight
}

func (LARight) TableName() string {
	return "LA_Right"
}

// LARightType Right type
type LARightType string

const (
	AgriActivity            LARightType = "agriActivity"
	BelowTheDepth                       = "belowTheDepth"
	BoatHarbour                         = "boatHarbour"
	CommonwealthAcquisition             = "commonwealthAcquisition"
	Covenant                            = "covenant"
	Easement                            = "easement"
	ExcludedArea                        = "excludedArea"
	//	Forest                              = "forest"	// TODO Check
	Freeholding         = "freeholding"
	Grazing             = "Grazing"
	HousingLand         = "housingLand"
	IndustrialState     = "industrialState"
	LandsLease          = "landsLease"
	Lease               = "lease"
	MainRoad            = "mainRoad"
	MarinePark          = "marinePark"
	MineTenure          = "mineTenure"
	NationalPark        = "nationalPark"
	Occupation          = "occupation"
	Ownership           = "ownership"
	PortAuthority       = "portAuthority"
	ProfitPrendre       = "profitPrendre"
	Railway             = "railway"
	Reserve             = "reserve"
	StateForest         = "stateForest"
	StateLand           = "stateLand"
	TimberReserve       = "timberReserve"
	TransferredProperty = "transferredProperty"
	WaterResource       = "waterResource"
	WaterRights         = "WaterRights"
)
