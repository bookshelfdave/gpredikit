package functions

import (
	"fmt"

	rt "github.com/bookshelfdave/gpredikit/runtime"
)

func RegisterBuiltins(reg *rt.ChkRegistry) {
	// core groups
	reg.RegisterFn(defGroupAll())
	reg.RegisterFn(defGroupAll())
	reg.RegisterFn(defGroupAny())
	reg.RegisterFn(defGroupNone())

	// logic things for testing groups
	reg.RegisterFn(defTrue())
	reg.RegisterFn(defFalse())

	// everything else
	reg.RegisterFn(defFileExists())
	reg.RegisterFn(defOnPath())
	reg.RegisterFn(defIsFile())
	reg.RegisterFn(defIsDir())
}

func defGroupAll() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "all",
		CheckFunction: func(actualParams *rt.ActualParams, runEnv *rt.RunEnv, parent *rt.ChkInstance) *rt.ChkResult {
			failures := false
			for _, inst := range parent.Children {
				title, has_title := inst.GetTitle()
				if has_title {
					fmt.Printf("Check '%s' %s\n", title, inst.BuildPath())
				} else {
					fmt.Printf("Check %s\n", inst.BuildPath())
				}

				rt.RunCheckMaybeRetry(inst, runEnv)
				if !inst.Result.PassFail {
					failures = true
				}
			}
			if failures {
				return rt.ResFail()
			} else {
				return rt.ResPass()
			}
		},
		FormalParams:    rt.EmptyFormalParams(),
		IsGroup:         true,
		AcceptsChildren: true,
	}
}

func defGroupAny() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "any",
		CheckFunction: func(actualParams *rt.ActualParams, runEnv *rt.RunEnv, parent *rt.ChkInstance) *rt.ChkResult {
			passes := false
			for _, inst := range parent.Children {
				rt.RunCheckMaybeRetry(inst, runEnv)
				if inst.Result.PassFail {
					passes = true
				}
			}
			if passes {
				return rt.ResPass()
			} else {
				return rt.ResFail()
			}
		},
		FormalParams:    rt.EmptyFormalParams(),
		IsGroup:         true,
		AcceptsChildren: true,
	}
}

func defGroupNone() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "none",
		CheckFunction: func(actualParams *rt.ActualParams, runEnv *rt.RunEnv, parent *rt.ChkInstance) *rt.ChkResult {
			passes := false
			for _, inst := range parent.Children {
				if inst.Result.PassFail {
					passes = true
				}
			}
			if passes {
				return rt.ResFail()
			} else {
				return rt.ResPass()
			}
		},
		FormalParams:    rt.EmptyFormalParams(),
		IsGroup:         true,
		AcceptsChildren: true,
	}
}

func defTrue() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "true",
		CheckFunction: func(actualParams *rt.ActualParams, runEnv *rt.RunEnv, parent *rt.ChkInstance) *rt.ChkResult {
			return rt.ResOk(true)
		},
		FormalParams:    rt.EmptyFormalParams(),
		IsGroup:         false,
		AcceptsChildren: false,
	}
}

func defFalse() *rt.ChkDef {
	return &rt.ChkDef{
		Name: "false",
		CheckFunction: func(actualParams *rt.ActualParams, runEnv *rt.RunEnv, parent *rt.ChkInstance) *rt.ChkResult {
			return rt.ResOk(false)
		},
		FormalParams:    rt.EmptyFormalParams(),
		IsGroup:         false,
		AcceptsChildren: false,
	}
}
