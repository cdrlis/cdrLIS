package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/common"

// Set of spatial units (4.1.23), with a geometric, and/or topological, and/or thematic coherence

type LALevel struct {
	LID          common.Oid
	Name         *string
	RegisterType *LARegisterType
	Structure    *LAStructureType
	Type         *LALevelContentType
}

//
// LA_RegisterType: the LA_RegisterType code list includes all the various register types, such as rural or
// urban, used in a specific land administration profile implementation. The LA_RegisterType code list is
// required only if the attribute registerType in LA_Level class is implemented. The code list shall provide a
// complete list of all codes with a name and description.

type LARegisterType int

const (
	// DefaultRegister Default register type
	DefaultRegister LARegisterType = 0
	Urban
	Rural
	Mining
	PublicSpace
	Forest
	All
)

//
// LA_StructureType: the LA_StructureType code list includes all the various spatial structure types, such as
// point or polygon, used in a specific land administration profile implementation. The LA_StructureType
// code list is required only if the attribute structure in LA_Level class is implemented. The code list shall
// provide a complete list of all codes with a name and description.
//

type LAStructureType int

const (
	// DefaultStructure Default structure type
	DefaultStructure LAStructureType = 0
	Point
	Polygon
	Text
	Topological
	UnstructuredLine
	Sketch
)

//
// LA_LevelContentType: the LA_LevelContentType code list includes all the various level content types,
// such as primary right or customary, used in a specific land administration profile implementation. The
// LA_LevelContentType code list is required only if the attribute type in LA_Level class is implemented. The
// code list shall provide a complete list of all codes with a name and description.
//

type LALevelContentType int

const (
	// DefaultLevel Default level type
	DefaultLevel LALevelContentType = 0
	Building
	Customary
	MixedLCT // Due to conflict with LASurfaceRelationType, Mixed can't be used.
	Network
	PrimaryRight
	Responsibility
	Restriction
	Informal
)
