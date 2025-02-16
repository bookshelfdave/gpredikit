package functions

import (
	"testing"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func TestFileExists(t *testing.T) {
	fn := DefFileExists()
	runEnv := &rt.RunEnv{}

	actualParams := rt.BuildActualParams().
		NewNamedParamString("path", "/bin/sh").
		Build()

	instance := &rt.ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)
	if !result.TestResult {
		t.Fail()
	}
}

func TestIsDir(t *testing.T) {
	fn := DefIsDir()
	runEnv := &rt.RunEnv{}

	actualParams := rt.BuildActualParams().
		NewNamedParamString("path", "/bin").
		Build()

	instance := &rt.ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)
	if !result.TestResult {
		t.Fail()
	}
}

func TestIsFile(t *testing.T) {
	fn := DefIsFile()
	runEnv := &rt.RunEnv{}

	actualParams := rt.BuildActualParams().
		NewNamedParamString("path", "/bin/sh").
		Build()

	instance := &rt.ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)
	if !result.TestResult {
		t.Fail()
	}
}

func TestOnPath(t *testing.T) {
	fn := DefOnPath()
	runEnv := &rt.RunEnv{}

	actualParams := rt.BuildActualParams().
		NewNamedParamString("cmd", "sh").
		Build()

	instance := &rt.ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)

	if !result.TestResult {
		t.Fail()
	}
}
