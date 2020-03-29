package spatialunit

// LAAreaValue Area value
type LAAreaValue struct {
	AreaSize Area
	Type     LAAreaType
}

// Area Area
type Area string

// LAAreaType Area type
type LAAreaType int

const (
	// DefualtArea Default Area type
	DefualtArea LAAreaType = 0
)
