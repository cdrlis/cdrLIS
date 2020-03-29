package administrative

// LARestriction Restriction
type LARestriction struct {
	LARRR
	PartyRequired *bool
	Type          LARestrictionType
}

// LARestrictionType Restriction type
type LARestrictionType int

const (
	// DefaultRestriction Default Restriction type
	DefaultRestriction LARestrictionType = 0
)
