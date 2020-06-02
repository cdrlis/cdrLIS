package common

import (
	"github.com/cdrlis/cdrLIS/ladm/common/metadata"
	"time"
)

// Type VersionedObject is introduced in the LADM to manage and maintain historical data in the database.
// History requires, that inserted and superseded data, are given a time-stamp. In this way, the contents of the
// database can be reconstructed, as they were at any historical moment. For more on history and dynamic
// aspects of LA systems, see Annex N.
type VersionedObject struct {
	BeginLifespanVersion time.Time 	`gorm:"column:beginlifespanversion;primary_key" json:"beginLifespanVersion"`
	EndLifespanVersion   *time.Time	`gorm:"column:endlifespanversion" json:"endLifespanVersion"`
	Quality              *metadata.DQ_Element `json:"quality"`
	Source               *metadata.CI_ResponsibleParty `json:"source"`
}