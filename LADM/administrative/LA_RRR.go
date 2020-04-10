package administrative

import (
	"time"

	"github.com/cdrlis/cdrLIS/LADM/common"
	"github.com/cdrlis/cdrLIS/LADM/party"
)

//
// Administrative::LA_RRR
//
// LA_RRR is an abstract class with three specialization classes:
// 1) LA_Right, with rights as instances. Rights are primarily in the domain of private or customary law.
//    Ownership rights are generally based on (national) legislation, and code lists in the LADM are in
//    support of this, see Annex J.
// 2) LA_Restriction, with restrictions as instances. Restrictions usually "run with the land", meaning that
//    they remain valid, even when the right to the land is transferred after the right was created (and
//    registered). A mortgage, an instance of class LA_Mortgage, is a special restriction of the ownership
//    right. It concerns the conveyance of a property by a debtor to a creditor, as a security for a financial
//    loan, with the condition that the property is returned, when the loan is paid off.
// 3) LA_Responsibility, with responsibilities as instances.

type LARRR struct {
	common.VersionedObject

	Description *string
	RID         common.Oid
	Share       *common.Fraction
	ShareCheck  *bool
	TimeSpec    *time.Time

	Party *party.LAParty // rrrParty
	Unit  *LABAunit      // unitRrr
}
