package comp

import (
	"errors"
	"fmt"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

type CompileErrors struct {
	Errors []error
}

func NewCompilerErrors() *CompileErrors {
	return &CompileErrors{
		Errors: make([]error, 0),
	}
}

func makeCheckInstance(chk *AstChkInstance, reg *rt.ChkRegistry) (*rt.ChkInstance, error) {
	newInst := &rt.ChkInstance{}

	// try to make all children first
	for _, child := range chk.Children {
		childInst, err := makeCheckInstance(child, reg)
		if err != nil {
			return nil, err
		}
		newInst.Children = append(newInst.Children, childInst)
	}

	if val, ok := reg.Fns[chk.FnName]; !ok {
		fmt.Println("Invalid function")
		errMsg := fmt.Sprintf("Unknown function '%s' specified at %s:%d:%d",
			chk.FnName, chk.Address.Filename, chk.Address.Line, chk.Address.Col)
		e := errors.New(errMsg)
		return nil, e
	} else {
		fmt.Printf("FN = %+v\n", val)
	}
	return newInst, nil
}

func CompileChecksToAsts(astFiles []*AstFile, reg *rt.ChkRegistry) {
	for _, astFile := range astFiles {
		for _, checkDef := range astFile.Checks {
			inst, err := makeCheckInstance(checkDef, reg)
			if err != nil {
				// just for now
				fmt.Printf("Error: %s\n", err)
			}
			fmt.Printf("%+v\n", inst)
		}
	}
}
