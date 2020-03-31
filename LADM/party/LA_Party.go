package party

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LAParty LA Party
type LAParty struct {
	shared.VersionedObject

	// LA_Party
	ExtPid *shared.Oid
	Name   NullableString
	Pid    shared.Oid
	Role   *LAPartyRoleType
	Type   LAPartyType
	Groups []LAPartyMember
}

// NullableString Nullable string
type NullableString *string

// LAPartyType Party type
type LAPartyType int

const (
	// DefaultParty Default party type
	DefaultParty LAPartyType = 0
)

// LAPartyRoleType Party role type
type LAPartyRoleType int

const (
	// DefaultPartyRole Default Party role type
	DefaultPartyRole LAPartyRoleType = 0
)
