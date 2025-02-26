package main

import (
	"github.com/bookshelfdave/gpredikit/comp"
	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func main() {
	files := []string{"./checks/first.pk"}
	// TODO: move all error reporting up to the top level
	compiledFilesOut := comp.CompileInputFiles(files)
	if compiledFilesOut == nil {
		return
	}
	of := rt.NewDefaultOutputFormatter()
	comp.RunChecks(compiledFilesOut, of)
}
