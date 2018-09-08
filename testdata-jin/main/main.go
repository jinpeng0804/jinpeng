package main

import (
	"fmt"
	"testdata"
)

func main() {
	// testSet := testdata.Tests{}
	// a := data.Struct1{}
	//a := testdata.GetTestDataSet(&testdata.Type1{}, testdata.ErrorManIsTure)
	a := testdata.GetTestDataSet(&testdata.Type1{}, testdata.ErrorManIsFalse)
	fmt.Println(a)
}
