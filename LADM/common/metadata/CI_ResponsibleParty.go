package metadata

import (
	"net/mail"
	"net/url"
	"time"
)

type CI_Citation struct {
	Title         string
	AltemateTitle *string
	Date          []CI_Date
	Edition       *string
	EditionDate   *time.Time
	//	Identifier				MD_Identifier
	CitedResponsibleParty *CI_ResponsibleParty
	PresentationForm      *CI_PresentationFormCode
	Series                *CI_Series
	OtherCitationDetails  *string
	CollectiveTitle       *string
	ISBN                  *string
	ISSN                  *string
}

type CI_Date struct {
	date     time.Time
	DateType CI_DateTypeCode
}

type CI_DateTypeCode string

const (
	Creation    CI_DateTypeCode = "creation"
	Publication                 = "publication"
	Revision                    = "revision"
)

type CI_ResponsibleParty struct {
	IndividualName   *string
	OrganizationName *string
	PositionName     *string
	ContactInfo      *CI_Contact
	Role             CI_RoleCode
}

type CI_Contact struct {
	Phone               CI_Telephone
	Address             CI_Address
	OnLineResource      CI_OnLineResource
	HoursOfService      string
	ContactInstructions string
}

type CI_Telephone struct {
	Voice     []string
	Facsimile []string
}

type CI_Address struct {
	DeliveryPoint         []string
	City                  string
	AdministrativeArea    string
	PostalCode            string
	Country               string
	ElectronicMailAddress []mail.Address
}

type CI_RoleCode int

const (
	ResourceProvider CI_RoleCode = iota
	Custodian
	Owner
	User
	Distributor
	Originator
	PointOfContact
	PrincipalInvestigator
	Processor
	Publisher
	Author
)

type CI_OnLineResource struct {
	Linkage            url.URL
	Protocol           string
	ApplicationProfile string
	Name               string
	Description        string
	Function           CI_OnLineFunctionCode
}

type CI_OnLineFunctionCode string

const (
	Download      CI_OnLineFunctionCode = "download"
	Information                         = "information"
	OffLineAccess                       = "offLineAccess"
	Order                               = "order"
	Search                              = "search"
)

type CI_PresentationFormCode string

const (
	DocumentDigital  CI_PresentationFormCode = "documentDigital"
	DocumentHardcopy                         = "documentHardcopy"
	ImageDigital                             = "imageDigital"
	ImageHardcopy                            = "imageHardcopy"
	MapDigital                               = "mapDigital"
	MapHardcopy                              = "mapHardcopy"
	ModelDigital                             = "modelDigital"
	ModelHardcopy                            = "modelHardcopy"
	ProfileDigital                           = "profileDigital"
	TableDigital                             = "tableDigital"
	TableHardcopy                            = "tableHardcopy"
	VideoDigital                             = "videoDigital"
	VideoHardcopy                            = "videoHardcopy"
)

type CI_Series struct {
	Name                string
	IssueIdentification string
	Page                string
}
