package administrative

import "github.com/cdrlis/cdrLIS/LADM/common"

// LABaunit Basic administrative unit
type LABaunit struct {
	common.VersionedObject

	Name *string
	Type LABAUnitType
	UID  common.Oid
}

// LABAUnitType BA unit type
type LABAUnitType int

const (
	// DefaultBAUnit Default BA unit type
	DefaultBAUnit LABAUnitType = 0
)
