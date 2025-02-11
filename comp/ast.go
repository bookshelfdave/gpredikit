package comp

func ProcessInputFiles() {
	Run()
}

type AstFile struct {
	Filename string
	Checks   []*AstCheckDef
	Tools    []*AstToolDef
}

type ContentAddress struct {
	Line     int
	Col      int
	Filename string
}

type AstCheckDef struct {
	FnName       string
	ActualParams []*ActualParam
	Children     []*AstCheckDef

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
