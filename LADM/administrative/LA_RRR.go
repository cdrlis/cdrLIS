package administrative

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LARRR right, restriction, responsibility
type LARRR struct {
	shared.VersionedObject

	Description *string
	RID         shared.Oid
	Share       *shared.Fraction
	ShareCheck  *bool
	TimeSpec    *string // TODO ISO8601_ISO14825_Type
}
