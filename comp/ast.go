package comp

import . "github.com/bookshelfdave/gpredikit/runtime"

func ProcessInputFiles() {
	Run()
}

type AstFile struct {
	Filename string
	Checks   []*AstChkInstance
	Tools    []*AstToolDef
}

type AstChkInstance struct {
	FnName       string
	ActualParams []*ActualParam
	Children     []*AstChkInstance

	IsNegated  bool
	IsRetrying bool
	IsGroup    bool

	Address ContentAddress
}

type AstToolDef struct {
	ToolName       string
	ClassParams    []*ActualParam
	InstanceParams map[string][]*ActualParam
	Address        ContentAddress
}

type AstToolInstanceParam struct {
	ParamName   string
	ParamsProps []*ActualParam
	Address     ContentAddress
}

type AddressableString struct {
	V       string
	Address ContentAddress
}
