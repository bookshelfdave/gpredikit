package functions

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func DefShell() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "shell!",
		CheckFunction: func(actualParams *rt.ActualParams, _ *rt.RunEnv, inst *rt.ChkInstance) *rt.ChkResult {
			cmd, _ := actualParams.GetStringOrDefault("cmd")
			delay := time.Duration(5 * time.Second)
			ctx, cancel := context.WithTimeout(context.Background(), delay)
			defer cancel()

			// not really what I want, but I want to see something work
			if err := exec.CommandContext(ctx, "bash", "-c", cmd).Run(); err != nil {
				fmt.Printf("fn err %s\n", err)
				return &rt.ChkResult{
					TestResult: false,
					Err:        err,
				}
			} else {
				fmt.Println("Success")
				return rt.ResPass()
			}

		},
		FormalParams:    rt.BuildFormalParams().WithRequiredString("cmd").WithOptionalDuration("timeout", time.Duration(30*time.Second)).Build(),
		AcceptsChildren: false,
		IsGroup:         false,
		IsQuery:         false,
		TemplateParams:  map[string]*rt.ActualParam{},
	}
}
