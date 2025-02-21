package functions

import (
	"fmt"
	"testing"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func TestOnePlusOne(t *testing.T) {
	cd := DefShell()
	actualParams := rt.MakeTestParams(cd, rt.TestParams{
		"cmd": "sleep 10",
	})

	runEnv := &rt.RunEnv{}
	result := cd.CheckFunction(actualParams, runEnv, nil)
	fmt.Printf("---> %+v", result)

	// if result != expected {
	// 	t.Errorf("Expected %d but got %d", expected, result)
	// }
}
