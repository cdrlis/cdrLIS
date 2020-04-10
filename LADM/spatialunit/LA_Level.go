package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/common"

// Set of spatial units (4.1.23), with a geometric, and/or topological, and/or thematic coherence

type LALevel struct {
	common.VersionedObject

	LID          common.Oid
	Name         *string
	RegisterType *LARegisterType
	Structure    *LAStructureType
	Type         *LALevelContentType

	Su []LASpatialUnit // suLevel
}

//
// LA_RegisterType: the LA_RegisterType code list includes all the various register types, such as rural or
// urban, used in a specific land administration profile implementation. The LA_RegisterType code list is
// required only if the attribute registerType in LA_Level class is implemented. The code list shall provide a
// complete list of all codes with a name and description.

type LARegisterType string

const (
	Urban       LARegisterType = "urban"
	Rural                      = "rural"
	Mining                     = "mining"
	PublicSpace                = "publicSpace"
	Forest                     = "forest"
	All                        = "all"
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
	Polygon                          = "polygon"
	Text                             = "text"
	Topological                      = "topological"
	UnstructuredLine                 = "unstructuredLine"
	Sketch                           = "sketch"
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
	Customary                         = "customary"
	MixedLCT                          = "mixed" // Due to conflict with LASurfaceRelationType, Mixed can't be used.
	Network                           = "network"
	PrimaryRight                      = "primaryRight"
	Responsibility                    = "responsibility"
	Restriction                       = "restriction"
	Informal                          = "informal"
)
