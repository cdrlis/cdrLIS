package shared

// VersionedObject Versioned object
type VersionedObject struct {
	BeginLifespanVersion string
	EndLifespanVersion   *string
	Quality              *DQElement
	Source               *CIResponsibleParty
}

// DQElement DQ Elemenet
type DQElement string // TODO external package

// CIResponsibleParty CI Responsible party
type CIResponsibleParty string // TODO external package
