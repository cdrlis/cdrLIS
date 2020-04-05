package external

import "github.com/cdrlis/cdrLIS/LADM/common"

//
// Class ExtPhysicalBuildingUnit is a class for the external registration of mapping data of building units.
// ExtPhysicalBuildingUnit is associated to class LA_LegalSpaceBuildingUnit, see Figure K.1.
//
type ExtPhysicalBuildingUnit struct {
	common.VersionedObject
	ExtAddressID *ExtAddress
}
