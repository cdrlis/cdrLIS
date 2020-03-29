package surveying

// LATransformation Transformation
type LATransformation struct {
	Transformation      CCOperationMethod
	TransformedLocation GMPoint
}

// CCOperationMethod Operation method
type CCOperationMethod string // TODO external package
