package spatialunit

// LAVolumeValue Volume value
type LAVolumeValue struct {
	VolumeSize Volume
	Type       LAVolumeType
}

// Volume Volume
type Volume string

// LAVolumeType Volume type
type LAVolumeType int

const (
	// DefualtVolume Default volume type
	DefualtVolume LAVolumeType = 0
)
