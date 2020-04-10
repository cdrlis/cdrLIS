package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/external"

// LALegalSpaceBuildingUnit Legal space building unit
type LALegalSpaceBuildingUnit struct {
	LASpatialUnit
	ExtPhysicalBuildingUnitID *external.ExtPhysicalBuildingUnit
	Type                      *LABuildingUnitType
}

//
// LA_BuildingUnitType: the LA_BuildingUnitType code list includes all the various building unit types, such
// as private or commercial, used in a specific land administration profile implementation. The
// LA_BuildingUnitType code list is required only if the attribute type in LA_LegalSpaceBuildingUnit class is
// implemented. The code list shall provide a complete list of all codes with a name and description.

type LABuildingUnitType string

const (
	Individual LABuildingUnitType = "individual"
	Shared                        = "shared"
)
