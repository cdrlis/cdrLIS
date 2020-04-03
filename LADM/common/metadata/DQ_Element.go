package metadata

import (
	"time"
)

type DQ_Element struct {
	NameOfMeasure               string
	MeasureIdentification       MD_Identifier
	MeasureDescription          string
	EvaluationMethodType        DQ_EvaluationMethodTypeCode
	EvaluationMethodDescription string
	EvaluationProcedure         CI_Address
	DateTime                    time.Time
	Result                      DQ_Result
}

type MD_Identifier struct {
	authority CI_DateTypeCode
	code      string
}

type DQ_EvaluationMethodTypeCode int

const (
	DirectInternal DQ_EvaluationMethodTypeCode = iota
	DirectExternal
	Indirect
)

type DQ_Result struct {
	Specification int
	Explanation   string
	Pass          bool
	//	valueType		RecordType
	//	ValueUnit		UnitOfMeasure
	ErrorStatistic string
	//	Value			Record
}
