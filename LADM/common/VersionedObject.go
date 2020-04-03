package common

import (
	"github.com/cdrlis/cdrLIS/LADM/common/metadata"
	"time"
)

// VersionedObject Versioned object
type VersionedObject struct {
	BeginLifespanVersion time.Time
	EndLifespanVersion   time.Time
	Quality              metadata.DQ_Element
	Source               metadata.CI_ResponsibleParty
}
