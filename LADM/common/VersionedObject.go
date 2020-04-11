package common

import (
	"github.com/cdrlis/cdrLIS/LADM/common/metadata"
	"time"
)

// Type VersionedObject is introduced in the LADM to manage and maintain historical data in the database.
// History requires, that inserted and superseded data, are given a time-stamp. In this way, the contents of the
// database can be reconstructed, as they were at any historical moment. For more on history and dynamic
// aspects of LA systems, see Annex N.
type VersionedObject struct {
	BeginLifespanVersion time.Time 	`gorm:"column:beginlifespanversion"`
	EndLifespanVersion   *time.Time	`gorm:"column:endlifespanversion"`
	Quality              *metadata.DQ_Element
	Source               *metadata.CI_ResponsibleParty
}

type TimeObject time.Time