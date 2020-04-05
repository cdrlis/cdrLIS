package surveying

import "github.com/cdrlis/cdrLIS/LADM/common/geometry"

// LATransformation Transformation
type LATransformation struct {
	Transformation      CCOperationMethod
	TransformedLocation *geometry.GMPoint
}

// CCOperationMethod Operation method
type CCOperationMethod string // TODO external package
