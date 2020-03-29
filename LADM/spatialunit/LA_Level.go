package spatialunit

import "github.com/cdrlis/cdrLIS/LADM/shared"

// LALevel Level
type LALevel struct {
	LID          shared.Oid
	Name         string
	RegisterType LARegisterType
	Structure    LAStructureType
	Type         LALevelContentType
}

// LARegisterType Register type
type LARegisterType int

const (
	// DefaultRegister Default register type
	DefaultRegister LARegisterType = 0
)

// LAStructureType Structure type
type LAStructureType int

const (
	// DefaultStructure Default structure type
	DefaultStructure LAStructureType = 0
)

// LALevelContentType Level content type
type LALevelContentType int

const (
	// DefaultLevel Default level type
	DefaultLevel LALevelContentType = 0
)
