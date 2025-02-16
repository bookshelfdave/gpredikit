package functions

import (
	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func RegisterBuiltIns(reg *rt.ChkRegistry) {
	reg.Fns["exists?"] = DefFileExists()
	reg.Fns["on_path?"] = DefOnPath()
	reg.Fns["is_file?"] = DefIsFile()
	reg.Fns["is_dir?"] = DefIsDir()
}
