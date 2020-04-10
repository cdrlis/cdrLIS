package surveying

import (
	"github.com/cdrlis/cdrLIS/LADM/common/geometry"
	"github.com/cdrlis/cdrLIS/LADM/common/metadata"
)

// LATransformation Transformation
type LATransformation struct {
	Transformation      CCOperationMethod
	TransformedLocation *geometry.GMPoint
}

// CCOperationMethod from ISO 19111
type CCOperationMethod struct {
	FormulaReference CCFormula
	SourceDimension  int
	TargetDimension  int
}

// CC Formula from ISO 19111
type CCFormula struct {
	Formula         string
	FormulaCitation metadata.CI_Citation
}
