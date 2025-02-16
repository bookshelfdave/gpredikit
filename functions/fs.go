package functions

import (
	"errors"
	"os"
	"os/exec"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func DefFileExists() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "exists?",
		CheckFunction: func(actualParams map[string]*rt.ActualParam, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path := actualParams["path"]
			spath, _ := path.GetString()

			if _, err := os.Stat(spath); err == nil {
				return rt.ChkPass()
			} else if errors.Is(err, os.ErrNotExist) {
				return rt.ChkFail()
			} else {
				return rt.ChkError(err)
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func DefIsDir() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "is_dir?",
		CheckFunction: func(actualParams map[string]*rt.ActualParam, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path := actualParams["path"]
			spath, _ := path.GetString()

			if stat, err := os.Stat(spath); err == nil {
				return rt.ChkOk(stat.IsDir())
			} else if errors.Is(err, os.ErrNotExist) {
				return rt.ChkOk(false)
			} else {
				return rt.ChkError(err)
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func DefIsFile() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "is_file?",
		CheckFunction: func(actualParams map[string]*rt.ActualParam, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path := actualParams["path"]
			spath, _ := path.GetString()

			if stat, err := os.Stat(spath); err == nil {
				return rt.ChkOk(stat.Mode().IsRegular())
			} else if errors.Is(err, os.ErrNotExist) {
				return rt.ChkOk(false)
			} else {
				return rt.ChkError(err)
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func DefOnPath() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "on_path?",
		CheckFunction: func(actualParams map[string]*rt.ActualParam, _ *rt.RunEnv, _ *rt.ChkInstance) *rt.ChkResult {
			path := actualParams["cmd"]
			spath, _ := path.GetString()

			// TODO: it would be nice to log the path that's returned
			if _, err := exec.LookPath(spath); err == nil {
				return rt.ChkPass()
			} else {
				return rt.ChkFail()
			}
		},
		FormalParams: rt.BuildFormalParams().
			WithRequiredString("cmd").
			Build(),
		IsGroup: false,
	}
}
