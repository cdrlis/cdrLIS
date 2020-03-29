package administrative

// LARight Right
type LARight struct {
	LARRR
	Type LARightType
}

// LARightType Right type
type LARightType int

const (
	// DefaultRight Default right type
	DefaultRight LARightType = 0
)
