package functions

import (
	"errors"
	"os"
	"os/exec"

	. "github.com/bookshelfdave/gpredikit/runtime"
)

func DefFileExists() *ChkDef {
	return &ChkDef{
		Name: "exists?",
		CheckFunction: func(actualParams map[string]*ActualParam, _ *RunEnv, _ *ChkInstance) *ChkResult {
			path, _ := actualParams["path"]
			spath, _ := path.GetString()

			if _, err := os.Stat(spath); err == nil {
				return ChkPass()
			} else if errors.Is(err, os.ErrNotExist) {
				return ChkFail()
			} else {
				return ChkError(err)
			}
		},
		FormalParams: BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func DefIsDir() *ChkDef {
	return &ChkDef{
		Name: "is_dir?",
		CheckFunction: func(actualParams map[string]*ActualParam, _ *RunEnv, _ *ChkInstance) *ChkResult {
			path, _ := actualParams["path"]
			spath, _ := path.GetString()

			if stat, err := os.Stat(spath); err == nil {
				return ChkOk(stat.IsDir())
			} else if errors.Is(err, os.ErrNotExist) {
				return ChkOk(false)
			} else {
				return ChkError(err)
			}
		},
		FormalParams: BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func DefIsFile() *ChkDef {
	return &ChkDef{
		Name: "is_file?",
		CheckFunction: func(actualParams map[string]*ActualParam, _ *RunEnv, _ *ChkInstance) *ChkResult {
			path, _ := actualParams["path"]
			spath, _ := path.GetString()

			if stat, err := os.Stat(spath); err == nil {
				return ChkOk(stat.Mode().IsRegular())
			} else if errors.Is(err, os.ErrNotExist) {
				return ChkOk(false)
			} else {
				return ChkError(err)
			}
		},
		FormalParams: BuildFormalParams().
			WithRequiredString("path").
			Build(),
		IsGroup: false,
	}
}

func DefOnPath() *ChkDef {
	return &ChkDef{
		Name: "on_path?",
		CheckFunction: func(actualParams map[string]*ActualParam, _ *RunEnv, _ *ChkInstance) *ChkResult {
			path, _ := actualParams["cmd"]
			spath, _ := path.GetString()

			// TODO: it would be nice to log the path that's returned
			if _, err := exec.LookPath(spath); err == nil {
				return ChkPass()
			} else {
				return ChkFail()
			}
		},
		FormalParams: BuildFormalParams().
			WithRequiredString("cmd").
			Build(),
		IsGroup: false,
	}
}
