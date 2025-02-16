package runtime

type ChkDef struct {
	Name            string
	CheckFunction   ChkFn
	FormalParams    map[string]*FormalParam
	AcceptsChildren bool
	IsGroup         bool
	IsQuery         bool
	TemplateParams  map[string]*ActualParam
}

type ChkOutput struct {
	Stdout *string
	Stderr *string
	Retval *int
}

type ChkResult struct {
	TestResult bool
	Err        error
	Output     *ChkOutput
}

func ChkOk(res bool) *ChkResult {
	return &ChkResult{
		TestResult: res,
	}
}

func ChkPass() *ChkResult {
	return &ChkResult{
		TestResult: true,
	}
}

func ChkFail() *ChkResult {
	return &ChkResult{
		TestResult: false,
	}
}

func ChkError(err error) *ChkResult {
	return &ChkResult{
		TestResult: false,
		Err:        err,
	}
}

// ChkInstance is a fully compiled AstChkInstance
type ChkInstance struct {
	Name                     string
	Title                    *string
	Def                      *ChkDef
	ActualParams             []*ActualParam
	MaterializedFormalParams []*FormalParam
	Children                 []*ChkInstance
	IsNegated                bool
	IsRetrying               bool
	IsQuery                  bool
	InstanceID               uint
	Address                  ContentAddress
}

type RunEnv struct{}

type ChkFn func(actualParams map[string]*ActualParam, runEnv *RunEnv, chkInst *ChkInstance) *ChkResult
