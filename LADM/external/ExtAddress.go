package external

import (
	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
)

//
// Class ExtAddress is a class for an external registration of addresses (an address being a direction for finding
// a location), see Figure K.1.
//

type ExtAddress struct {
	common.VersionedObject
	AddressAreaName   *string
	AddressCoordinate *geometry.GMPoint
	AddressID         common.Oid
	BuildingName      *string
	BuildingNumber    *string
	City              *string
	Country           *string
	PostalCode        *string
	PostBox           *string
	State             *string
	StreetName        *string
}
