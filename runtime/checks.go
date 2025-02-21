package runtime

import "fmt"

type RunEnv struct{}

type ChkFn func(actualParams *ActualParams, runEnv *RunEnv, chkInst *ChkInstance) *ChkResult

type ChkDef struct {
	Name            string
	CheckFunction   ChkFn
	FormalParams    map[string]*FormalParam
	AcceptsChildren bool
	IsGroup         bool
	IsQuery         bool
	TemplateParams  map[string]*ActualParam
}

func (c *ChkDef) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("check definition requires a name")
	}

	if c.CheckFunction == nil {
		return fmt.Errorf("check definition requires a check function")
	}

	if c.FormalParams == nil {
		return fmt.Errorf("check definition requires formal parameters (use EmptyFormalParams() for no params)")
	}

	return nil
}

func EmptyFormalParams() map[string]*FormalParam {
	return make(map[string]*FormalParam)
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

func ResOk(res bool) *ChkResult {
	return &ChkResult{
		TestResult: res,
	}
}

func ResPass() *ChkResult {
	return &ChkResult{
		TestResult: true,
	}
}

func ResFail() *ChkResult {
	return &ChkResult{
		TestResult: false,
	}
}

func ResError(err error) *ChkResult {
	return &ChkResult{
		TestResult: false,
		Err:        err,
	}
}

// ChkInstance is a fully compiled AstChkInstance
type ChkInstance struct {
	Name         string
	Parent       *ChkInstance
	Title        *string
	Def          *ChkDef
	ActualParams *ActualParams
	Children     []*ChkInstance
	IsNegated    bool
	IsRetrying   bool
	IsQuery      bool
	InstanceID   uint
	Address      ContentAddress
}

func (c *ChkInstance) BuildPath() string {
	if c.Parent != nil {
		return fmt.Sprintf("%s.%s", c.Parent.BuildPath(), c.Name)
	} else {
		return c.Name
	}
}

var nextFileId = 0

type CompiledFileOut struct {
	Id        int // uint?
	Filename  string
	Errors    []error
	Instances []*ChkInstance
}

func NewCompiledFileOut(filename string) *CompiledFileOut {
	id := nextFileId
	nextFileId++
	return &CompiledFileOut{
		Id:       id,
		Filename: filename,
	}
}

func (c *CompiledFileOut) AddError(err error) {
	c.Errors = append(c.Errors, err)
}

func (c *CompiledFileOut) AddErrors(errs []error) {
	c.Errors = append(c.Errors, errs...)
}

func (c *CompiledFileOut) AddChkInstance(instance *ChkInstance) {
	c.Instances = append(c.Instances, instance)
}
