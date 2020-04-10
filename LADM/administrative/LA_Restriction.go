package administrative

//
// Administrative::LA_Restriction
//
// An instance of class LA_Restriction is a restriction. LA_Restriction is a subclass of class LA_RRR.
// LA_Mortgage is a specialization of LA_Restriction (6.4.6); see Figure 10.

type LARestriction struct {
	LARRR

	PartyRequired *bool
	Type          LARestrictionType
}

// LARestrictionType Restriction type
type LARestrictionType string

const (
	AdminPublicServitude LARestrictionType = "adminPublicServitude"
	Monument                               = "monument"
	MonumentPartly                         = "monumentPartly"
	NoBuilding                             = "noBuilding"
	Servitude                              = "servitude"
	ServitudePartly                        = "servitudePartly"
)
