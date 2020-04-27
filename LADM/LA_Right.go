package ladm

//
// Administrative::LA_Right
//
// An instance of class LA_Right is a right. LA_Right is a subclass of class LA_RRR. A right may be associated
// to zero or more [0..*] mortgages (i.e. a mortgage is associated to the affected basic administrative unit but it
// may also be specifically associated to the right which is the object of the mortgage); see Figure 10.
type LARight struct {
	LARRR

	Type LARightType `gorm:"column:type" json:"type"`

	mortgage []LAMortgage // mortgageRight
}

func (LARight) TableName() string {
	return "LA_Right"
}

// LARightType Right type
type LARightType string

const (
	AgriActivity            LARightType = "agriActivity"
	BelowTheDepth           LARightType = "belowTheDepth"
	BoatHarbour             LARightType = "boatHarbour"
	CommonwealthAcquisition LARightType = "commonwealthAcquisition"
	Covenant                LARightType = "covenant"
	Easement                LARightType = "easement"
	ExcludedArea            LARightType = "excludedArea"
	Forest1                 LARightType = "forest" // TODO Check
	Freeholding             LARightType = "freeholding"
	Grazing                 LARightType = "Grazing"
	HousingLand             LARightType = "housingLand"
	IndustrialState         LARightType = "industrialState"
	LandsLease              LARightType = "landsLease"
	Lease                   LARightType = "lease"
	MainRoad                LARightType = "mainRoad"
	MarinePark              LARightType = "marinePark"
	MineTenure              LARightType = "mineTenure"
	NationalPark            LARightType = "nationalPark"
	Occupation              LARightType = "occupation"
	Ownership               LARightType = "ownership"
	PortAuthority           LARightType = "portAuthority"
	ProfitPrendre           LARightType = "profitPrendre"
	Railway                 LARightType = "railway"
	Reserve                 LARightType = "reserve"
	StateForest             LARightType = "stateForest"
	StateLand               LARightType = "stateLand"
	TimberReserve           LARightType = "timberReserve"
	TransferredProperty     LARightType = "transferredProperty"
	WaterResource           LARightType = "waterResource"
	WaterRights             LARightType = "WaterRights"
)
