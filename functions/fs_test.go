package functions

import (
	"testing"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func TestFileExists(t *testing.T) {
	cd := defFileExists()
	runEnv := &rt.RunEnv{}

	actualParams := rt.MakeTestParams(cd, rt.TestParams{
		"path": "/bin/sh",
	})

	instance := &rt.ChkInstance{}

	result := cd.CheckFunction(actualParams, runEnv, instance)
	if !result.PassFail {
		t.Fail()
	}
}

func TestIsDir(t *testing.T) {
	cd := defIsDir()
	runEnv := &rt.RunEnv{}

	actualParams := rt.MakeTestParams(cd, rt.TestParams{
		"path": "/bin/sh",
	})
	instance := &rt.ChkInstance{}

	result := cd.CheckFunction(actualParams, runEnv, instance)
	if !result.PassFail {
		t.Fail()
	}
}

func TestIsFile(t *testing.T) {
	cd := defIsFile()
	runEnv := &rt.RunEnv{}

	actualParams := rt.MakeTestParams(cd, rt.TestParams{
		"path": "/bin/sh",
	})
	instance := &rt.ChkInstance{}

	result := cd.CheckFunction(actualParams, runEnv, instance)
	if !result.PassFail {
		t.Fail()
	}
}

func TestOnPath(t *testing.T) {
	cd := defOnPath()
	runEnv := &rt.RunEnv{}

	actualParams := rt.MakeTestParams(cd, rt.TestParams{
		"path": "/bin/sh",
	})
	instance := &rt.ChkInstance{}

	result := cd.CheckFunction(actualParams, runEnv, instance)

	if !result.PassFail {
		t.Fail()
	}
}
