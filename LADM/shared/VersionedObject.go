package shared

// VersionedObject Versioned object
type VersionedObject struct {
	ID                   string
	BeginLifespanVersion string
	EndLifespanVersion   string
}
