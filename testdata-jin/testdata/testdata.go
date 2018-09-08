package testdata

import "data"

type Tests struct {
	Dis  string
	Par  interface{}
	Want interface{}
}

const (
	ErrorManIsTure = iota
	ErrorManIsFalse
)

type Type1 struct {
	data data.Struct1
}

func (d *Type1) ManError(errorMan data.StructTmp) {
	d.data.Sex = errorMan
}
func (d *Type1) GetData() interface{} {
	return *d
}

type Type2 struct {
	data data.Struct2
}

type testDataType interface {
	ManError(errorMan data.StructTmp)
	GetData() interface{}
}

func GetTestDataSet(dataType testDataType, errorType int) Tests {
	TestsRet := Tests{}
	//dataType.getTestSets(errorType)
	TestsRet = getErrorType(dataType, errorType)
	return TestsRet
}

func getErrorType(dataType testDataType, errorType int) Tests {
	TestsRet := Tests{}
	switch errorType {
	case ErrorManIsTure:
		TestsRet.Dis = "ErrorManIsTure"
		tmp := getManError(errorType)
		dataType.ManError(tmp)
		TestsRet.Par = dataType.GetData()
		TestsRet.Want = "exp"
	case ErrorManIsFalse:
		TestsRet.Dis = "ErrorManIsFalse"
		tmp := getManError(errorType)
		dataType.ManError(tmp)
		TestsRet.Par = dataType.GetData()
		TestsRet.Want = "asd"
	}
	return TestsRet
}

func getManError(errorType int) data.StructTmp {
	ret := data.StructTmp{}
	switch errorType {
	case ErrorManIsTure:
		ret.Man = true
	case ErrorManIsFalse:
		ret.Man = false
	}
	return ret
}
