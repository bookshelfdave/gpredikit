package runtime

import (
	"fmt"
	"os/exec"

	"github.com/logrusorgru/aurora/v4"
)

type OutputFormatter interface {
	Init()
	CheckStart(cd *ChkInstance)
	CheckEnd(cd *ChkInstance)
	Term()
}

// TODO: this code is pretty horrible, clean it up, ok?
type DefaultOutputFormatter struct {
	au *aurora.Aurora
}

func NewDefaultOutputFormatter() DefaultOutputFormatter {
	return DefaultOutputFormatter{
		au: aurora.New(aurora.WithColors(true)),
	}
}

func (f DefaultOutputFormatter) Init() {
}

func (f DefaultOutputFormatter) CheckStart(ci *ChkInstance) {

	i := 0
	for tmp := ci.Parent; tmp != nil; tmp = tmp.Parent {
		i += 1
		fmt.Printf("   ")
	}

	fmt.Printf("- Check %s: ", ci.BuildPath())

	if len(ci.Children) > 0 {
		fmt.Println("")
	}

}

func (f DefaultOutputFormatter) CheckEnd(ci *ChkInstance) {
	if len(ci.Children) > 0 {
		i := 0
		for tmp := ci.Parent; tmp != nil; tmp = tmp.Parent {
			i += 1
			fmt.Printf("   ")
		}
		fmt.Printf("- Group result %s: ", ci.BuildPath())
	}

	if ci.Result.Err != nil {
		err := fmt.Sprintf("error %+v", ci.Result.Err)
		fmt.Println(f.au.Yellow(err))
	} else {
		if ci.Result.PassFail {
			fmt.Println(f.au.Green("PASS"))
		} else {
			fmt.Println(f.au.Red("FAIL"))

		}
	}

}

func (f DefaultOutputFormatter) Term() {
}

type RunEnv struct {
	Formatter OutputFormatter
}

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
	PassFail bool
	Err      error
	Output   *ChkOutput
}

func ResOk(res bool) *ChkResult {
	return &ChkResult{
		PassFail: res,
	}
}

func ResPass() *ChkResult {
	return &ChkResult{
		PassFail: true,
	}
}

func ResFail() *ChkResult {
	return &ChkResult{
		PassFail: false,
	}
}

func ResError(err error) *ChkResult {
	return &ChkResult{
		PassFail: false,
		Err:      err,
	}
}

// ChkInstance is a fully compiled AstChkInstance
type ChkInstance struct {
	Name         string
	Parent       *ChkInstance
	Def          *ChkDef
	ActualParams *ActualParams
	Children     []*ChkInstance
	IsNegated    bool
	IsRetrying   bool
	IsQuery      bool
	InstanceID   uint
	Address      ContentAddress
	Result       *ChkResult
}

func (inst *ChkInstance) BuildPath() string {
	if inst.Parent != nil {
		return fmt.Sprintf("%s.%s", inst.Parent.BuildPath(), inst.Name)
	} else {
		return inst.Name
	}
}

func (inst *ChkInstance) Run(runEnv *RunEnv) *ChkResult {
	return inst.Def.CheckFunction(inst.ActualParams, runEnv, inst)
}

func (inst *ChkInstance) RunHookIfDefined(hook_name string, runEnv *RunEnv) error {
	if hook_ap, ok := inst.ActualParams.Get(hook_name); ok {
		hook_cmdline, err := hook_ap.GetString()
		if err != nil {
			return fmt.Errorf("non-string value passed for hook %s: %w", hook_name, err)
		}

		// TODO: add hook timeouts?
		// TODO: make shell configurable
		//if err := exec.CommandContext(ctx, "bash", "-c", hook_cmdline).Run(); err != nil {
		out, err := exec.Command("bash", "-c", hook_cmdline).Output()
		if err != nil {
			fmt.Printf("Hook err %s\n", err)
		} else {
			fmt.Printf("%s", out)
		}
	}
	return nil
}

func (inst *ChkInstance) GetTitle() (string, bool) {
	if t, ok := inst.ActualParams.Get("title"); ok {
		if ts, err := t.GetString(); err != nil {
			// I suspect that I'm not using this correctly
			fmt.Printf("Warning: Type error while getting 'title' for %s at file %s %d:%d",
				inst.Name, inst.Address.Filename, inst.Address.Line, inst.Address.Col)
			return "", false
		} else {
			return ts, true
		}

	} else {
		return "", false
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
