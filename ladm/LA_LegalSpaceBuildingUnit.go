package ladm

import (
	"github.com/cdrlis/cdrLIS/ladm/common"
	"github.com/cdrlis/cdrLIS/ladm/external"
)

// LALegalSpaceBuildingUnit Legal space building unit
type LALegalSpaceBuildingUnit struct {
	common.VersionedObject
	ID          string         `gorm:"column:id;primary_key" json:"-"`
	SpatialUnit *LASpatialUnit `gorm:"foreignkey:ID,BeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"spatialUnit,omitempty"`

	ExtPhysicalBuildingUnitID *external.ExtPhysicalBuildingUnit `gorm:"column:extaddressid" json:"extAddressID"`
	Type                      *LABuildingUnitType               `gorm:"column:type" json:"type"`
}

func (LALegalSpaceBuildingUnit) TableName() string {
	return "LA_LegalSpaceBuildingUnit"
}

//
// LA_BuildingUnitType: the LA_BuildingUnitType code list includes all the various building unit types, such
// as private or commercial, used in a specific land administration profile implementation. The
// LA_BuildingUnitType code list is required only if the attribute type in LA_LegalSpaceBuildingUnit class is
// implemented. The code list shall provide a complete list of all codes with a name and description.

type LABuildingUnitType string

const (
	Individual LABuildingUnitType = "individual"
	Shared     LABuildingUnitType = "shared"
)
