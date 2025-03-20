package comp

import (
	rt "github.com/bookshelfdave/gpredikit/runtime"
)

type AstFile struct {
	Filename string
	Checks   []*AstChkInstance
	Tools    []*AstToolDef
}

type AstChkInstance struct {
	FnName       string
	ActualParams map[string]*rt.ActualParam
	Children     []*AstChkInstance

	IsNegated  bool
	IsRetrying bool
	IsGroup    bool

	Address rt.ContentAddress
}

func NewAstChkInstance() *AstChkInstance {
	return &AstChkInstance{
		ActualParams: make(map[string]*rt.ActualParam),
	}
}

type AstToolDef struct {
	ToolName         string
	DesignTimeParams []*rt.ActualParam
	RuntimeParams    map[string]map[string]*rt.ActualParam
	Address          rt.ContentAddress
}

type AstToolInstanceParam struct {
	ParamName   string
	ParamsProps map[string]*rt.ActualParam
	Address     rt.ContentAddress
}

type AddressableString struct {
	V       string
	Address rt.ContentAddress
}
