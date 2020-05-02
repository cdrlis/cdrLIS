package ladm

import "github.com/cdrlis/cdrLIS/LADM/common"

// Set of spatial units (4.1.23), with a geometric, and/or topological, and/or thematic coherence

type LALevel struct {
	common.VersionedObject

	ID           string              `gorm:"column:id;primary_key" json:"-"`

	LID          common.Oid          `gorm:"column:lid" json:"lID"`
	Name         *string             `gorm:"column:name" json:"name"`
	RegisterType *LARegisterType     `gorm:"column:registertype" json:"registerType"`
	Structure    *LAStructureType    `gorm:"column:structure" json:"structure"`
	Type         *LALevelContentType `gorm:"column:type" json:"type"`

	SU []LASpatialUnit `gorm:"foreignkey:LevelID,LevelBeginLifespanVersion;association_foreignkey:ID,BeginLifespanVersion" json:"level"`
}

func (LALevel) TableName() string {
	return "LA_Level"
}

//
// LA_RegisterType: the LA_RegisterType code list includes all the various register types, such as rural or
// urban, used in a specific land administration profile implementation. The LA_RegisterType code list is
// required only if the attribute registerType in LA_Level class is implemented. The code list shall provide a
// complete list of all codes with a name and description.

type LARegisterType string

const (
	Urban       LARegisterType = "urban"
	Rural       LARegisterType = "rural"
	Mining      LARegisterType = "mining"
	PublicSpace LARegisterType = "publicSpace"
	Forest      LARegisterType = "forest"
	All         LARegisterType = "all"
)

//
// LA_StructureType: the LA_StructureType code list includes all the various spatial structure types, such as
// point or polygon, used in a specific land administration profile implementation. The LA_StructureType
// code list is required only if the attribute structure in LA_Level class is implemented. The code list shall
// provide a complete list of all codes with a name and description.
//

type LAStructureType string

const (
	Point            LAStructureType = "point"
	Polygon          LAStructureType = "polygon"
	Text             LAStructureType = "text"
	Topological      LAStructureType = "topological"
	UnstructuredLine LAStructureType = "unstructuredLine"
	Sketch           LAStructureType = "sketch"
)

//
// LA_LevelContentType: the LA_LevelContentType code list includes all the various level content types,
// such as primary right or customary, used in a specific land administration profile implementation. The
// LA_LevelContentType code list is required only if the attribute type in LA_Level class is implemented. The
// code list shall provide a complete list of all codes with a name and description.
//

type LALevelContentType string

const (
	Building       LALevelContentType = "building"
	Customary      LALevelContentType = "customary"
	MixedLCT       LALevelContentType = "mixed" // Due to conflict with LASurfaceRelationType, Mixed can't be used.
	Network        LALevelContentType = "network"
	PrimaryRight   LALevelContentType = "primaryRight"
	Responsibility LALevelContentType = "responsibility"
	Restriction    LALevelContentType = "restriction"
	Informal       LALevelContentType = "informal"
)
