package functions

import (
	"testing"

	. "github.com/bookshelfdave/gpredikit/runtime"
)

func TestFileExists(t *testing.T) {
	fn := DefFileExists()
	runEnv := &RunEnv{}

	actualParams := BuildActualParams().
		NewNamedParamString("path", "/bin/sh").
		Build()

	instance := &ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)
	if !result.TestResult {
		t.Fail()
	}
}

func TestIsDir(t *testing.T) {
	fn := DefIsDir()
	runEnv := &RunEnv{}

	actualParams := BuildActualParams().
		NewNamedParamString("path", "/bin").
		Build()

	instance := &ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)
	if !result.TestResult {
		t.Fail()
	}
}

func TestIsFile(t *testing.T) {
	fn := DefIsFile()
	runEnv := &RunEnv{}

	actualParams := BuildActualParams().
		NewNamedParamString("path", "/bin/sh").
		Build()

	instance := &ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)
	if !result.TestResult {
		t.Fail()
	}
}

func TestOnPath(t *testing.T) {
	fn := DefOnPath()
	runEnv := &RunEnv{}

	actualParams := BuildActualParams().
		NewNamedParamString("cmd", "sh").
		Build()

	instance := &ChkInstance{}

	result := fn.CheckFunction(actualParams, runEnv, instance)

	if !result.TestResult {
		t.Fail()
	}
}
