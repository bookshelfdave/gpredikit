package commands

import (
	"context"
	"fmt"
	"os"

	"github.com/bookshelfdave/gpredikit/comp"
	rt "github.com/bookshelfdave/gpredikit/runtime"
	"github.com/urfave/cli/v3"
)

func VerifyInputFiles(files []string) error {
	dups := make(map[string]bool, len(files))
	for _, f := range files {
		if _, ok := dups[f]; ok {
			return fmt.Errorf("file %s specified as input more than once", f)
		}
		dups[f] = true
		if _, err := os.Stat(f); err != nil {
			return fmt.Errorf("invalid input file %s", f)
		}
	}
	return nil
}

func RunCommand() *cli.Command {
	return &cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run predikit test files",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			if cmd.Args().Len() == 0 {
				fmt.Println("Specify a file or space separated list of files to compile")
				return nil
			}
			files := cmd.Args().Slice()
			if err := VerifyInputFiles(files); err != nil {
				fmt.Printf("Fatal error: %s\n", err)
				return nil
			}
			// TODO: move all error reporting up to the top level
			compiledFilesOut := comp.CompileInputFiles(files)
			if compiledFilesOut == nil {
				return nil
			}
			of := rt.NewDefaultOutputFormatter()
			comp.RunChecks(compiledFilesOut, of)
			return nil
		},
	}
}
