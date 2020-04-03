package administrative

import "github.com/cdrlis/cdrLIS/LADM/common"

// LARRR right, restriction, responsibility
type LARRR struct {
	common.VersionedObject

	Description *string
	RID         common.Oid
	Share       *common.Fraction
	ShareCheck  *bool
	TimeSpec    *string // TODO ISO8601_ISO14825_Type
}
