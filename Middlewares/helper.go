package Middlewares

import (
	"fmt"
	"reflect"
	"testing"
)

func CheckNillError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckDeepEqual(expected any, got any, i int, t *testing.T) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Testcase: %v FAILED (Expected %v Found %v", i+1, got, expected)
	} else {
		fmt.Println("Testcase:", i+1, " PASSED")
	}
}

func CheckError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("Error Found from function: %s", err.Error())
	}
}
