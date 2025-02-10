package comp

func ProcessInputFiles() {
	Run()
}

type AstFile struct {
	filename string
}

type ContentAddress struct {
}

type AstCheckDef struct {
	FnName         string
	ActualParams   string
	Children       []AstCheckDef
	ContentAddress string

	IsNegated  bool
	IsRetrying bool
	IsGroup    bool
}
