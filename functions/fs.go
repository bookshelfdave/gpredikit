package functions

import (
	"errors"
	"os"
	"os/exec"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func defFileExists() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "exists?",
		CheckFunction: func(actualParams *rt.ActualParams, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path, _ := actualParams.GetStringOrDefault("path")

			if _, err := os.Stat(path); err == nil {
				return rt.ResPass()
			} else if errors.Is(err, os.ErrNotExist) {
				return rt.ResFail()
			} else {
				return rt.ResError(err)
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func defIsDir() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "is_dir?",
		CheckFunction: func(actualParams *rt.ActualParams, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path, _ := actualParams.GetStringOrDefault("path")

			if stat, err := os.Stat(path); err == nil {
				return rt.ResOk(stat.IsDir())
			} else if errors.Is(err, os.ErrNotExist) {
				return rt.ResOk(false)
			} else {
				return rt.ResError(err)
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func defIsFile() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "is_file?",
		CheckFunction: func(actualParams *rt.ActualParams, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path, _ := actualParams.GetStringOrDefault("path")

			if stat, err := os.Stat(path); err == nil {
				return rt.ResOk(stat.Mode().IsRegular())
			} else if errors.Is(err, os.ErrNotExist) {
				return rt.ResOk(false)
			} else {
				return rt.ResError(err)
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func defOnPath() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "on_path?",
		CheckFunction: func(actualParams *rt.ActualParams, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path, _ := actualParams.GetStringOrDefault("path")

			// TODO: it would be nice to log the path that's returned
			if _, err := exec.LookPath(path); err == nil {
				return rt.ResPass()
			} else {
				return rt.ResFail()
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("cmd").
			Build(),
		IsGroup: false,
	}
}
