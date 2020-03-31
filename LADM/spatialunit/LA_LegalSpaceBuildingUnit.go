package spatialunit

// LALegalSpaceBuildingUnit Legal space building unit
type LALegalSpaceBuildingUnit struct {
	LASpatialUnit

	ExtPhysicalBuildingUnitID ExtPhysicalBuildingUnit
	Type                      LABuildingUnitType
}

// ExtPhysicalBuildingUnit External physical building unit
type ExtPhysicalBuildingUnit string

// LABuildingUnitType Building unit type
type LABuildingUnitType int

const (
	// DefaultBuilding Default building type
	DefaultBuilding LABuildingUnitType = 0
)
